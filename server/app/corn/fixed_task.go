package corn

import (
	"time"
	"x_admin/app/service/cornService"
	"x_admin/app/service/monitorService"
	"x_admin/core"
	"x_admin/util"
)

// robfig/cron 基础使用示例
//
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
var FixedTasks = NewCronManager()

func init() {
	// 固定任务管理器

	FixedTasks.Start()

	// 每10秒执行一次拉取定时任务"*/10 * * * * *"
	FixedTasks.AddTask("loadTasks", "0 * * * * *", cornService.Task{
		Lock: false,
		// LockTTL:  10 * time.Second,
		TaskCode: "loadTasks",
		TaskDesc: "拉取定时任务",
		TaskFunc: func() {
			RunTaskList := loadTasks()
			core.Logger.Debug("拉取到的任务数量: ", len(RunTaskList))
			if err := DynamicTasks.AddTasksBeforeRemoveAll(RunTaskList); err != nil {
				core.Logger.Error("添加任务失败", err)
			}
		},
	})

	// 每5秒执行一次广播当前在线用户数
	FixedTasks.AddTask("onlineCount", "*/5 * * * * *", cornService.Task{
		Lock:     true,
		LockTTL:  2 * time.Second,
		TaskCode: "onlineCount",
		TaskDesc: "广播当前在线用户数",
		TaskFunc: func() {
			// 存入redis
			util.RedisUtil.RPush("onlineCount", []any{core.Ws.GetOnlineCount()}, 10)

			// 广播当前在线用户数
			core.Ws.SendToAll(map[string]any{
				"onlineCount": core.Ws.GetOnlineCount(),
			})
		},
	})

	// 每2秒执行一次收集服务器信息并推送到Redis
	FixedTasks.AddTask("CollectAndPushServerInfo", "*/2 * * * * *", cornService.Task{
		Lock:     false,
		LockTTL:  2 * time.Second,
		TaskCode: "CollectAndPushServerInfo",
		TaskDesc: "收集服务器信息并推送到Redis",
		TaskFunc: func() {
			if err := monitorService.MonitorServerService.CollectAndPushServerInfo(); err != nil {
				core.Logger.Error("收集服务器信息并推送到Redis失败", err)
			}
		},
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
