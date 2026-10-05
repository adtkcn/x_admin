package corn

import (
	"strconv"
	"time"
	"x_admin/app/iot/pipeline"
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

	// 每n秒从 IoT 到期表（ZSET）弹出已到期设备置离线：报文/心跳由 touchAlive 刷新
	// “到期时间=到达时间+设备 offline_timeout（秒，设备列已冗余产品值）”，
	FixedTasks.AddTask("iotOfflineScan", "*/5 * * * * *", corn_service.Task{
		LockTTL:  25,
		TaskCode: "iotOfflineScan",
		TaskDesc: "弹出到期成员并置 IoT 设备离线",
		TaskFunc: func() {
			pipeline.ScanOffline()
		},
	})

	// 每5分钟跑一次 IoT 离线判定慢路径：全扫库侧在线设备按 seen+窗口现算兜底——
	// 超窗置离线、未超窗但 score 缺失/已过期的回填到期表
	// （防 Redis 丢数据后设备永远在线，并顺带修正改过窗口的旧 score）。
	FixedTasks.AddTask("iotOfflineSweep", "0 */5 * * * *", corn_service.Task{
		LockTTL:  240,
		TaskCode: "iotOfflineSweep",
		TaskDesc: "IoT 离线判定慢路径与到期表重建",
		TaskFunc: func() {
			pipeline.SweepOffline()
		},
	})

	// 每天凌晨3点按保留期(IOT.DataKeepDays)清理 IoT 时序历史，分批小事务删除；
	// 配置为 0 表示永久保留，任务空转
	FixedTasks.AddTask("iotDataClean", "0 0 3 * * *", corn_service.Task{
		LockTTL:  600,
		TaskCode: "iotDataClean",
		TaskDesc: "IoT 时序数据保留期清理",
		TaskFunc: func() {
			pipeline.CleanData()
		},
	})

}
