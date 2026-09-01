package common_controller

import (
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
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
	file, err := util.VerifyUtil.VerifyFile(c, "file")
	if response.IsFailWithResp(c, err) {
		return
	}
	fileHash, err := common_service.UploadService.UploadFile(file)
	// 访问地址 = GET /api/uploads/:id，由文件流路由按 id 查 x_common_file_hash.FilePath 返回
	resp := common_schema.CommonUploadFileResp{
		// ID:         fileHash.ID,
		FileHashId: fileHash.ID,
		Name:       file.Filename,
		Uri:        util.UrlUtil.HashUrl(fileHash.ID, file.Filename), // 访问地址（完整可访问 URL）
		// Path:       fileHash.FilePath,                 // 磁盘存储 key = <id>.<ext>
		Ext:     fileHash.Ext,
		Size:    fileHash.FileSize,
		Instant: false,
	}
	response.CheckAndRespWithData(c, resp, err)

	// 压缩80%转webp
	common_service.UploadService.ConvertImage(fileHash.ID, fileHash.FilePath, 80, 0, 0)
	// 缩略图
	common_service.UploadService.ConvertImage(fileHash.ID, fileHash.FilePath, 80, 200, 200)
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
	if record.ID != "" {
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

// Serve 按文件哈希ID返回文件流,如果是图片，则返回压缩后的图片，否则返回原文件
//
//	@Summary		按文件哈希ID获取文件流
//	@Description	通过 id 查询 x_common_file_hash.FilePath，读取物理文件并以流形式返回
//	@Tags			common_file-文件
//	@Param			id				path	string	true	"文件哈希ID"
//	@Param			file_name		path	string	true	"文件名"
//	@Param			quality			query	int		false	"图片质量"
//	@Param			scale_width		query	int		false	"图片缩放宽度"
//	@Param			scale_height	query	int		false	"图片缩放高度"
//	@Success		200				{file}	binary	"文件流"
//	@Router			/api/uploads/{id} [get]
func (fh *UploadHandler) Serve(c *gin.Context) {
	id := c.Param("id")
	file_name := c.Param("file_name")
	file_name_ext := util.UrlUtil.GetFileExt(file_name)
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	// 解析派生参数：quality / scale_width / scale_height
	// 任一非零则视为请求派生版本，按 pid 查询对应派生记录；否则查询主文件。
	quality, _ := strconv.Atoi(c.Query("quality"))
	if quality == 0 {
		quality = 80
	}
	scaleWidth, _ := strconv.Atoi(c.Query("scale_width"))
	scaleHeight, _ := strconv.Atoi(c.Query("scale_height"))
	needDerived := (quality != 0 || scaleWidth != 0 || scaleHeight != 0) && util.ToolsUtil.Contains([]string{"jpg", "jpeg", "png", "webp"}, file_name_ext)

	// 使用 singleflight 合并同一 key 的并发请求，避免缓存击穿重复打 DB
	// key 纳入派生参数维度，保证不同 quality/缩放请求互不干扰
	sfKey := "file:serve:" + id
	if needDerived {
		sfKey += ":" + strconv.Itoa(quality) + ":" + strconv.Itoa(scaleWidth) + ":" + strconv.Itoa(scaleHeight)
	}
	res, _, _ := fh.requestGroup.Do(sfKey, func() (any, error) {
		if needDerived {
			filePath := common_service.FileHashService.GetDerivedPath(id, quality, scaleWidth, scaleHeight)
			// 派生文件存在，则返回
			if filePath != "" {
				return filePath, nil
			}
			// 推入异步队列，下一次请求时大概率获取到压缩文件
			common_service.UploadService.ConvertImage(id, file_name, quality, scaleWidth, scaleHeight)
		}
		//派生文件不存在，则返回主文件
		return common_service.FileHashService.GetFilePath(id), nil
	})
	filePath, _ := res.(string)
	if filePath == "" {
		response.NotFound(c, "文件不存在")
		return
	}

	// 读取文件路径
	absPath := filepath.Join(config.FileConfig.UploadDirectory, filePath)
	f, err := os.Open(absPath)
	if err != nil {
		core.Logger.Errorf("FileHandler.Serve open err: id=%s path=%s err=%+v", id, absPath, err)
		response.NotFound(c, "文件不存在")
		return
	}
	defer f.Close()

	ext := util.UrlUtil.GetFileExt(filePath)
	ctype := mime.TypeByExtension("." + ext)
	if ctype == "" {
		ctype = "application/octet-stream"
	}
	// 替换拼接真实后缀
	file_name = util.UrlUtil.ReplaceExt(file_name, "."+ext)

	// c.Header("Content-Disposition", `inline; filename="`+file_name+`"`)
	c.Header("Content-Disposition", `inline;filename*=UTF-8''`+url.PathEscape(file_name))
	c.Header("Content-Type", ctype)
	c.Header("Cache-Control", "public, max-age=31536000, immutable")
	http.ServeContent(c.Writer, c.Request, filepath.Base(absPath), time.Now(), f)
}
