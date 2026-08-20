package queue

import (
	"context"
	"encoding/json/v2"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisBackend 基于 Redis List 的跨实例工作队列（RPUSH/BRPOP）。
type redisBackend struct {
	client *redis.Client
	prefix string

	// stopCh：关闭后停止取新消息（已在途的仍处理完）。由 Close() 关闭。
	stopCh chan struct{}
	// wg：统计已取出未处理完的消息数，Close() 等待其归零。
	wg sync.WaitGroup
	// closeOnce：保证 stopCh 只关闭一次。
	closeOnce sync.Once
}

func newRedisBackend(client *redis.Client, prefix string) *redisBackend {
	return &redisBackend{
		client: client,
		prefix: prefix,
		stopCh: make(chan struct{}),
	}
}

func (b *redisBackend) key(name string) string {
	return b.prefix + name
}

// Enqueue 将消息追加到队列尾部。入队动作本身计入 wg（推前 Add，写成功后 Done），保证入队正常完成。
func (b *redisBackend) Enqueue(name string, payload any) error {
	byte, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	b.wg.Add(1)
	defer b.wg.Done()
	return b.client.RPush(context.Background(), b.key(name), byte).Err()
}

// Consume 单连接 BRPOP 取消息 → 有界 channel(cap=concurrency) → worker 池并发处理。
// 仅占用 1 条 Redis 连接；ctx 取消或 Close() 只停止取新，已在途消息必处理完（wg.Wait）。
func (b *redisBackend) Consume(ctx context.Context, name string, handler Handler, concurrency int) error {
	if concurrency <= 0 {
		concurrency = 1
	}
	key := b.key(name)

	// 有界 channel 提供背压：满则阻塞取消息协程。
	ch := make(chan []byte, concurrency)

	// worker 池：并发处理；handler panic 不影响其他 worker；处理完才 wg.Done（不受 ctx 打断）。
	for i := 0; i < concurrency; i++ {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("[queue:redis] worker of %q panicked: %v", name, r)
				}
			}()
			for body := range ch {
				if err := handler(ctx, body); err != nil {
					log.Printf("[queue:redis] handle message from %q failed: %v", name, err)
				}
				b.wg.Done()
			}
		}()
	}

	// 取消息协程：BRPOP 阻塞拉取；ctx 取消或 stopCh 关闭即停取（仅停取，不丢在途）。
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-b.stopCh:
				return
			default:
			}
			res, err := b.client.BRPop(ctx, 0, key).Result()
			if err != nil {
				select {
				case <-ctx.Done():
					return
				case <-b.stopCh:
					return
				default:
					time.Sleep(100 * time.Millisecond) // 网络抖动退避
					continue
				}
			}
			if len(res) < 2 {
				continue
			}
			b.wg.Add(1)
			ch <- []byte(res[1]) // 取出即计入在途，保证被处理
		}
	}()

	return nil
}

// Close 优雅关闭：停取新消息，并阻塞等待已在途消息全部处理完（wg.Wait）。客户端由 core 管理，不在此关闭。
func (b *redisBackend) Close() error {
	b.closeOnce.Do(func() {
		close(b.stopCh)
	})
	b.wg.Wait()
	return nil
}
