package common_schema

// CommonFileHashAddReq 新增文件哈希记录请求
type CommonFileHashAddReq struct {
	FileMd5  string `json:"file_md5" binding:"required"`
	FileSize int64  `json:"file_size" binding:"required"`
	FilePath string `json:"file_path" binding:"required"`
	Ext      string `json:"ext"`
}

// CommonFileHashResp 文件哈希记录响应
type CommonFileHashResp struct {
	ID       string `json:"id"`
	FileMd5  string `json:"file_md5"`
	FileSize int64  `json:"file_size"`
	FilePath string `json:"file_path"`
	Ext      string `json:"ext"`
}
