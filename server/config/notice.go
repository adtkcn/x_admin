package config

// NoticeChannel 通知渠道定义（由后端定义，前端按此渲染）
type NoticeChannel struct {
	Key            string // 渠道标识：site/email/app 等
	Label          string // 渠道展示名称
	DefaultEnabled uint8  // 默认开关：0关闭 1开启
}

// NoticeConfigStruct 消息通知配置
type NoticeConfigStruct struct {
	EmailDelaySeconds int             // 邮件延迟推送秒数：WebSocket推送后若未读，延迟N秒后补发邮件，默认120
	Channels          []NoticeChannel // 通知渠道清单
}

// DefaultNoticeChannels 默认渠道清单
var DefaultNoticeChannels = []NoticeChannel{
	// {Key: "site", Label: "站内消息（无用）", DefaultEnabled: 1},
	{Key: "email", Label: "邮件通知", DefaultEnabled: 1},
	// {Key: "app", Label: "APP推送（暂无）", DefaultEnabled: 1},
}

var NoticeConfig = NoticeConfigStruct{
	EmailDelaySeconds: 120,
	Channels:          DefaultNoticeChannels,
}
