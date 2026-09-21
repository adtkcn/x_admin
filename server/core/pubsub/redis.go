package pubsub

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// redisEmitter Redis 后端：基于 Redis Pub/Sub 跨实例广播事件。
//
// 频道映射：事件类型 eventType 映射到 Redis 频道 {prefix}{eventType}。
// 通过 PSubscribe("{prefix}*") 一次性订阅本前缀下的所有事件类型，
// 收到消息后从频道名还原 eventType 再本地分发。
//
// Emit 语义：Publish 到 Redis，由各实例（含自身）的订阅循环收到后分发，
// 从而保证跨实例广播且行为一致。
type redisEmitter struct {
	client *redis.Client
	prefix string
	reg    *registry

	psub   *redis.PubSub
	ctx    context.Context
	cancel context.CancelFunc

	closeOnce sync.Once
	mu        sync.RWMutex
	closed    bool
}

// newRedisEmitter 创建 Redis 后端 Emitter，并启动后台订阅循环。
// prefix 为频道前缀（建议带业务/项目前缀以隔离），用于匹配 PSubscribe 的通配。
func newRedisEmitter(client *redis.Client, prefix string) *redisEmitter {
	ctx, cancel := context.WithCancel(context.Background())
	r := &redisEmitter{
		client: client,
		prefix: prefix,
		reg:    newRegistry(),
		ctx:    ctx,
		cancel: cancel,
	}
	// 订阅本前缀下的所有事件频道
	r.psub = client.PSubscribe(ctx, prefix+"*")
	go r.listen()
	return r
}

// listen 后台消费 Redis 订阅消息，还原事件类型并本地分发。
func (r *redisEmitter) listen() {
	for {
		msg, err := r.psub.ReceiveMessage(r.ctx)
		if err != nil {
			select {
			case <-r.ctx.Done():
				return
			default:
				// 断线重试，等待底层自动重连
				time.Sleep(100 * time.Millisecond)
				continue
			}
		}
		eventType := strings.TrimPrefix(msg.Channel, r.prefix)
		r.reg.dispatch(eventType, []byte(msg.Payload))
	}
}

func (r *redisEmitter) On(eventType string, h Handler) func() {
	return r.reg.on(eventType, h)
}

func (r *redisEmitter) Off(eventType string) {
	r.reg.off(eventType)
}

// Emit 发布事件到 Redis 频道 {prefix}{eventType}。
func (r *redisEmitter) Emit(eventType string, payload []byte) error {
	r.mu.RLock()
	closed := r.closed
	r.mu.RUnlock()
	if closed {
		return nil
	}
	return r.client.Publish(r.ctx, r.prefix+eventType, payload).Err()
}

func (r *redisEmitter) Close() error {
	r.closeOnce.Do(func() {
		r.mu.Lock()
		r.closed = true
		r.mu.Unlock()
		r.cancel()
		if r.psub != nil {
			r.psub.Close()
		}
		r.reg.reset()
	})
	return nil
}
