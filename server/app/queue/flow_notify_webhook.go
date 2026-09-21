package queue

import (
	"context"
	"encoding/json/v2"

	"x_admin/app/schema/queue_schema"
	"x_admin/app/service/flow_service"
)

// ProcessFlowNotifyWebhook 消费流程 Webhook 通知任务：解析载荷后异步回调。
func ProcessFlowNotifyWebhook(ctx context.Context, body []byte) error {
	var payload queue_schema.FlowNotifyWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}
	webhookSvc := flow_service.NewFlowHistoryService()
	return webhookSvc.ProcessFlowWebhook(payload)
}
