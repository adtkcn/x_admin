package corn

import (
	"x_admin/core"
	"x_admin/util"

	"gorm.io/gorm"
)

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

// 定义任务结构体
type TaskInfo struct {
	TaskCode string // 任务编码
	TaskDesc string // 任务描述
	TaskFunc func() // 任务函数
}

var TaskInfoList = []TaskInfo{
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

	Disabled bool      // 是否禁用
	TaskInfo *TaskInfo // 任务信息
}

// 从数据库加载任务
func LoadTasks(db *gorm.DB) []RunTask {
	// 从数据库加载任务
	var tasks []RunTask
	err := db.Where("disabled = ?", false).Find(&tasks).Error
	if err != nil {
		core.Logger.Error("从数据库加载任务失败", err)
		return []RunTask{}
	}
	var RunTaskList = []RunTask{} // 运行中的任务列表

	for _, task := range tasks {
		if task.Disabled == true {
			continue
		}
		for _, info := range TaskInfoList {
			if task.TaskCode == info.TaskCode {
				task.TaskInfo = &info
				RunTaskList = append(RunTaskList, task)
				break
			}
		}
	}
	return RunTaskList
}
func init() {
	// var db = core.GetDB()
	// var RunTaskList = LoadTasks(db)
	tm := core.NewCronManager()
	tm.Start()

	// for _, task := range RunTaskList {
	// 	if task.TaskInfo != nil {
	// 		err := tm.AddTask(task.TaskId, task.CronExpr, task.TaskInfo.TaskFunc)
	// 		if err != nil {
	// 			core.Logger.Error("添加任务失败", err)
	// 		}
	// 	}
	// }
	// defer tm.Stop()
	// 添加一个每5秒执行的任务
	tm.AddTask("broadcast", "*/5 * * * * *", func() {
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
}
