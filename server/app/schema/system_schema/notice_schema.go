package system_schema

import "github.com/adtkcn/x_null"

// SystemNoticeListReq 通知列表参数
type SystemNoticeListReq struct {
	Type   string `json:"type" form:"type"`       // 通知类型筛选
	IsRead int    `json:"is_read" form:"is_read"` // 0=未读 1=已读 -1=全部
}

// SystemNoticeDetailReq 通知详情参数
type SystemNoticeDetailReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemNoticeReadReq 标记已读参数
type SystemNoticeReadReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemNoticeDelReq 删除通知参数
type SystemNoticeDelReq struct {
	ID string `json:"id" form:"id" binding:"required"` // 主键
}

// SystemNoticeResp 通知返回信息
type SystemNoticeResp struct {
	ID         string      `json:"id"`          // 主键
	Type       string      `json:"type"`        // 通知类型
	Title      string      `json:"title"`       // 标题
	Content    string      `json:"content"`     // 正文
	ReceiverID string      `json:"receiver_id"` // 接收人ID
	SenderID   string      `json:"sender_id"`   // 发送人ID
	URL        string      `json:"url"`         // 跳转URL
	IsRead     uint8       `json:"is_read"`     // 0未读 1已读
	ReadTime   x_null.Time `json:"read_time"`   // 阅读时间
	CreateTime x_null.Time `json:"create_time"` // 创建时间
}

// SystemNoticeUnreadCountResp 未读数量响应
type SystemNoticeUnreadCountResp struct {
	Count int64 `json:"count"` // 未读数量
}

// NoticeChannelSetting 单个渠道的当前开关状态
type NoticeChannelSetting struct {
	Key     string `json:"key"`     // 渠道标识：site/email/app 等
	Label   string `json:"label"`   // 渠道展示名称
	Enabled uint8  `json:"enabled"` // 当前开关：0关闭 1开启
}

// SystemNoticeSettingResp 通知偏好响应
// Channels 为后端定义的渠道清单（含当前开关状态），前端按此渲染
type SystemNoticeSettingResp struct {
	Channels []NoticeChannelSetting `json:"channels"`
}

// SystemNoticeSettingSaveReq 保存通知偏好参数：channel(渠道) -> is_enabled(开关)
type SystemNoticeSettingSaveReq struct {
	Settings map[string]uint8 `json:"settings" form:"settings" binding:"required"` // 渠道开关映射
}
