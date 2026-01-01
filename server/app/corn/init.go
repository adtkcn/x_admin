package corn

import (
	"x_admin/app/schema"
	"x_admin/app/service/cornService"
	"x_admin/core"
	"x_admin/util"
)

// robfig/cron 基础使用示例
//	func init() {
//		c := cron.New(cron.WithSeconds())
//		c.AddFunc("*/5 * * * * *", func() {
//			fmt.Println("定时任务：每5秒执行一次")
//			core.Ws.SendToRoom("room1", []byte("hello room1"))
//			core.Ws.SendToAll([]byte("hello all"))
//		})
//		// 启动定时任务
//		c.Start()
//	}

// 从数据库加载任务
func loadTasks() []cornService.RunTask {
	var Status = core.NullInt{}
	Status.SetValue(1)
	allList, err := cornService.SystemCornService.ListAll(schema.SystemCornListReq{
		Status: Status,
	})
	if err != nil {
		core.Logger.Error("加载任务失败", err)
		return nil
	}

	var RunTaskList = []cornService.RunTask{} // 运行中的任务列表

	for _, task := range allList {
		if task.Status.ValueOrZero() == 0 {
			continue
		}
		for _, info := range cornService.TaskInfoList {
			if task.TaskCode.ValueOrZero() == info.TaskCode {

				RunTaskList = append(RunTaskList, cornService.RunTask{
					TaskId:   task.Id,
					TaskName: task.TaskName.ValueOrZero(),
					TaskCode: task.TaskCode.ValueOrZero(),
					CronExpr: task.CornExpr.ValueOrZero(),
					Status:   task.Status.ValueOrZero() == 1,
					Task:     &info,
				})
				break
			}
		}
	}
	return RunTaskList
}

func init() {
	// 动态任务管理器
	DynamicTasks := NewCronManager()
	DynamicTasks.Start()

	// 固定任务管理器
	FixedTasks := NewCronManager()
	FixedTasks.Start()

	// 每10秒执行一次拉取定时任务
	FixedTasks.AddTask("loadTasks", "*/10 * * * * *", func() {
		RunTaskList := loadTasks()
		core.Logger.Info("拉取到的任务数量: ", len(RunTaskList))
		if err := DynamicTasks.AddTasksBeforeRemoveAll(RunTaskList); err != nil {
			core.Logger.Error("添加任务失败", err)
		}
	})
	// 每5秒执行一次广播当前在线用户数
	FixedTasks.AddTask("onlineCount", "*/5 * * * * *", func() {
		// core.Ws.SendToRoom("room1", map[string]any{
		// 	"message": "Hello Room1!",
		// })
		// 存入redis
		util.RedisUtil.RPush("onlineCount", []any{core.Ws.GetOnlineCount()}, 10)

		// 广播当前在线用户数
		core.Ws.SendToAll(map[string]any{
			"onlineCount": core.Ws.GetOnlineCount(),
		})
	})

	// FixedTasks.AddTask("WriteInfluxdb2", "*/10 * * * * *", func() {
	// 	var alarm_event_list = []map[string]any{
	// 		{
	// 			"tid":   "284",
	// 			"site":  "4c",
	// 			"grade": "1",

	// 			"channel":    strconv.Itoa(util.ToolsUtil.Random(1, 16)),
	// 			"type":       util.ToolsUtil.Random(1, 7), // 告警类型1-7
	// 			"start_time": time.Now().Unix() - int64(util.ToolsUtil.Random(1, 20)),
	// 			"end_time":   time.Now().Unix(),
	// 			"max":        util.ToolsUtil.Random(100, 200),
	// 			"min":        util.ToolsUtil.Random(20, 100),
	// 		},
	// 	}
	// 	core.WriteInfluxdb2(alarm_event_list)
	// })
}
