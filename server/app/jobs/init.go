package jobs

import (
	"fmt"
	"x_admin/core"

	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("定时任务：每5秒执行一次")
		core.Ws.SendToRoom("room1", []byte("hello room1"))
		core.Ws.SendToAll([]byte("hello all"))
	})
	// 启动定时任务
	c.Start()
}
