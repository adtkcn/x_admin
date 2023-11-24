package storage

//SettingStorageDetailReq 存储详情参数
type SettingStorageDetailReq struct {
	Alias string `form:"alias" binding:"required,oneof=local qiniu qcloud aliyun"` // 别名: [local,qiniu,qcloud,aliyun]
}

//SettingStorageEditReq 存储编辑参数
type SettingStorageEditReq struct {
	Alias     string `form:"alias" binding:"required,oneof=local qiniu qcloud aliyun"` // 别名: [local,qiniu,qcloud,aliyun]
	Status    int    `form:"status" binding:"oneof=0 1"`                               // 状态: 0/1
	Bucket    string `form:"bucket"`                                                   // 存储空间名
	SecretKey string `form:"secretKey"`                                                // SK
	AccessKey string `form:"accessKey"`                                                // AK
	Domain    string `form:"domain"`                                                   // 访问域名
	Region    string `form:"region"`                                                   // 地区,腾讯存储特有
}

//SettingStorageChangeReq 存储切换参数
type SettingStorageChangeReq struct {
	Alias  string `form:"alias" binding:"required,oneof=local qiniu qcloud aliyun"` // 别名: [local,qiniu,qcloud,aliyun]
	Status int    `form:"status" binding:"oneof=0 1"`                               // 状态: 0/1
}
