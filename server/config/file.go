package config

type FileConfigStruct struct {
	UploadDirectory string `mapstructure:"UploadDirectory"` // 文件目录
	PublicPrefix    string `mapstructure:"PublicPrefix"`    // 资源访问前缀

	UploadImageSize int64 `mapstructure:"UploadImageSize"` // 上传图片大小限制
	UploadVideoSize int64 `mapstructure:"UploadVideoSize"` // 上传视频大小限制
	UploadFileSize  int64 `mapstructure:"UploadFileSize"`  // 上传文件大小限制

	UploadImageExt []string `mapstructure:"UploadImageExt"` // 上传图片扩展
	UploadVideoExt []string `mapstructure:"UploadVideoExt"` // 上传视频扩展
	UploadFileExt  []string `mapstructure:"UploadFileExt"`  // 上传文件扩展

	// 分片上传配置
	ChunkSize       int64  `mapstructure:"ChunkSize"`       // 分片大小（字节），默认 5MB
	ChunkTmpDir     string `mapstructure:"ChunkTmpDir"`     // 本地分片临时目录，默认 "./uploads/.tmp"
	ChunkExpireHour int    `mapstructure:"ChunkExpireHour"` // 未完成分片过期时间（小时），默认 24
	// 预签名配置
	PresignSecret string `mapstructure:"PresignSecret"` // 预签名 URL 密钥（留空则自动生成随机密钥）
	PresignExpire int    `mapstructure:"PresignExpire"` // 预签名有效期（秒），默认 3600
}

// var uploadImageExtDefault = []string{"png", "jpg", "jpeg", "gif", "ico", "bmp", "webp", "avif"}

var FileConfig = FileConfigStruct{
	// 资源访问前缀
	PublicPrefix: "/api/uploads", // /api/uploads
	// 上传文件路径
	UploadDirectory: "/tmp/uploads/x_admin_go/",
	UploadImageSize: 20 * 1024 * 1024,   // 20MB
	UploadVideoSize: 2000 * 1024 * 1024, // 2000MB
	UploadFileSize:  1024 * 1024 * 1024, //1GB
	// 上传图片扩展
	UploadImageExt: []string{"png", "jpg", "jpeg", "gif", "ico", "bmp", "webp", "avif"},
	// 上传音视频扩展
	UploadVideoExt: []string{"mp4", "avi", "flv", "wmv", "rmvb", "mov", "mp3", "wav", "flac", "m4a"},

	UploadFileExt: []string{"pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx", "zip", "rar", "7z", "txt"},

	// 分片上传配置
	ChunkSize:       5 * 1024 * 1024, // 5MB
	ChunkTmpDir:     "./uploads/.tmp",
	ChunkExpireHour: 24,
	// 预签名配置（Secret 留空则自动生成随机密钥）
	PresignSecret: "",
	PresignExpire: 3600,
}
