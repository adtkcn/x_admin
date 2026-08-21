package queue

import (
	"context"
	"encoding/json/v2"
	"time"

	"x_admin/core"
	"x_admin/util"
)

// ProcessEmailCode 消费邮箱验证码发送任务。
// 验证码有时效性：入队超过 60 秒（如队列卡死/积压）则丢弃，不再发送。
func ProcessEmailCode(ctx context.Context, body []byte) error {
	var task util.EmailCodeTask
	if err := json.Unmarshal(body, &task); err != nil {
		return err
	}
	if task.CreatedAt > 0 && time.Now().Unix()-task.CreatedAt > 60 {
		core.Logger.Warnf("邮箱验证码任务已过期丢弃: to=%s age=%ds", task.To, time.Now().Unix()-task.CreatedAt)
		return nil
	}
	return util.EmailUtil.SendEmail(util.EmailOptions{
		To:       []string{task.To},
		Subject:  task.Subject,
		HTMLBody: task.HTMLBody,
	})
}
