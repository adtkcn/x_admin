package queue_schema

// FlowNotifyEmailPayload 流程邮件通知异步任务载荷
type FlowNotifyEmailPayload struct {
	To       []string `json:"to"`        // 收件人邮箱列表
	Subject  string   `json:"subject"`   // 邮件主题
	HTMLBody string   `json:"html_body"` // 邮件 HTML 正文
}

// FlowNotifyWebhookPayload 流程 Webhook 回调异步任务载荷
type FlowNotifyWebhookPayload struct {
	URL     string `json:"url"`     // 回调地址
	Content string `json:"content"` // 回调内容
}
