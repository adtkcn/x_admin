package model

import (
	"x_admin/core"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// FlowApply 申请流程实体
type FlowApply struct {
	Id                  string                `gorm:"primarykey;type:char(36);comment:'申请id'"` // 申请id
	TemplateId          string                `gorm:"comment:'模板'"`                            // 模板
	ApplyUserId         int                   `gorm:"comment:'申请人id'"`                         // 申请人id
	ApplyUserNickname   string                `gorm:"comment:'申请人昵称'"`                         // 申请人昵称
	FlowName            string                `gorm:"comment:'流程名称'"`                          // 流程名称
	FlowGroup           int                   `gorm:"comment:'流程分类'"`                          // 流程分类
	FlowRemark          string                `gorm:"comment:'流程描述'"`                          // 流程描述
	FlowFormData        string                `gorm:"comment:'表单配置'"`                          // 表单配置
	FlowProcessData     string                `gorm:"comment:'流程配置'"`                          // 流程配置
	FlowProcessDataList string                `gorm:"comment:'流程配置list数据'"`                    // 流程配置list数据
	FormValue           string                `gorm:"comment:'表单值'"`                           // 表单值
	Status              int                   `gorm:"comment:'状态：1待提交，2审批中，3审批完成，4审批失败'"`      // 状态：0待提交，1审批中，2审批完成，3审批失败
	IsDelete            soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	UpdateTime          core.NullTime         `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	CreateTime          core.NullTime         `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	DeleteTime          core.NullTime         `gorm:"default:null;comment:'删除时间'"`   // 删除时间
}

// 自动在创建时设置 UUIDv7
func (u *FlowApply) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.Id = id.String()
	return nil
}
