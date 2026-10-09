package iot_schema

import (
	"x_admin/core/request"

	"github.com/adtkcn/x_null"
)

// ============ 告警历史 ============

// AlertListReq 告警历史分页查询：可按设备/编号/产品/区域/状态/级别/触发时间区间过滤
// 查询条件统一用 x_null.String：不传/空串均视为“不筛选”（severity 的 -1 是真级别，不能当哨兵）
type AlertListReq struct {
	DeviceId  x_null.String `json:"device_id" form:"device_id"`   // 设备ID
	Sn        x_null.String `json:"sn" form:"sn"`                 // 设备编号(模糊)
	ProductId x_null.String `json:"product_id" form:"product_id"` // 产品ID
	SpaceId   x_null.String `json:"space_id" form:"space_id"`     // 区域ID(含子区域)
	Status    x_null.String `json:"status" form:"status"`         // 状态: 0已恢复 1告警中
	Severity  x_null.String `json:"severity" form:"severity"`     // 级别: -1故障 1轻微 2严重
	StartTime x_null.String `json:"start_time" form:"start_time"` // 触发时间起(2006-01-02 15:04:05)
	EndTime   x_null.String `json:"end_time" form:"end_time"`     // 触发时间止(2006-01-02 15:04:05)
	request.PageReq
}

// AlertResp 一条告警事件记录：属性快照已反序列化为 map，便于前端直接渲染
type AlertResp struct {
	ID          string         `json:"id"`           // 主键
	DeviceId    string         `json:"device_id"`    // 设备ID
	Sn          string         `json:"sn"`           // 设备编号
	ProductId   string         `json:"product_id"`   // 产品ID
	ProductName string         `json:"product_name"` // 产品名称
	SpaceId     string         `json:"space_id"`     // 区域ID
	SpaceName   string         `json:"space_name"`   // 区域名称
	Status      uint8          `json:"status"`       // 状态: 1=告警中, 0=已恢复
	Severity    int8           `json:"severity"`     // 级别: -1=故障, 1=轻微, 2=严重(数值即危害序,3+预留)
	Values      map[string]any `json:"values"`       // 触发时属性快照
	TriggeredAt x_null.Time    `json:"triggered_at"` // 首次触发时间
	RecoveredAt string         `json:"recovered_at"` // 恢复时间(未恢复为空串)
}

// AlertDelReq 批量删除告警历史
type AlertDelReq struct {
	IDs []string `json:"ids" form:"ids" binding:"required,min=1"` // 主键集合
}
