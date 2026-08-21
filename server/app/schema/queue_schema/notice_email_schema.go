package queue_schema

// NoticeEmailTask 通知邮件补推任务载荷（推入队列异步发送）
type NoticeEmailTask struct {
	To        string   `json:"to"`
	Subject   string   `json:"subject"`
	HTMLBody  string   `json:"html_body"`
	NoticeIDs []string `json:"notice_ids"` // 发送成功后标记 is_emailed 的通知ID
	CreatedAt int64    `json:"created_at"` // 入队时间戳（秒），用于过期判定
}

// NoticeEmailDelayTask 通知邮件延迟补推任务载荷（推入延迟队列）。
// 到达 EmailDelaySeconds 后才被 ConsumeDelayed 取出，再由消费者合成邮件推入 notice:email 即时队列。
type NoticeEmailDelayTask struct {
	ReceiverID   string `json:"receiver_id"`   // 接收人ID（消费者据此查邮箱并合成邮件）
	PreviewCount int    `json:"preview_count"` // 邮件正文展示的最新未读条数
}
