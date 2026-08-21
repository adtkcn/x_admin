package queue

import (
	"context"
	"encoding/json/v2"

	"x_admin/app/schema/queue_schema"
	"x_admin/app/service/notice_service"
	"x_admin/core"
)

// ProcessNoticeEmailDelay 消费通知邮件延迟补推任务。
// 消息到期（延迟 EmailDelaySeconds）后取出，依 ReceiverID 查邮箱并合成邮件，
// 推入 notice:email 即时队列异步发送（发送状态回写由 ProcessNoticeEmail 负责）。
// 「用户已读则不发」的判定放在 PushUserEmail 内（查询 is_read=0），
// 因此本消费者无需额外判断：用户若在延迟期内已读，PushUserEmail 查询为空、直接跳过。
func ProcessNoticeEmailDelay(ctx context.Context, body []byte) error {
	var task queue_schema.NoticeEmailDelayTask
	if err := json.Unmarshal(body, &task); err != nil {
		return err
	}
	if task.ReceiverID == "" {
		return nil
	}
	admins, err := notice_service.NoticeService.GetAdminsByIDs([]string{task.ReceiverID})
	if err != nil {
		core.Logger.Errorf("通知邮件延迟补推查询管理员失败: receiverID=%s err=%v", task.ReceiverID, err)
		return nil // 不重试，避免延迟队列反复堆积
	}
	admin, ok := admins[task.ReceiverID]
	if !ok || admin.Email == "" {
		return nil
	}
	notice_service.NoticeService.PushUserEmail(core.GetDB(), task.ReceiverID, admin.Email, task.PreviewCount)
	return nil
}
