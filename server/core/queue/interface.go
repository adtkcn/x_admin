// Package queue 提供一个与业务解耦的异步任务队列。
//
// 设计目标：
//   - 统一接口：调用方只依赖 Queue 接口，不感知底层是哪种后端。
//   - 多后端：内置 redis 实现，可平滑扩展（新增后端只需实现 Queue 接口）。
//   - 无业务语义：消息体统一为 []byte，业务层自行序列化；队列不关心任务类型。
//   - 易扩展：新增后端只需实现 Queue 接口，并在 config.New 中注册。
//
// 语义为「点对点」工作队列：同一条消息只会被一个消费者取出（多个 worker 竞争消费），
// 适用于异步任务、后台作业等场景。Pub/Sub 广播语义请使用 core/pubsub。
package queue

import "context"

// Handler 消费一条消息的回调。body 为原始字节载荷，由业务层自行反序列化。
// 返回 error 表示处理失败；当前实现仅记录日志，不自动重试（重试/死信队列见 README）。
type Handler func(ctx context.Context, body []byte) error

// Queue 统一队列接口。底层后端可为 local 或 redis，调用方只依赖本接口。
type Queue interface {
	// Enqueue 向指定队列 name 投递一条消息）。
	Enqueue(name string, body any) error
	// Consume 启动消费者，从队列 name 持续拉取消息并交给 handler 处理。
	// concurrency 为 worker 并发数（<=0 时回落为 1）。ctx 取消后所有 worker 优雅退出。
	Consume(ctx context.Context, name string, handler Handler, concurrency int) error
	// Close 释放底层资源（redis 不主动关闭共享客户端）。
	Close() error
}
