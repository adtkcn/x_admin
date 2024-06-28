package model

import (
	"x_admin/core"

	"gorm.io/plugin/soft_delete"
)

// FlowTemplate 流程模板实体
type FlowTemplate struct {
	Id                  int    `gorm:"primarykey;comment:''"` //
	FlowName            string `gorm:"comment:'流程名称'"`        // 流程名称
	FlowGroup           int    `gorm:"comment:'流程分类'"`        // 流程分类
	FlowRemark          string `gorm:"comment:'流程描述'"`        // 流程描述
	FlowFormData        string `gorm:"comment:'表单配置'"`        // 表单配置
	FlowProcessData     string `gorm:"comment:'流程配置'"`        // 流程配置
	FlowProcessDataList string `gorm:"comment:'流程配置list数据'"`  // 流程配置list数据

	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	UpdateTime core.TsTime           `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	CreateTime core.TsTime           `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	DeleteTime core.TsTime           `gorm:"default:null;comment:'删除时间'"`   // 删除时间
}
