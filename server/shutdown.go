package main

import (
	"context"
	"net/http"
	"time"

	"x_admin/app/service/iot_service/pipeline"
	"x_admin/core"
	"x_admin/util/file_util"
)

// gracefulClose 按序优雅关闭全部资源：
// 先停止接收新 HTTP 请求并等待在途请求完成（根 ctx 已在外部取消，
// 队列 worker / ASR588 监听器 / cron 调度此时均已停止派发），最后释放
// WS / 队列 / Redis / DB 连接。
func gracefulClose(server *http.Server) {
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

	// 3.5) 队列已停生产、数据库未关：把时序批写缓冲里剩余数据同步落库，避免退出丢最后一批
	pipeline.StopSink()

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

	// 6) 关闭安全文件根目录。
	if err := file_util.Close(); err != nil {
		core.Logger.Errorf("安全文件根目录关闭失败: %v", err)
	}
}
