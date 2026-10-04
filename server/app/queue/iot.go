package queue

import (
	"context"
	"x_admin/app/iot/pipeline"
	"x_admin/config"
	"x_admin/core"
)

// Start 拉起处理流水线：接入与处理完全解耦，ctx 取消后 worker 优雅退出
func IotStart(ctx context.Context, body []byte) error {
	if !config.IotConfig.Enable {
		core.Logger.Infof("[iot处理] IOT.Enable=false，跳过流水线启动")
		return nil
	}
	// var payload email.EmailOptions
	// if err := json.Unmarshal(body, &payload); err != nil {
	// 	return err
	// }
	return pipeline.HandleIngest(ctx, body)
}
