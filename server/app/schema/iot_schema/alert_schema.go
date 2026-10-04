package iot_schema

import (
	"time"
	"x_admin/core/request"
)

// ============ 告警历史 ============

// AlertListReq 告警历史分页查询：可按设备/编号/产品/区域/状态/触发时间区间过滤
type AlertListReq struct {
	DeviceId  string `json:"device_id" form:"device_id"`                             // 设备ID
	Sn        string `json:"sn" form:"sn"`                                           // 设备编号
	ProductId string `json:"product_id" form:"product_id"`                           // 产品ID
	SpaceId   string `json:"space_id" form:"space_id"`                               // 区域ID
	Status    int8   `json:"status" form:"status,default=-1" binding:"oneof=-1 0 1"` // 状态: -1全部 0已恢复 1告警中
	Severity  int8   `json:"severity" form:"severity,default=-1"`                    // 级别: -1全部 1轻微 2严重 3故障
	StartTime string `json:"start_time" form:"start_time"`                           // 触发时间起(2006-01-02 15:04:05)
	EndTime   string `json:"end_time" form:"end_time"`                               // 触发时间止(2006-01-02 15:04:05)
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
	Severity    uint8          `json:"severity"`     // 级别: 1=轻微, 2=严重, 3=故障
	Values      map[string]any `json:"values"`       // 触发时属性快照
	TriggeredAt time.Time      `json:"triggered_at"` // 首次触发时间
	RecoveredAt string         `json:"recovered_at"` // 恢复时间(未恢复为空串)
}

// AlertDelReq 批量删除告警历史
type AlertDelReq struct {
	IDs []string `json:"ids" form:"ids" binding:"required,min=1"` // 主键集合
}
