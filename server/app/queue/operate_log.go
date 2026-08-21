package queue

import (
	"context"
	"encoding/json/v2"

	"x_admin/app/middleware"
	"x_admin/app/model/system_model"
	"x_admin/core"
	"x_admin/util"
)

// ProcessOperateLog 消费操作日志落库任务：由中间件投递，消费者在此写库，
// 解耦请求与 DB 写入。
func ProcessOperateLog(ctx context.Context, body []byte) error {
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
}
