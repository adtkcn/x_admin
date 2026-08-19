package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"x_admin/app/middleware"
	"x_admin/app/model/system_model"
	"x_admin/app/schema/flow_schema"
	"x_admin/app/service/flow_service"
	"x_admin/app/service/notice_service"
	"x_admin/core"
	"x_admin/util"
)

// Start 统一注册并启动所有消费者（各自队列/并发），ctx 取消后优雅退出。
func Start(ctx context.Context) {
	// 流程邮件通知：队列投递为 flow_schema.FlowNotifyEmailPayload
	// emailSvc := flow_service.NewFlowHistoryService()
	core.Queue.Consume(ctx, "flow_notify_email", func(ctx context.Context, body []byte) error {
		var payload util.EmailOptions
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		if err := util.EmailUtil.SendEmail(payload); err != nil {
			return fmt.Errorf("发送邮件失败: %w", err)
		}
		return nil

	}, 3)

	// 流程 Webhook 通知：队列投递为 flow_schema.FlowNotifyWebhookPayload
	webhookSvc := flow_service.NewFlowHistoryService()
	core.Queue.Consume(ctx, "flow_notify_webhook", func(ctx context.Context, body []byte) error {
		var payload flow_schema.FlowNotifyWebhookPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		return webhookSvc.ProcessFlowWebhook(payload)
	}, 3)

	// 邮箱验证码发送：队列投递为 util.EmailCodeTask
	core.Queue.Consume(ctx, util.QueueEmailCode, func(ctx context.Context, body []byte) error {
		var task util.EmailCodeTask
		if err := json.Unmarshal(body, &task); err != nil {
			return err
		}
		// 验证码有时效性：入队超过 60 秒（如队列卡死/积压）则丢弃，不再发送
		if task.CreatedAt > 0 && time.Now().Unix()-task.CreatedAt > 60 {
			core.Logger.Warnf("邮箱验证码任务已过期丢弃: to=%s age=%ds", task.To, time.Now().Unix()-task.CreatedAt)
			return nil
		}
		return util.EmailUtil.SendEmail(util.EmailOptions{
			To:       []string{task.To},
			Subject:  task.Subject,
			HTMLBody: task.HTMLBody,
		})
	}, 3)

	// 通知邮件补推：队列投递为 notice_service.NoticeEmailTask
	core.Queue.Consume(ctx, notice_service.QueueNoticeEmail, func(ctx context.Context, body []byte) error {
		var task notice_service.NoticeEmailTask
		if err := json.Unmarshal(body, &task); err != nil {
			return err
		}
		// // 入队超过 60 秒（如队列卡死/积压）则丢弃，不再发送
		// if task.CreatedAt > 0 && time.Now().Unix()-task.CreatedAt > 60 {
		// 	core.Logger.Warnf("通知邮件任务已过期丢弃: to=%s age=%ds", task.To, time.Now().Unix()-task.CreatedAt)
		// 	return nil
		// }
		sendErr := util.EmailUtil.SendEmail(util.EmailOptions{
			To:       []string{task.To},
			Subject:  task.Subject,
			HTMLBody: task.HTMLBody,
		})
		// 无论成功失败都回写最终状态：2发送成功 / 3发送失败
		status := notice_service.EmailStatusSuccess
		if sendErr != nil {
			status = notice_service.EmailStatusFailed
			core.Logger.Errorf("发送通知邮件失败: to=%s err=%v", task.To, sendErr)
		}
		if err := notice_service.NoticeService.MarkEmailStatus(task.NoticeIDs, status); err != nil {
			core.Logger.Errorf("回写通知邮件状态失败: to=%s err=%v", task.To, err)
		}
		// 不返回 sendErr：避免队列无限重试，最终状态已落库
		return nil
	}, 3)

	// 操作日志落库：由中间件投递，消费者在此写库，解耦请求与 DB 写入
	core.Queue.Consume(ctx, middleware.QueueOperateLog, func(ctx context.Context, body []byte) error {
		var payload middleware.OperateLogPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			return err
		}
		log := system_model.SystemLogOperate{
			AdminId:   payload.AdminId,
			Type:      payload.Type,
			Title:     payload.Title,
			Ip:        payload.Ip,
			Url:       payload.Url,
			Method:    payload.Method,
			Args:      payload.Args,
			Error:     payload.Error,
			Status:    payload.Status,
			StartTime: util.NullTimeUtil.ParseTime(payload.StartTime),
			EndTime:   util.NullTimeUtil.ParseTime(payload.EndTime),
			TaskTime:  payload.TaskTime,
		}
		if err := core.GetDB().Create(&log).Error; err != nil {
			core.Logger.Errorf("写入操作日志失败: %v", err)
			return err
		}
		return nil
	}, 3)
}
