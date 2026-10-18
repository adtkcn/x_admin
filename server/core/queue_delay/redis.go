package queue_delay

import (
	"context"
	"encoding/json/v2"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisDelayedBackend 基于 Redis ZSET 的延迟队列（ZADD/ZRANGEBYSCORE/ZREM）。
type redisDelayedBackend struct {
	client *redis.Client
	prefix string

	// stopCh：关闭后停止取新消息（已在途的仍处理完）。由 Close() 关闭。
	stopCh chan struct{}
	// wg：统计已取出未处理完的消息数，Close() 等待其归零。
	wg sync.WaitGroup
	// closeOnce：保证 stopCh 只关闭一次。
	closeOnce sync.Once
}

func newRedisDelayedBackend(client *redis.Client, prefix string) *redisDelayedBackend {
	return &redisDelayedBackend{
		client: client,
		prefix: prefix,
		stopCh: make(chan struct{}),
	}
}

// delayedKey 延迟池键(ZSET)，score=到期毫秒时间戳。
func (b *redisDelayedBackend) delayedKey(name string) string {
	return b.prefix + name
}

// EnqueueAt 写入延迟池，到达 at 后才可被 ConsumeDelayed 取出。入队动作同样计入 wg。
func (b *redisDelayedBackend) EnqueueAt(name string, payload any, at time.Time) error {
	byte, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	b.wg.Add(1)
	defer b.wg.Done()
	return b.client.ZAdd(context.Background(), b.delayedKey(name),
		redis.Z{Score: float64(at.UnixMilli()), Member: string(byte)}).Err()
}

// EnqueueDelay 在 delay 之后才可消费。
func (b *redisDelayedBackend) EnqueueDelay(name string, payload any, delay time.Duration) error {
	return b.EnqueueAt(name, payload, time.Now().Add(delay))
}

// takeDueScript：原子取出并删除到期 member（KEYS[1]=ZSET, ARGV[1]=nowMillis, ARGV[2]=batch）。
// ZRANGEBYSCORE 默认按 score 升序，故取到期天然先到期先出（FIFO）；同一毫秒的多条按 member 字典序输出。
var takeDueScript = redis.NewScript(`
local tasks = redis.call('ZRANGEBYSCORE', KEYS[1], 0, ARGV[1], 'LIMIT', 0, ARGV[2])
if #tasks == 0 then return 0 end
redis.call('ZREM', KEYS[1], unpack(tasks))
return tasks
`)

// ConsumeDelayed 直接消费延迟池（ZSET）中已到期消息，不经即时队列 Consume。
// 每轮最多取 concurrency 条（与 worker 对齐，避免退出窗口囤积过多已删未处理消息）；
// 取空后 sleep，空轮询指数退避(500ms→5s)。ctx 取消或 Close() 只停取新，已在途必处理完。
func (b *redisDelayedBackend) ConsumeDelayed(ctx context.Context, name string, handler Handler, concurrency int) error {
	if concurrency <= 0 {
		concurrency = 1
	}
	zKey := b.delayedKey(name)

	// 有界 channel 提供背压：容量=并发数，满则阻塞轮询协程。
	ch := make(chan []byte, concurrency)

	// worker 池：并发处理；recover 按消息隔离——panic 既不能丢 wg 计数
	// （少 Done 会让 Close 的 wg.Wait 永久卡死），也不能带走整个 worker（并发能力递减）。
	for i := 0; i < concurrency; i++ {
		go func() {
			for body := range ch {
				b.handleOne(ctx, handler, name, body)
			}
		}()
	}

	go func() {
		// poller 退出时关闭 ch：worker 排空在途消息后自然退出，不永久滞留；
		// 只有本协程向 ch 发送，关闭与发送不会并发
		defer close(ch)
		const interval = 100 * time.Millisecond
		const maxInterval = 3 * time.Second
		batch := int64(concurrency) // 每轮取数上限
		sleep := interval
		for {
			// 停止取新：ctx 取消或 Close() 触发（仅停取，不丢在途）。
			select {
			case <-ctx.Done():
				return
			case <-b.stopCh:
				return
			default:
			}
			got := false
			// 循环取直到到期消息清空或达本轮上限。
			for {
				select {
				case <-ctx.Done():
					return
				case <-b.stopCh:
					return
				default:
				}
				res, err := takeDueScript.Run(ctx, b.client,
					[]string{zKey}, time.Now().UnixMilli(), batch).Result()
				if err != nil {
					log.Printf("[delayed:queue:redis] take due delayed messages of %q failed: %v", name, err)
					time.Sleep(100 * time.Millisecond)
					break
				}
				due, ok := res.([]interface{})
				if !ok || len(due) == 0 {
					break
				}
				got = true
				for _, m := range due {
					s, ok := m.(string)
					if !ok {
						continue
					}
					b.wg.Add(1)
					ch <- []byte(s) // 取出即计入在途，保证被处理
				}
				if int64(len(due)) < batch {
					break
				}
			}
			// 有消息重置退避，否则指数退避至上限。
			if got {
				sleep = interval
			} else {
				sleep *= 2
				if sleep > maxInterval {
					sleep = maxInterval
				}
			}
			select {
			case <-ctx.Done():
				return
			case <-b.stopCh:
				return
			case <-time.After(sleep):
			}
		}
	}()

	return nil
}

// handleOne 处理一条消息并结束在途计数：handler panic 也会执行 wg.Done，
// worker 本身驻留继续消费后续消息
func (b *redisDelayedBackend) handleOne(ctx context.Context, handler Handler, name string, body []byte) {
	defer b.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[delayed:queue:redis] handler of %q panicked: %v", name, r)
		}
	}()
	if err := handler(ctx, body); err != nil {
		log.Printf("[delayed:queue:redis] handle delayed message from %q failed: %v", name, err)
	}
}

// Close 优雅关闭：停取新消息，并阻塞等待已在途消息全部处理完（wg.Wait）。客户端由 core 管理，不在此关闭。
func (b *redisDelayedBackend) Close() error {
	b.closeOnce.Do(func() {
		close(b.stopCh)
	})
	b.wg.Wait()
	return nil
}
