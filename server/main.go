package main

import (
	"context"
	"os/signal"
	"syscall"

	"x_admin/core"
	"x_admin/routes"
)

//	@description	x_admin后台管理系统
//	@termsOfService	http://x.adtk.cn

//	@contact.name	xh
//	@contact.url	http://x.adtk.cn
//	@contact.email	x@adtk.cn

// @license.name	MIT License
// @license.url	https://gitee.com/xiangheng/x_admin/blob/main/LICENSE
// @BasePath		/
func main() {
	// 根 context：捕获 SIGINT/SIGTERM，用于驱动全链路优雅关闭。
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 刷新日志缓冲（进程退出前 flush）
	defer core.Logger.Sync()

	// 1) 引导：初始化各子系统并拉起后台消费者（队列 / cron / IoT 流水线）
	if err := bootstrap(ctx); err != nil {
		return
	}

	// 2) 起服务：建 HTTP server、启动监听，阻塞直到收到关闭信号或运行期致命错误
	server := newServer(routes.InitRouter())
	printDocHints()
	serveUntilStop(server, ctx, stop)

	// 3) 关闭：按序释放全部资源
	gracefulClose(server)
	core.Logger.Infof("服务已优雅关闭")
}
