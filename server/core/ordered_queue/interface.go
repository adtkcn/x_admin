// Package ordered_queue 提供「按 key 分片、每片单消费者」的有序队列。
//
// 与 core/queue 的区别只有一个，但这个区别决定了它为什么要独立成包：
//   - queue：竞争消费，吞吐优先。同一条消息只被一个 worker 取到，但**不保证处理顺序**。
//   - ordered_queue：分片消费，顺序优先。同一分片键的消息严格按投递顺序被串行处理。
//
// 典型用途是设备报文处理：同一设备先后两条报文若被并发处理，后到的可能先落库，
// 在线状态与告警状态机就会出现「旧值覆盖新值」。按 deviceId 分片后，同设备必同片、
// 同片全局必只有一个消费者，顺序由结构本身保证，业务侧无需再加锁。
//
// 语义边界（与 queue 一致，刻意不做）：
//   - at-most-once：BLPOP 取出即出队，处理中进程崩溃会丢当前这一条。
//   - 不重试：有序队列就地重试会阻塞同片全部后续消息，需要重试应配死信队列。
//   - 只保证同分片键有序：不保证跨键顺序，也不保证入队侧顺序（多生产者并发 RPUSH 本身有竞态）。
//
// Redis 键布局（name 为队列名，idx 为分片号）：
//   - {Prefix}{name}:{idx}       List，分片消息体（RPUSH 右进 / BLPOP 左出，合成 FIFO）
//   - {Prefix}{name}:lease:{idx} String，分片租约（值为持有者 ID，带 TTL）
package ordered_queue

import (
	"context"
)

// Handler 消费一条消息的回调。body 为原始字节载荷，由业务层自行反序列化。
// 返回 error 仅记录日志，不重试（有序队列重试会阻塞同片后续消息）。
type Handler func(ctx context.Context, body []byte) error

// OrderedQueue 有序队列接口。底层后端可为 redis，调用方只依赖本接口。
type OrderedQueue interface {
	// Enqueue 按 shardKey（如 deviceId）哈希投递到对应分片。
	// 同一 shardKey 的先后投递顺序，即该 key 的处理顺序。
	Enqueue(name string, shardKey string, payload any) error
	// Consume 为队列 name 启动分片消费：内部为每个分片起一个认领协程，
	// 抢到该分片租约就串行消费，抢不到说明已由其他实例持有；本实例下线/失锁后
	// 其他实例自动接管。调用方无需指定并发数与分片归属。
	// ctx 取消后停止取新消息，已在途的消息处理完。
	Consume(ctx context.Context, name string, handler Handler) error
	// Close 释放底层资源（redis 不主动关闭共享客户端）。
	Close() error
}
