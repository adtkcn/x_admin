package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"x_admin/app/model/user_model"
	"x_admin/app/task"

	"x_admin/config"
	"x_admin/core"
	"x_admin/plugin"
	"x_admin/routes"

	app_corn "x_admin/app/corn"

	"github.com/gin-gonic/gin"
)

// initServer 初始化server
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":" + strconv.Itoa(config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,  //从连接建立到读取完整请求头和 body的最大时间
		WriteTimeout:   100 * time.Second, // 从读取完请求到写完响应的最大时间
		MaxHeaderBytes: 8192,              // 8KB,请求头最大字节数
	}
}

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

	plugin.RegisterNullValidator()
	// 刷新日志缓冲
	defer core.Logger.Sync()

	// 自动迁移用户表
	core.AutoMigrate(
		&user_model.User{},
		&user_model.UserAuth{},
	)

	// 初始化微信 SDK 客户端（小程序+公众号）
	plugin.InitWechatClients()

	// 基础全局队列实例由 core 包创建（见 core.Queue）
	// 此处仅作为启动钩子拉起消费者，ctx 取消后所有 worker 优雅退出。
	task.Start(ctx)

	// 初始化router
	router := routes.InitRouter()
	// 初始化server
	s := initServer(router)

	fmt.Println("格式化文档注释:", "swag fmt")
	fmt.Println("生成文档:", "swag init")
	fmt.Printf("文档: http://localhost:%v/api/static/api/index.html\n", config.AppConfig.Port)

	// 在独立 goroutine 中运行服务，主协程负责等待关闭信号并优雅退出。
	serverErr := make(chan error, 1)
	go func() {
		core.Logger.Infof("HTTP 服务启动，监听 :%d", config.AppConfig.Port)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	// 等待关闭信号或运行期致命错误。
	select {
	case err := <-serverErr:
		core.Logger.Errorf("HTTP 服务异常退出: %v", err)
	case <-ctx.Done():
		core.Logger.Infof("收到关闭信号，开始优雅关闭...")
		// 立即停掉所有后台派发，避免关闭窗口期触发新任务：
		// 取消根 ctx（队列 worker / ASR588 监听器退出）+ 停止 cron 调度。
		stop()
		app_corn.FixedTasks.Stop()
		app_corn.DynamicTasks.Stop()
	}

	// 按序优雅关闭所有资源。
	shutdown(s)

	core.Logger.Infof("服务已优雅关闭")
}

// shutdown 按序优雅关闭全部资源：
// 先停止接收新 HTTP 请求并等待在途请求完成（根 ctx 已在外部取消，
// 队列 worker / ASR588 监听器 / cron 调度此时均已停止派发），最后释放
// WS / 队列 / Redis / DB 连接。
func shutdown(server *http.Server) {
	// 1) 停止接收新请求，等待在途请求完成（最多等待 30s）。
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		core.Logger.Errorf("HTTP 服务优雅关闭失败（强制退出）: %v", err)
	} else {
		core.Logger.Infof("HTTP 服务已停止接收新请求，在途请求处理完毕")
	}

	// 2) 关闭 WebSocket 管理器（停心跳、删在线 SET、关 PubSub）。
	if core.Ws != nil {
		core.Ws.Close()
	}

	// 3) 关闭队列（释放底层资源）。
	if core.Queue != nil {
		if err := core.Queue.Close(); err != nil {
			core.Logger.Errorf("队列关闭失败: %v", err)
		}
	}
	if core.QueueDelay != nil {
		if err := core.QueueDelay.Close(); err != nil {
			core.Logger.Errorf("延迟队列关闭失败: %v", err)
		}
	}

	// 4) 关闭 Redis 连接。
	if core.Redis != nil {
		if err := core.Redis.Close(); err != nil {
			core.Logger.Errorf("Redis 关闭失败: %v", err)
		}
	}

	// 5) 关闭数据库连接。
	if core.GetDB() != nil {
		if db, err := core.GetDB().DB(); err == nil {
			if closeErr := db.Close(); closeErr != nil {
				core.Logger.Errorf("数据库连接关闭失败: %v", closeErr)
			}
		}
	}
}
