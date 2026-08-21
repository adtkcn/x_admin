package queue

import (
	"context"

	"x_admin/app/schema/queue_schema"
	"x_admin/core"
)

// Start 统一注册并启动所有消费者（各自队列/并发），ctx 取消后优雅退出。
func Start(ctx context.Context) {
	// 流程邮件通知
	core.Queue.Consume(ctx, queue_schema.QueueFlowNotifyEmail, ProcessFlowNotifyEmail, 3)

	// 流程 Webhook 通知
	core.Queue.Consume(ctx, queue_schema.QueueFlowNotifyWebhook, ProcessFlowNotifyWebhook, 3)

	// 邮箱验证码发送
	core.Queue.Consume(ctx, queue_schema.QueueEmailCode, ProcessEmailCode, 3)

	// 通知邮件补推
	core.Queue.Consume(ctx, queue_schema.QueueNoticeEmail, ProcessNoticeEmail, 3)

	// 图片异步转 webp：上传 jpg/png 后投递，worker 读取原图转 webp 并回写记录
	core.Queue.Consume(ctx, queue_schema.QueueImageWebp, ProcessImageWebp, 3)

	// 操作日志落库：由中间件投递，消费者在此写库，解耦请求与 DB 写入
	core.Queue.Consume(ctx, queue_schema.QueueOperateLog, ProcessOperateLog, 3)
}
