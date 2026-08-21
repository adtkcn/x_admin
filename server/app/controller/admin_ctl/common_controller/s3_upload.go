package common_controller

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"strconv"
	"strings"

	"x_admin/app/schema/common_schema"
	"x_admin/app/service/common_service"
	"x_admin/core"
	"x_admin/core/response"
	"x_admin/plugin/storage"
	"x_admin/util"

	"github.com/gin-gonic/gin"
)

// S3UploadHandler S3 标准协议控制器（文件存本地，API 对齐 S3 REST）
type S3UploadHandler struct{}

// S3Handler S3 统一入口：根据 HTTP 方法 + query 参数分发到具体操作
// 路由: ANY /api/admin/s3/*fileKey
func (h S3UploadHandler) S3Handler(c *gin.Context) {
	method := c.Request.Method
	rawKey := strings.TrimPrefix(c.Param("fileKey"), "/")

	// 去掉 S3 桶名前缀（forcePathStyle 时会带上 Bucket 名）
	parts := strings.SplitN(rawKey, "/", 2)
	var fileKey string
	if len(parts) == 2 {
		fileKey = parts[1]
	} else {
		fileKey = rawKey
	}

	// ---- 扩展端点（非 S3 标准，因 Gin wildcard 限制内聚在此） ----
	if method == "POST" && fileKey == "generateKey" {
		h.GenerateKey(c)
		return
	}
	if method == "POST" && fileKey == "checkInstant" {
		h.CheckInstant(c)
		return
	}
	if method == "POST" && fileKey == "registerHash" {
		h.RegisterHash(c)
		return
	}

	q := c.Request.URL.Query()
	hasUploads := q.Has("uploads")
	uploadId := q.Get("uploadId")
	partNumberStr := q.Get("partNumber")
	engine := storage.GetStorageEngine()

	switch {
	case method == "POST" && hasUploads:
		h.createMultipartUpload(c, engine, fileKey)
	case method == "PUT" && uploadId != "" && partNumberStr != "":
		h.uploadPart(c, engine, fileKey, uploadId, partNumberStr)
	case method == "POST" && uploadId != "":
		h.completeMultipartUpload(c, engine, fileKey, uploadId)
	case method == "DELETE" && uploadId != "":
		h.abortMultipartUpload(c, engine, fileKey, uploadId)
	case method == "GET" && uploadId != "":
		h.listParts(c, engine, fileKey, uploadId)
	case method == "PUT":
		h.putObject(c, engine, fileKey)
	case method == "HEAD":
		h.headObject(c, engine, fileKey)
	default:
		s3Error(c, 405, "MethodNotAllowed", "不支持的操作")
	}
}

// ---- 秒传检查 ----

