package corn_service

import (
	"time"
	"x_admin/core"
)

// 定义任务结构体
type Task struct {
	Lock    bool          // 是否使用分布式锁
	LockTTL time.Duration // 锁过期时间

	TaskCode string // 任务编码
	TaskDesc string // 任务描述

	TaskFunc func() // 任务函数
}

// 任务结构
type RunTask struct {
	TaskId   string // 任务ID
	TaskName string // 任务名称
	CronExpr string // cron表达式
	Task     *Task  // 任务信息
}

// 定义任务列表
var TaskInfoList = []Task{
	{
		Lock:    true,
		LockTTL: 10 * time.Second,

		TaskCode: "exampleTask",
		TaskDesc: "这是一个示例任务",

		TaskFunc: func() {

			core.Logger.Debug("执行示例任务: exampleTask")
			// 模拟任务执行时间
			time.Sleep(7 * time.Second)
		},
	},
}
