package queue

import (
	"context"
	"x_admin/app/iot/pipeline"
	"x_admin/config"
	"x_admin/core"
)

// IotStart iot:ingest 队列消费回调：逐条报文交给处理流水线，接入与处理完全解耦
func IotStart(ctx context.Context, body []byte) error {
	if !config.IotConfig.Enable {
		core.Logger.Infof("[iot处理] IOT.Enable=false，跳过流水线启动")
		return nil
	}
	return pipeline.HandleIngest(ctx, body)
}
