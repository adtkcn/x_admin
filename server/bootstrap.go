package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	app_corn "x_admin/app/corn"
	fabu_model "x_admin/app/model/fabu_model"
	iot_model "x_admin/app/model/iot_model"
	"x_admin/app/model/user_model"
	"x_admin/app/queue"
	"x_admin/app/service/iot_service/pipeline"
	"x_admin/config"
	"x_admin/core"

	"x_admin/util/validation"

	"github.com/gin-gonic/gin"
)

// bootstrap 启动阶段的初始化：显示注册 x_null 校验器 → 自动迁移表结构 → 拉起队列消费者。
// 微信 SDK、安全文件根已改为各包 init 自注册（由侧效应导入/包加载触发）。
func bootstrap(ctx context.Context) error {
	// 显示配置 gin 绑定校验器：注册 x_null 自定义类型 + 中文错误翻译
	validation.Register()

	autoMigrate()

	// 基础全局队列实例由 core 包创建（见 core.Queue）
	// 此处仅作为启动钩子拉起所有队列消费者（含 IoT 报文流水线），ctx 取消后 worker 优雅退出。
	// 必须在 AutoMigrate 之后（IoT Sink 写 x_iot_device_data）。
	// 拉起消费前先重建告警运行态：服务停机期 Redis 活动态可能丢失，而库里仍「告警中」的
	// 记录只靠活动态才能收尾，不回灌则重启后恢复报文会永久卡在告警中。
	pipeline.RecoverAlertRuntime()

	queue.Start(ctx)

	return nil
}

// autoMigrate 自动迁移业务表结构。
func autoMigrate() {
	core.AutoMigrate(
		&user_model.User{},
		&user_model.UserAuth{},
		&fabu_model.FabuApp{},
		&fabu_model.FabuAppVersion{},
		&fabu_model.FabuWgt{},
		&iot_model.IotProductCategory{},
		&iot_model.IotProduct{},
		&iot_model.IotRuleChain{},
		&iot_model.IotDeviceSpace{},
		&iot_model.IotDevice{},
		// 时序数据：本轮先用 MySQL 暂代 TDengine（iot.md 第四阶段接入后此表可下线）
		&iot_model.IotDeviceData{},
		// 告警历史：产品 goja 告警脚本 + 处理层状态机（持续/恢复）收敛成事件级记录
		&iot_model.IotAlertHistory{},
	)
}

// newServer 初始化 HTTP server
func newServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           ":" + strconv.Itoa(config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,  //从连接建立到读取完整请求头和 body的最大时间
		WriteTimeout:   100 * time.Second, // 从读取完请求到写完响应的最大时间
		MaxHeaderBytes: 8192,              // 8KB,请求头最大字节数
	}
}

// printDocHints 打印接口文档相关提示
func printDocHints() {
	fmt.Println("格式化文档注释:", "swag fmt")
	fmt.Println("生成文档:", "swag init")
	fmt.Printf("文档: http://localhost:%v/api/static/api/index.html\n", config.AppConfig.Port)
}

// serveUntilStop 在独立 goroutine 中启动 HTTP 服务，主协程阻塞等待关闭信号或运行期致命错误。
// 收到关闭信号时先停掉后台派发（取消根 ctx + 停止 cron），避免关闭窗口期再触发新任务；
// 本函数返回后由 main 继续执行 gracefulClose 释放资源。
func serveUntilStop(server *http.Server, ctx context.Context, stop context.CancelFunc) {
	serverErr := make(chan error, 1)
	go func() {
		core.Logger.Infof("HTTP 服务启动，监听 :%d", config.AppConfig.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	// 等待关闭信号或运行期致命错误。
	select {
	case err := <-serverErr:
		core.Logger.Errorf("HTTP 服务异常退出: %v", err)
	case <-ctx.Done():
		core.Logger.Infof("收到关闭信号，开始优雅关闭...")
		// 取消根 ctx（队列 worker / ASR588 监听器退出）+ 停止 cron 调度。
		stop()
		app_corn.FixedTasks.Stop()
		app_corn.DynamicTasks.Stop()
	}
}
