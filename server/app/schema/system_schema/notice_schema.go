package system_schema

import "github.com/adtkcn/x_null"

// SystemNoticeListReq 通知列表参数
type SystemNoticeListReq struct {
	Type   string `form:"type"`   // 通知类型筛选
	IsRead int    `form:"isRead"` // 0=未读 1=已读 -1=全部
}

// SystemNoticeDetailReq 通知详情参数
type SystemNoticeDetailReq struct {
	ID string `form:"id" binding:"required"` // 主键
}

// SystemNoticeReadReq 标记已读参数
type SystemNoticeReadReq struct {
	ID string `form:"id" binding:"required"` // 主键
}

// SystemNoticeDelReq 删除通知参数
type SystemNoticeDelReq struct {
	ID string `form:"id" binding:"required"` // 主键
}

// SystemNoticeResp 通知返回信息
type SystemNoticeResp struct {
	ID         string      `json:"id"`         // 主键
	Type       string      `json:"type"`       // 通知类型
	Title      string      `json:"title"`      // 标题
	Content    string      `json:"content"`    // 正文
	ReceiverID string      `json:"receiverId"` // 接收人ID
	SenderID   string      `json:"senderId"`   // 发送人ID
	URL        string      `json:"url"`        // 跳转URL
	IsRead     uint8       `json:"isRead"`     // 0未读 1已读
	ReadTime   x_null.Time `json:"readTime"`   // 阅读时间
	CreateTime x_null.Time `json:"createTime"` // 创建时间
}

// SystemNoticeUnreadCountResp 未读数量响应
type SystemNoticeUnreadCountResp struct {
	Count int64 `json:"count"` // 未读数量
}

// SystemNoticeSettingResp 通知偏好响应
type SystemNoticeSettingResp struct {
	SiteEnabled  uint8 `json:"siteEnabled"`  // 站内信开关
	EmailEnabled uint8 `json:"emailEnabled"` // 邮件开关
}

// SystemNoticeSettingSaveReq 保存通知偏好参数
type SystemNoticeSettingSaveReq struct {
	SiteEnabled  uint8 `form:"siteEnabled" binding:"oneof=0 1"`  // 站内信开关
	EmailEnabled uint8 `form:"emailEnabled" binding:"oneof=0 1"` // 邮件开关
}
