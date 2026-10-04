package queue_schema

// 队列名常量：producer（各 service 投递）与 consumer（task 包注册）共用，
// 避免队列名散落在 util/middleware/notice_service 等多处。
const (
	QueueFlowNotifyEmail   = "flow_notify_email"
	QueueFlowNotifyWebhook = "flow_notify_webhook"
	QueueEmailCode         = "email:code:send"
	QueueNoticeEmail       = "notice:email"
	QueueNoticeEmailDelay  = "notice:email:delay" // 通知邮件延迟补推（到达延迟时间后才投递到 notice:email）
	QueueImageWebp         = "image_webp"
	QueueOperateLog        = "operate_log"

	// IoT 设备接入：接入层（HTTP/MQTT/TCP/WS）投递，iot/pipeline 消费
	// 离线判定由 app/corn 的 iotOfflineScan（弹到期表）与 iotOfflineSweep（慢路径重建）定时执行
	QueueIotIngest = "iot:ingest"
)
