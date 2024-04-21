package model

//MonitorClient 客户端信息实体
type MonitorClient struct {
	Id int `gorm:"primarykey;comment:'uuid'" excel:"name:uuid;"` // uuid

	ProjectKey string `gorm:"comment:'项目key'" excel:"name:项目key;"` // 项目key

	ClientId string `gorm:"comment:'sdk生成的客户端id'" excel:"name:sdk生成的客户端id;"` // sdk生成的客户端id

	UserId string `gorm:"comment:'用户id'" excel:"name:用户id;"` // 用户id

	Os string `gorm:"comment:'系统'" excel:"name:系统;"` // 系统

	Browser string `gorm:"comment:'浏览器'" excel:"name:浏览器;"` // 浏览器

	City string `gorm:"comment:'城市'" excel:"name:城市;"` // 城市

	Width int `gorm:"comment:'屏幕'" excel:"name:屏幕;"` // 屏幕

	Height int `gorm:"comment:'屏幕高度'" excel:"name:屏幕高度;"` // 屏幕高度

	Ua string `gorm:"comment:'ua记录'" excel:"name:ua记录;"` // ua记录

	ClientTime int64 `gorm:"comment:'客户端时间'" excel:"name:客户端时间;"` // 客户端时间

	CreateTime int64 `gorm:"autoCreateTime;comment:'创建时间'" excel:"name:创建时间;"` // 创建时间

}
