package queue

import (
	"context"
	"encoding/json/v2"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// popTimeout 阻塞取空的超时，必须有限。用 0 永久阻塞时，只调 Close()（不取消 ctx）
// 会让取消息协程醒不过来，wg.Wait 挂到下一条消息才退；2s 把「感知停机」的代价钉死在 2s 内。
const popTimeout = 2 * time.Second

// redisBackend 基于 Redis List 的跨实例工作队列（RPUSH 右进 / BLPOP 左出，先进先出）。
// 注意不可改用 BRPOP：它与 RPUSH 同端进出，突发批量投递会变成后进先出。
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

// Enqueue 将消息追加到队列尾部（右进，与 BLPOP 的左出合成先进先出）。入队动作本身计入 wg（推前 Add，写成功后 Done），保证入队正常完成。
func (b *redisBackend) Enqueue(name string, payload any) error {
	byte, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	b.wg.Add(1)
	defer b.wg.Done()
	return b.client.RPush(context.Background(), b.key(name), byte).Err()
}

// Consume 单连接 BLPOP 取消息（左出，先投先取）→ 有界 channel(cap=concurrency) → worker 池并发处理。
// 仅占用 1 条 Redis 连接；ctx 取消或 Close() 只停止取新，已在途消息必处理完（wg.Wait）。
// 本包只保证「取用顺序」为投递顺序；worker 池并发，处理完成顺序不保证，需要严格有序请用 core/ordered_queue。
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

	// 取消息协程：BLPOP 阻塞拉取；ctx 取消或 stopCh 关闭即停取（仅停取，不丢在途）。
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-b.stopCh:
				return
			default:
			}
			res, err := b.client.BLPop(ctx, popTimeout, key).Result()
			if err != nil {
				if errors.Is(err, redis.Nil) {
					continue // 空轮超时：无消息，回到循环查停机位
				}
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
