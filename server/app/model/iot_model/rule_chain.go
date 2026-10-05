package iot_model

import (
	"uuid"

	"github.com/adtkcn/x_null"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// IotRuleChain 可视化规则链模板（替代产品级三段 goja 脚本的解码/过滤/告警配置）。
//
// 独立成表、可被多个产品复用（对齐 x_flow_template 惯例）：graph_data 存前端
// LogicFlow 原生图 {nodes, edges}（不拍成单父树，保留带关系标签的多出口能力），
// 处理层按有向图逐报文执行；产品通过 rule_chain_id 引用一条链。
type IotRuleChain struct {
	ID         string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
	Name       string                `gorm:"type:varchar(100);not null;default:'';comment:'规则链名称'" json:"name"`
	Remark     string                `gorm:"type:varchar(255);not null;default:'';comment:'备注'" json:"remark"`
	GraphData  string                `gorm:"type:text;comment:'规则链图JSON(LogicFlow nodes+edges)'" json:"graph_data"`
	IsStop     uint8                 `gorm:"not null;default:0;comment:'是否停用: 0=否, 1=是'" json:"is_stop"`
	IsDelete   soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
	CreateTime x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
	UpdateTime x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
	DeleteTime x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (IotRuleChain) TableName() string { return "x_iot_rule_chain" }

func (m *IotRuleChain) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewV7().String()
	return nil
}
