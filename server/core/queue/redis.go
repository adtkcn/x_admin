package queue

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisBackend 基于 Redis List 的跨实例工作队列（RPUSH 投递 / BRPOP 阻塞消费）。
// 适合多实例部署：任意实例投递的任务可被任一实例的 worker 取到。
type redisBackend struct {
	client *redis.Client
	prefix string
}

func newRedisBackend(client *redis.Client, prefix string) *redisBackend {
	return &redisBackend{client: client, prefix: prefix}
}

func (b *redisBackend) key(name string) string {
	return b.prefix + name
}

// Enqueue 使用 RPUSH 将消息追加到队列尾部。
func (b *redisBackend) Enqueue(name string, payload any) error {
	byte, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return b.client.RPush(context.Background(), b.key(name), byte).Err()
}

// Consume 单连接 BRPOP 取消息，投入有界 channel（cap=concurrency），
// 再由 concurrency 个 worker 并发处理。
//
// 对比原生多连接方案：取与处理解耦，仅占用 1 条 Redis 连接（而非 concurrency 条），
// 在保持相同并发能力与背压语义（channel 满即阻塞取消息协程）的前提下大幅减少连接数；
// Redis 保证同一消息只被一个 worker 取到。ctx 取消后先停止取消息，再 drain 在途任务优雅退出。
func (b *redisBackend) Consume(ctx context.Context, name string, handler Handler, concurrency int) error {
	if concurrency <= 0 {
		concurrency = 1
	}
	key := b.key(name)

	// 有界 channel：容量=并发数，满了则阻塞取消息协程，天然提供背压，避免内存无限堆积。
	ch := make(chan []byte, concurrency)

	// worker 池：从 channel 取消息并发处理
	for i := 0; i < concurrency; i++ {
		go func() {
			for body := range ch {
				if err := handler(ctx, body); err != nil {
					log.Printf("[queue:redis] handle message from %q failed: %v", name, err)
				}
			}
		}()
	}

	// 单连接取消息协程：BRPOP 阻塞拉取，取到即投入有界 channel
	go func() {
		for {
			res, err := b.client.BRPop(ctx, 0, key).Result()
			if err != nil {
				if ctx.Err() != nil {
					// ctx 取消：关闭 channel，通知 worker 退出（drain 在途任务）
					close(ch)
					return
				}
				// 网络抖动：短暂退避后重试，避免 CPU 空转
				time.Sleep(100 * time.Millisecond)
				continue
			}
			if len(res) < 2 {
				continue
			}
			// channel 满时此处阻塞，自动背压；ctx 取消后 range 结束，worker 退出
			ch <- []byte(res[1])
		}
	}()

	return nil
}

// Close Redis 客户端由 core 统一管理，这里不主动关闭共享连接。
func (b *redisBackend) Close() error {
	return nil
}
