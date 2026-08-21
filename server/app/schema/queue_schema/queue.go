package queue_schema

// 队列名常量：producer（各 service 投递）与 consumer（task 包注册）共用，
// 避免队列名散落在 util/middleware/notice_service 等多处。
const (
	QueueFlowNotifyEmail   = "flow_notify_email"
	QueueFlowNotifyWebhook = "flow_notify_webhook"
	QueueEmailCode         = "email:code:send"
	QueueNoticeEmail       = "notice:email"
	QueueImageWebp         = "image_webp"
	QueueOperateLog        = "operate_log"
)
