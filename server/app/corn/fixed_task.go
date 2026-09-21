package corn

import (
	"strconv"
	"time"
	"x_admin/app/service/corn_service"
	"x_admin/app/service/monitor_service"
	"x_admin/core"
	"x_admin/plugin/storage"
	"x_admin/util"
)

var FixedTasks = NewCronManager()

func init() {
	// 固定任务管理器

	FixedTasks.Start()

	// 定时执行一次拉取定时任务
	FixedTasks.AddTask("loadTasks", "40 */10 * * * *", corn_service.Task{
		LockTTL:  0,
		TaskCode: "loadTasks",
		TaskDesc: "拉取定时任务",
		TaskFunc: func() {
			loadTasks()
		},
	})

	// 每5秒执行一次websocket广播当前在线用户数
	FixedTasks.AddTask("onlineCount", "*/5 * * * * *", corn_service.Task{

		LockTTL:  3,
		TaskCode: "onlineCount",
		TaskDesc: "广播当前在线用户数",
		TaskFunc: func() {
			count := core.Ws.GetOnlineCount()

			// 存入redis，保留最近1小时数据（3600/5=720条）
			// 字符串拼接格式: "15:04:05,count"，避免 JSON 序列化开销
			record := time.Now().Format("15:04:05") + "," + strconv.Itoa(count)
			util.RedisUtil.RPush("onlineCount", []any{record}, 720)

			// 广播当前在线用户数
			core.Ws.SendToAll("onlineCount", map[string]any{
				"count": count,
			})
		},
	})

	// 每3秒执行一次收集服务器信息并推送到Redis
	FixedTasks.AddTask("CollectAndPushServerInfo", "*/3 * * * * *", corn_service.Task{

		LockTTL:  0,
		TaskCode: "CollectAndPushServerInfo",
		TaskDesc: "收集服务器信息并推送到Redis",
		TaskFunc: func() {
			if err := monitor_service.MonitorServerService.CollectAndPushServerInfo(); err != nil {
				core.Logger.Error("收集服务器信息并推送到Redis失败", err)
			}
		},
	})

	// 每天凌晨1点删除三个月前的错误监控数据
	FixedTasks.AddTask("DelMonitorErrorListThreeMonthAgo", "0 0 1 * * *", corn_service.Task{

		LockTTL:  10,
		TaskCode: "DelMonitorErrorListThreeMonthAgo",
		TaskDesc: "删除三个月前的错误监控数据",
		TaskFunc: func() {
			if err := monitor_service.MonitorErrorListService.DelThreeMonthAgo(); err != nil {
				core.Logger.Error("删除三个月前的错误监控数据失败", err)
			}
		},
	})

	// 每小时执行一次清理过期的分片临时目录
	FixedTasks.AddTask("CleanChunkTmpDir", "0 0 * * * *", corn_service.Task{
		LockTTL:  30,
		TaskCode: "CleanChunkTmpDir",
		TaskDesc: "清理过期的分片临时目录",
		TaskFunc: func() {
			storage.CleanChunkTmpDir()
		},
	})

}
