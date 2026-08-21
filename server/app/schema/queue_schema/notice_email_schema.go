package queue_schema

// NoticeEmailTask 通知邮件补推任务载荷（推入队列异步发送）
type NoticeEmailTask struct {
	To        string   `json:"to"`
	Subject   string   `json:"subject"`
	HTMLBody  string   `json:"html_body"`
	NoticeIDs []string `json:"notice_ids"` // 发送成功后标记 is_emailed 的通知ID
	CreatedAt int64    `json:"created_at"` // 入队时间戳（秒），用于过期判定
}
