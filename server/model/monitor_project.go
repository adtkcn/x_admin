package model

import "x_admin/core"

//MonitorProject 错误项目实体
type MonitorProject struct {
	Id int `gorm:"primarykey;comment:'项目id'" excel:"name:项目id;"` // 项目id

	ProjectKey string `gorm:"comment:'项目uuid'" excel:"name:项目uuid;"` // 项目uuid

	ProjectName string `gorm:"comment:'项目名称'" excel:"name:项目名称;"` // 项目名称

	ProjectType string `gorm:"comment:'项目类型go java web node php 等'" excel:"name:项目类型"` // 项目类型go java web node php 等

	IsDelete int `gorm:"comment:'是否删除: 0=否, 1=是'" excel:"name:是否删除: 0=否, 1=是;"` // 是否删除: 0=否, 1=是

	UpdateTime core.TsTime `gorm:"autoUpdateTime;comment:'更新时间'" excel:"name:更新时间;"` // 更新时间

	CreateTime core.TsTime `gorm:"autoCreateTime;comment:'创建时间'" excel:"name:创建时间;"` // 创建时间

	DeleteTime core.TsTime `gorm:"default:null;comment:'删除时间'" excel:"name:删除时间;"` // 删除时间

}
