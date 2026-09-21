package queue

import (
	"context"
	"encoding/json/v2"
	"fmt"

	"x_admin/plugin/email"
)

// ProcessFlowNotifyEmail 消费流程邮件通知任务：解析载荷后发送邮件。
func ProcessFlowNotifyEmail(ctx context.Context, body []byte) error {
	var payload email.EmailOptions
	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}
	if err := email.SendEmail(payload); err != nil {
		return fmt.Errorf("发送邮件失败: %w", err)
	}
	return nil
}
