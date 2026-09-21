// Package queue_delay 提供与业务解耦的延迟（定时）投递队列，与 core/queue 即时队列互不耦合。
//
// 设计目标：
//   - 独立包：延迟投递能力单独成包，调用方可按需引入，不影响即时队列 core/queue。
//   - 多后端：内置 redis 实现（基于 ZSET 的到期轮询），可平滑扩展。
//   - 无业务语义：消息体统一为 []byte，业务层自行序列化。
//
// 延迟消息投递后不会立即被消费，而是到达指定时间后才可被 ConsumeDelayed 取出处理。
// 采用「直接消费延迟池」模式：不经过就绪 List、不依赖即时队列的 Consume。
package queue_delay

import (
	"context"
	"time"
)

// Handler 消费一条消息的回调。签名与 core/queue 的 Handler 一致，但本包独立定义、不依赖 core/queue。
type Handler func(ctx context.Context, body []byte) error

// DelayedQueue 延迟投递能力接口。消息到达指定时间后才可被取出消费。
type DelayedQueue interface {
	// EnqueueAt 在指定的绝对时间点 at 之后才可被 ConsumeDelayed 消费。
	EnqueueAt(name string, payload any, at time.Time) error
	// EnqueueDelay 在 delay 时长之后才可被消费（等价于 EnqueueAt(name, payload, time.Now().Add(delay))）。
	EnqueueDelay(name string, payload any, delay time.Duration) error
	// ConsumeDelayed 直接消费延迟池(name 对应 ZSET)中已到期的消息，不经过即时队列的 Consume/就绪 List。
	// 周期性轮询到期消息，取出后原子删除并交给 handler，语义与 Queue.Consume 一致（同一条消息仅被一个 worker 处理）。
	// concurrency 为并发处理数（<=0 回落为 1）。ctx 取消后优雅退出。
	ConsumeDelayed(ctx context.Context, name string, handler Handler, concurrency int) error
	// Close 释放底层资源（redis 不主动关闭共享客户端）。
	Close() error
}
