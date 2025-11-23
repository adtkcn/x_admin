package corn

import (
	"x_admin/core"
	"x_admin/util"
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
func init() {
	tm := core.NewCronManager()
	tm.Start()
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
