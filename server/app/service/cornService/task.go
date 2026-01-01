package cornService

import "x_admin/core"

// 定义任务结构体
type Task struct {
	TaskCode string // 任务编码
	TaskDesc string // 任务描述
	TaskFunc func() // 任务函数
}

var TaskInfoList = []Task{
	{
		TaskCode: "exampleTask",
		TaskDesc: "这是一个示例任务，每分钟执行一次",
		TaskFunc: func() {
			core.Logger.Info("执行示例任务: exampleTask")
		},
	},
}

// 任务结构
type RunTask struct {
	TaskId   string // 任务ID
	TaskName string // 任务名称
	TaskCode string // 任务编码
	CronExpr string // cron表达式

	Status bool  // 是否禁用
	Task   *Task // 任务信息
}
