package plugin

import "io"

// MultipartPart S3 分片上传的部件信息
type MultipartPart struct {
	PartNumber int    `json:"partNumber"`
	ETag       string `json:"etag"`
}

// StorageEngine 存储引擎接口（本地文件存储）
type StorageEngine interface {
	// PutObject 上传完整对象（非分片）
	PutObject(key string, data io.Reader, size int64) (string, error)

	// ObjectExists 检查对象是否存在（用于秒传）
	ObjectExists(key string) (bool, error)

	// GetObjectURL 获取对象访问 URL
	GetObjectURL(key string) (string, error)

	// ---- 分片上传相关 ----

	// InitMultipartUpload 初始化分片上传，返回 uploadId
	InitMultipartUpload(key string) (string, error)

	// UploadPart 上传单个分片，返回 etag
	UploadPart(key string, uploadId string, partNumber int, data io.Reader, size int64) (string, error)

	// CompleteMultipartUpload 完成分片合并
	CompleteMultipartUpload(key string, uploadId string, parts []MultipartPart) (string, error)

	// AbortMultipartUpload 中止分片上传，清理服务端分片数据
	AbortMultipartUpload(key string, uploadId string) error

	// ListParts 列出已上传的分片（用于断点续传）
	ListParts(key string, uploadId string) ([]MultipartPart, error)
}
