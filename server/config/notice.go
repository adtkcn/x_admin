package config

// NoticeConfigStruct 消息通知配置
type NoticeConfigStruct struct {
	EmailDelaySeconds int // 邮件延迟推送秒数：WebSocket推送后若未读，延迟N秒后补发邮件，默认60
	AppDelaySeconds   int // App推送延迟秒数，默认120
	CleanDays         int // 已读通知保留天数，默认30
	MaxUnread         int // 单用户最大未读数，超限自动标已读，默认200
}

var NoticeConfig = NoticeConfigStruct{
	EmailDelaySeconds: 60,
	AppDelaySeconds:   120,
	CleanDays:         30,
	MaxUnread:         200,
}
