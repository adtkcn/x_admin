package common_controller

import (
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
	"x_admin/app/schema/common_schema"
	"x_admin/app/service/common_service"
	"x_admin/config"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/util"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

// UploadHandler 上传控制器
type UploadHandler struct {
	// 防止缓存击穿：同一文件 id 的并发 Serve 请求合并为一次 GetFilePath 查询
	requestGroup singleflight.Group
}

// @Summary		上传文件
// @Description	上传文件
// @Tags			common_upload-上传
// @Param			token	header		string														true	"token"
// @Param			cid		body		string														false	"分类ID"
// @Param			file	formData	file														true	"文件"
// @Success		200		{object}	response.Response{data=common_schema.CommonUploadFileResp}	"成功"
// @Router			/api/admin/common/upload/file [post]
func (uh *UploadHandler) UploadFile(c *gin.Context) {
	var uReq common_schema.CommonUploadImageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &uReq)) {
		return
	}
	file, ve := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, ve) {
		return
	}
	res, err := common_service.UploadService.UploadFile(file)
	// 访问地址 = GET /api/uploads/:id，由文件流路由按 id 查 x_common_file_hash.FilePath 返回
	resp := common_schema.CommonUploadFileResp{
		// ID:         res.ID,
		FileHashId: res.ID,
		Name:       file.Filename,
		Uri:        util.UrlUtil.HashUrl(res.ID, file.Filename), // 访问地址（完整可访问 URL）
		// Path:       res.FilePath,                 // 磁盘存储 key = <id>.<ext>
		Ext:     res.Ext,
		Size:    res.FileSize,
		Instant: false,
	}
	response.CheckAndRespWithData(c, resp, err)

	// 上传成功后异步转 webp（条件判断在 MaybeConvertWebp 内）
	common_service.UploadService.MaybeConvertWebp(res.ID, res.FilePath, res.Ext, res.FileSize)
}

// @Summary		文件秒传检查
// @Description	根据文件 MD5 查询是否已上传，命中则返回已有的文件哈希记录ID
// @Tags			common_upload-上传
// @Param			token		header		string				true	"token"
// @Param			file_md5	body		string				true	"文件MD5"
// @Param			file_name	body		string				true	"文件名"
// @Success		200			{object}	response.Response	"成功"
// @Router			/api/admin/common/upload/checkInstant [post]
func (uh *UploadHandler) CheckInstant(c *gin.Context) {
	var req struct {
		FileMd5  string `json:"file_md5" binding:"required"`
		FileName string `json:"file_name" binding:"required"`
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyJSON(c, &req)) {
		return
	}
	record, err := common_service.FileHashService.FindByMd5(req.FileMd5)
	if err != nil {
		core.Logger.Errorf("CheckInstant err: %v", err)
	}
	if record != nil {
		resp := common_schema.CommonUploadFileResp{
			// ID:         record.ID,
			FileHashId: record.ID,
			Name:       req.FileName,
			Uri:        util.UrlUtil.HashUrl(record.ID, req.FileName), // 访问地址（完整可访问 URL）
			// Path:       record.FilePath,                 // 磁盘存储 key
			Ext:     record.Ext,
			Size:    record.FileSize,
			Instant: true,
		}
		response.CheckAndRespWithData(c, resp, nil)
		return
	}

	// 秒传未命中：返回统一结构（instant=false），上传流程继续
	response.CheckAndRespWithData(c, common_schema.CommonUploadFileResp{Instant: false}, nil)
}

// Serve 按文件哈希ID返回文件流
//
//	@Summary		按文件哈希ID获取文件流
//	@Description	通过 id 查询 x_common_file_hash.FilePath，读取物理文件并以流形式返回
//	@Tags			common_file-文件
//	@Param			id	path	string	true	"文件哈希ID"
//	@Success		200	{file}	binary	"文件流"
//	@Router			/api/uploads/{id} [get]
func (fh *UploadHandler) Serve(c *gin.Context) {
	id := c.Param("id")
	file_name := c.Param("file_name")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	// 通过 id 查询磁盘存储 key（含扩展名）
	// 使用 singleflight 合并同一 id 的并发请求，避免缓存击穿重复打 DB

	// filePath := common_service.FileHashService.GetFilePath(id)

	res, _, _ := fh.requestGroup.Do("file:serve:"+id, func() (any, error) {
		return common_service.FileHashService.GetFilePath(id), nil
	})
	filePath, _ := res.(string)
	if filePath == "" {
		response.NotFound(c, "文件不存在")
		return
	}

	absPath := filepath.Join(config.FileConfig.UploadDirectory, filePath)
	f, err := os.Open(absPath)
	if err != nil {
		core.Logger.Errorf("FileHandler.Serve open err: id=%s path=%s err=%+v", id, absPath, err)
		response.NotFound(c, "文件不存在")
		return
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(filePath))
	ctype := mime.TypeByExtension(ext)
	if ctype == "" {
		ctype = "application/octet-stream"
	}
	// c.Header("Content-Disposition", `inline; filename="`+file_name+`"`)
	c.Header("Content-Disposition", `inline; filename="`+file_name+`"; filename*=UTF-8''`+url.PathEscape(file_name))
	c.Header("Content-Type", ctype)
	c.Header("Cache-Control", "public, max-age=31536000, immutable")
	http.ServeContent(c.Writer, c.Request, filepath.Base(absPath), time.Time{}, f)
}