// CheckInstant 秒传：根据文件 MD5 查询是否已上传
func (h S3UploadHandler) CheckInstant(c *gin.Context) {
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
		engine := storage.GetStorageEngine()
		url, _ := engine.GetObjectURL(record.FilePath)
		resp := common_schema.CommonUploadFileResp{
			// ID:         record.ID,
			FileHashId: record.ID,
			Name:       req.FileName,
			Uri:        url, // 访问地址（完整可访问 URL）
			// Path:       record.FilePath, // 相对路径
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

// ---- 密钥生成 ----

// GenerateKey 生成存储 key，路径结构对齐直传：年月日/时/分/{uuid}.ext
// 与 storage.buildSaveName 保持一致，使直传与 S3 分片落盘到同一目录层级。
func (h S3UploadHandler) GenerateKey(c *gin.Context) {
	var req struct {
		FileName string `json:"file_name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		s3Error(c, 400, "InvalidArgument", "参数错误: "+err.Error())
		return
	}
	key := util.UrlUtil.BuildFileSavePath(req.FileName)
	c.JSON(200, gin.H{"key": key})
}

// ---- 后置：注册文件哈希（秒传数据源） ----

// RegisterHash 上传完成后由前端调用，将 MD5 与 fileKey 关联存入哈希表。
// id 为哈希表主键，FilePath 存真实存储 key（fileKey），文件名无需与 id 一致。
func (h S3UploadHandler) RegisterHash(c *gin.Context) {
	var req struct {
		FileMd5  string `json:"file_md5" binding:"required"`
		FileSize int64  `json:"file_size"`
		FileKey  string `json:"file_key" binding:"required"`
		FileName string `json:"file_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		s3Error(c, 400, "InvalidArgument", "参数错误: "+err.Error())
		return
	}
	ext := util.UrlUtil.GetFileExt(req.FileName)
	// FilePath 存真实存储 key（fileKey），文件名无需与 id 一致；id 由 Create 内部生成
	id, err := common_service.FileHashService.Create(req.FileMd5, req.FileSize, req.FileKey, ext)
	if err != nil {
		core.Logger.Errorf("RegisterHash err: %v", err)
		response.CheckAndRespWithData(c, common_schema.CommonUploadFileResp{}, err)
		return
	}
	resp := common_schema.CommonUploadFileResp{
		// ID:         id,
		FileHashId: id,
		Name:       req.FileName,
		Uri:        util.UrlUtil.HashUrl(id), // 访问地址（完整可访问 URL）
		Ext:        ext,
		Size:       req.FileSize,
		Instant:    false,
	}
	response.CheckAndRespWithData(c, resp, nil)

	// 上传成功后异步转 webp（条件判断在 MaybeConvertWebp 内）
	common_service.UploadService.MaybeConvertWebp(id, req.FileKey, ext, req.FileSize)
}

// ---- CreateMultipartUpload ----

func (h S3UploadHandler) createMultipartUpload(c *gin.Context, engine storage.StorageEngine, fileKey string) {
	uploadId, err := engine.InitMultipartUpload(fileKey)
	if err != nil {
		core.Logger.Errorf("createMultipartUpload err: %v", err)
		s3Error(c, 500, "InternalError", "初始化分片上传失败")
		return
	}
	c.Header("Content-Type", "application/xml")
	c.String(200, xmlInitiateResult(fileKey, uploadId))
}

// ---- UploadPart ----

func (h S3UploadHandler) uploadPart(c *gin.Context, engine storage.StorageEngine, fileKey, uploadId, partNumberStr string) {
	partNumber, err := strconv.Atoi(partNumberStr)
	if err != nil || partNumber < 1 {
		s3Error(c, 400, "InvalidArgument", "分片序号错误")
		return
	}
	bodyData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		s3Error(c, 400, "InvalidRequest", "读取分片数据失败")
		return
	}
	etag, err := engine.UploadPart(fileKey, uploadId, partNumber, bytes.NewReader(bodyData), int64(len(bodyData)))
	if err != nil {
		core.Logger.Errorf("uploadPart err: %v", err)
		s3Error(c, 500, "InternalError", "上传分片失败")
		return
	}
	c.Header("ETag", fmt.Sprintf(`"%s"`, etag))
	c.Status(200)
}

// ---- CompleteMultipartUpload ----

type completeMultipartUploadXML struct {
	XMLName xml.Name `xml:"CompleteMultipartUpload"`
	Parts   []struct {
		PartNumber int    `xml:"PartNumber"`
		ETag       string `xml:"ETag"`
	} `xml:"Part"`
}

func (h S3UploadHandler) completeMultipartUpload(c *gin.Context, engine storage.StorageEngine, fileKey, uploadId string) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		s3Error(c, 400, "InvalidRequest", "读取请求体失败")
		return
	}
	var req completeMultipartUploadXML
	if err := xml.Unmarshal(bodyBytes, &req); err != nil {
		core.Logger.Errorf("completeMultipartUpload xml parse err: %v", err)
		s3Error(c, 400, "MalformedXML", "XML 格式错误")
		return
	}
	parts := make([]storage.MultipartPart, 0, len(req.Parts))
	for _, p := range req.Parts {
		etag := strings.Trim(p.ETag, `"`)
		parts = append(parts, storage.MultipartPart{PartNumber: p.PartNumber, ETag: etag})
	}

	filePath, err := engine.CompleteMultipartUpload(fileKey, uploadId, parts)
	if err != nil {
		core.Logger.Errorf("completeMultipartUpload err: %v", err)
		s3Error(c, 500, "InternalError", "合并分片失败")
		return
	}
	core.Logger.Infof("S3 CompleteMultipartUpload 完成, key=%s, filePath=%s", fileKey, filePath)
	url, _ := engine.GetObjectURL(filePath)
	c.Header("Content-Type", "application/xml")
	c.String(200, xmlCompleteResult(fileKey, filePath, url))
}

