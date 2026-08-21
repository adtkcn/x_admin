package queue

import (
	"context"
	"encoding/json/v2"

	"x_admin/app/schema/queue_schema"
	"x_admin/app/service/notice_service"
	"x_admin/core"
	"x_admin/util"
)

// ProcessNoticeEmail 消费通知邮件补推任务（notice:email）。
// 将 PushUserEmail 标记的「发送中」(1) 回写为最终状态：2 发送成功 / 3 发送失败，
// 状态 1 -> 2/3；无论成败都不返回 error，避免队列无限重试。
func ProcessNoticeEmail(ctx context.Context, body []byte) error {
	var task queue_schema.NoticeEmailTask
	if err := json.Unmarshal(body, &task); err != nil {
		return err
	}
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
}
