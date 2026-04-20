package model

import (
	"github.com/adtkcn/x_null"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// FlowHistory 流程历史实体
type FlowHistory struct {
	Id string `gorm:"primarykey;type:char(36);comment:'历史id'"` // 历史id

	ApplyId string `gorm:"comment:'申请id'"` // 申请id

	TemplateId string `gorm:"comment:'模板id'"` // 模板id

	ApplyUserId string `gorm:"comment:'申请人id'"` // 申请人id

	ApplyUserNickname string `gorm:"comment:'申请人昵称'"` // 申请人昵称

	ApproverId string `gorm:"comment:'审批人id'"` // 审批人id

	ApproverNickname string `gorm:"comment:'审批用户昵称'"` // 审批用户昵称

	NodeId    string `gorm:"comment:'节点'"`   // 节点
	NodeType  string `gorm:"comment:'节点类型'"` // 节点类型
	NodeLabel string `gorm:"comment:'节点名称'"` //节点名称

	FormValue string `gorm:"comment:'表单值'"` // 表单值

	PassStatus int `gorm:"comment:'通过状态：1待处理，2通过，3拒绝'"` // 通过状态：1待处理，2通过，3拒绝

	PassRemark string `gorm:"comment:'通过备注'"` // 通过备注

	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除: 0=否, 1=是'"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;comment:'更新时间'"` // 更新时间
	CreateTime x_null.Time           `gorm:"autoCreateTime;comment:'创建时间'"` // 创建时间
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'"`   // 删除时间

}

// 自动在创建时设置 UUIDv7
func (u *FlowHistory) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.Id = id.String()
	return nil
}