// ---- AbortMultipartUpload ----

func (h S3UploadHandler) abortMultipartUpload(c *gin.Context, engine storage.StorageEngine, fileKey, uploadId string) {
	if err := engine.AbortMultipartUpload(fileKey, uploadId); err != nil {
		core.Logger.Errorf("abortMultipartUpload err: %v", err)
	}
	c.Status(204)
}

// ---- ListParts ----

func (h S3UploadHandler) listParts(c *gin.Context, engine storage.StorageEngine, fileKey, uploadId string) {
	parts, err := engine.ListParts(fileKey, uploadId)
	if err != nil {
		core.Logger.Errorf("listParts err: %v", err)
		s3Error(c, 500, "InternalError", "查询分片列表失败")
		return
	}
	c.Header("Content-Type", "application/xml")
	c.String(200, xmlListPartsResult(fileKey, uploadId, parts))
}

// ---- PutObject ----

func (h S3UploadHandler) putObject(c *gin.Context, engine storage.StorageEngine, fileKey string) {
	bodyData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		s3Error(c, 400, "InvalidRequest", "读取文件数据失败")
		return
	}
	_, err = engine.PutObject(fileKey, bytes.NewReader(bodyData), int64(len(bodyData)))
	if err != nil {
		core.Logger.Errorf("putObject err: %v", err)
		s3Error(c, 500, "InternalError", "上传文件失败")
		return
	}
	core.Logger.Infof("S3 PutObject 完成, key=%s", fileKey)
	c.Status(200)
}

// ---- HeadObject ----

func (h S3UploadHandler) headObject(c *gin.Context, engine storage.StorageEngine, fileKey string) {
	exists, err := engine.ObjectExists(fileKey)
	if err != nil {
		core.Logger.Errorf("headObject err: %v", err)
	}
	if exists {
		c.Status(200)
	} else {
		c.Status(404)
	}
}

// ---- S3 XML 模板 ----

func xmlEscape(s string) string { return html.EscapeString(s) }

func xmlInitiateResult(key, uploadId string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<InitiateMultipartUploadResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <Bucket>files</Bucket>
  <Key>%s</Key>
  <UploadId>%s</UploadId>
</InitiateMultipartUploadResult>`, xmlEscape(key), xmlEscape(uploadId))
}

func xmlCompleteResult(key, filePath, url string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<CompleteMultipartUploadResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <Location>%s</Location>
  <Bucket>files</Bucket>
  <Key>%s</Key>
  <ETag>"%s"</ETag>
</CompleteMultipartUploadResult>`, xmlEscape(url), xmlEscape(key), xmlEscape(filePath))
}

func xmlListPartsResult(key, uploadId string, parts []storage.MultipartPart) string {
	partsXML := ""
	for _, p := range parts {
		partsXML += fmt.Sprintf("  <Part><PartNumber>%d</PartNumber><ETag>\"%s\"</ETag></Part>\n",
			p.PartNumber, xmlEscape(p.ETag))
	}
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<ListPartsResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
  <Bucket>files</Bucket>
  <Key>%s</Key>
  <UploadId>%s</UploadId>
  <IsTruncated>false</IsTruncated>
%s</ListPartsResult>`, xmlEscape(key), xmlEscape(uploadId), partsXML)
}

func xmlError(code, message string) string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Error><Code>%s</Code><Message>%s</Message></Error>`, xmlEscape(code), xmlEscape(message))
}

func s3Error(c *gin.Context, httpStatus int, code, message string) {
	c.Header("Content-Type", "application/xml")
	c.String(httpStatus, xmlError(code, message))
	c.Abort()
}
