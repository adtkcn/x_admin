package corn

import (
	"time"
	"x_admin/app/service/common_service"
	"x_admin/app/service/corn_service"
	"x_admin/app/service/monitor_service"
	"x_admin/app/service/notice_service"
	"x_admin/core"
	"x_admin/plugin"
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
	FixedTasks.AddTask("loadTasks", "40 * * * * *", corn_service.Task{
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
	FixedTasks.AddTask("onlineCount", "*/5 * * * * *", corn_service.Task{
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
	FixedTasks.AddTask("CollectAndPushServerInfo", "*/5 * * * * *", corn_service.Task{
		Lock:     false,
		LockTTL:  2 * time.Second,
		TaskCode: "CollectAndPushServerInfo",
		TaskDesc: "收集服务器信息并推送到Redis",
		TaskFunc: func() {
			if err := monitor_service.MonitorServerService.CollectAndPushServerInfo(); err != nil {
				core.Logger.Error("收集服务器信息并推送到Redis失败", err)
			}
		},
	})

	// 每天凌晨1点删除三个月前的错误监控数据
	FixedTasks.AddTask("DelMonitorErrorListThreeMonthAgo", "0 1 * * * *", corn_service.Task{
		Lock:     true,
		LockTTL:  10 * time.Minute,
		TaskCode: "DelMonitorErrorListThreeMonthAgo",
		TaskDesc: "删除三个月前的错误监控数据",
		TaskFunc: func() {
			if err := monitor_service.MonitorErrorListService.DelThreeMonthAgo(); err != nil {
				core.Logger.Error("删除三个月前的错误监控数据失败", err)
			}
		},
	})

	// 每60秒执行一次邮件延迟补推
	FixedTasks.AddTask("EmailDelayPush", "*/60 * * * * *", corn_service.Task{
		Lock:     true,
		LockTTL:  55 * time.Second,
		TaskCode: "EmailDelayPush",
		TaskDesc: "邮件延迟补推：扫描未读通知，对配置了邮箱的用户发送邮件提醒",
		TaskFunc: func() {
			notice_service.NoticeService.ProcessEmailDelayPush()
		},
	})

	// 每小时执行一次清理过期的分片临时目录
	FixedTasks.AddTask("CleanChunkTmpDir", "0 0 * * * *", corn_service.Task{
		Lock:     true,
		LockTTL:  10 * time.Minute,
		TaskCode: "CleanChunkTmpDir",
		TaskDesc: "清理过期的分片临时目录",
		TaskFunc: func() {
			plugin.CleanChunkTmpDir()
		},
	})

	// 每天凌晨2点清理上传超过7天且无业务引用的文件
	FixedTasks.AddTask("CleanOrphanFiles", "0 0 2 * * *", corn_service.Task{
		Lock:     true,
		LockTTL:  30 * time.Minute,
		TaskCode: "CleanOrphanFiles",
		TaskDesc: "清理超过7天未使用的文件（无业务引用）",
		TaskFunc: func() {
			common_service.FileRefService.CleanOrphanFiles(7)
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
