package model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// FlowTemplate 流程模板实体
type FlowTemplate struct {
	Id                  string `gorm:"primarykey;type:char(36);comment:''"` //
	FlowName            string `gorm:"comment:'流程名称'"`                      // 流程名称
	FlowGroup           int    `gorm:"comment:'流程分类'"`                      // 流程分类
	FlowRemark          string `gorm:"comment:'流程描述'"`                      // 流程描述
	FlowFormData        string `gorm:"comment:'表单配置'"`                      // 表单配置
	FlowProcessData     string `gorm:"comment:'流程配置'"`                      // 流程配置
	FlowProcessDataList string `gorm:"comment:'流程配置list数据'"`                // 流程配置list数据

	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	CreateTime x_null.Time           `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'"`   // 删除时间
}

// 自动在创建时设置 UUIDv7
func (u *FlowTemplate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.Id = id.String()
	return nil
}
