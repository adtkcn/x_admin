package ws

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

// RedisPubSub Redis Pub/Sub 消息总线实现（集群模式）。
//
// 使用 Redis Pub/Sub 实现跨实例消息广播，支持多节点部署。
// 所有实例订阅同一个 Redis 频道，任一实例发布消息时，所有实例都会收到。
//
// 架构设计：
//   - 双缓冲机制：Redis 订阅消息先由 listen() 反序列化后写入内部 channel，
//     再由 Manager.subscribeLoop 消费，避免阻塞 Redis 客户端连接
//   - 非阻塞写入：Publish 和 listen 中的 channel 写入都是非阻塞的，
//     channel 满时丢弃消息，保证高并发下的系统稳定性
//   - 自动重连：依赖 go-redis 内置的重连机制，Redis 断线后自动重连并重新订阅
//
// 消息流转：
//
//	Publish() → Redis PUBLISH → 所有实例的 listen() 收到
//	→ 反序列化写入内部 channel → Manager.subscribeLoop 消费 → 本地推送
//
// 注意事项：
//   - 消息会广播到所有实例（包括发送者自身），各实例自行过滤本地连接
//   - Redis 频道名应使用项目前缀（如 "x:ws:broadcast"），避免多项目冲突
//   - 消息体大小受 Redis 限制（理论 512MB），实际应控制在合理范围
type RedisPubSub struct {
	client    *redis.Client      // Redis 客户端实例
	pubsub    *redis.PubSub      // Redis Pub/Sub 订阅对象
	ch        chan PubSubMessage // 内部消息通道，缓冲大小 256
	ctx       context.Context
	cancel    context.CancelFunc
	channel   string       // Redis 频道名，如 "x:ws:broadcast"
	closeOnce sync.Once    // 保证 Close 只执行一次
	closed    bool         // 标记是否已关闭，防止 Close 后 listen/Publish 写入 channel
	mu        sync.RWMutex // 保护 closed 标志
}

// NewRedisPubSub 创建 Redis 消息总线。
//
// 参数：
//   - client: Redis 客户端实例，由 core.Redis 提供
//   - channel: Redis 频道名，建议使用项目前缀（如 config.RedisConfig.RedisPrefix + "ws:broadcast"）
//
// 初始化流程：
//  1. 创建带缓冲的内部 channel（容量 256）
//  2. 调用 Redis SUBSCRIBE 订阅指定频道
//  3. 启动 listen() goroutine 持续监听 Redis 消息并转发到内部 channel
//
// 返回后可通过 Subscribe() 获取内部 channel 的只读引用，供 Manager.subscribeLoop 消费。
func NewRedisPubSub(client *redis.Client, channel string) *RedisPubSub {
	ctx, cancel := context.WithCancel(context.Background())
	r := &RedisPubSub{
		client:  client,
		ch:      make(chan PubSubMessage, 256),
		ctx:     ctx,
		cancel:  cancel,
		channel: channel,
	}
	// 订阅 Redis 频道，go-redis 会自动处理连接池和重连
	r.pubsub = client.Subscribe(ctx, channel)
	// 启动监听协程，将 Redis 消息转发到内部 channel
	go r.listen()
	return r
}

// listen 监听 Redis 订阅消息并转发到内部 channel。
//
// 该函数在独立 goroutine 中运行，持续从 Redis Pub/Sub 接收消息：
//  1. 从 r.pubsub.Channel() 获取 Redis 消息 channel
//  2. 反序列化 JSON 消息为 PubSubMessage 结构
//  3. 非阻塞写入内部 channel，channel 满时丢弃消息并记录日志
//
// 退出条件：
//   - ctx 被取消（调用 Close() 时）
//   - Redis 消息 channel 被关闭（Redis 连接异常时）
//
// 设计要点：
//   - 使用非阻塞写入避免阻塞 Redis 客户端连接，防止影响其他 Redis 操作
//   - 反序列化失败时记录日志并继续处理下一条消息，不会中断监听
func (r *RedisPubSub) listen() {
	ch := r.pubsub.Channel()
	for {
		select {
		case <-r.ctx.Done():
			// context 被取消，正常退出
			return
		case msg, ok := <-ch:
			if !ok {
				// Redis 消息 channel 被关闭，退出
				return
			}
			// 反序列化 JSON 消息
			var psMsg PubSubMessage
			if err := json.Unmarshal([]byte(msg.Payload), &psMsg); err != nil {
				log.Printf("[ws] redis pubsub unmarshal error: %v", err)
				continue
			}
			// 非阻塞写入内部 channel，防止慢消费者阻塞 Redis 监听
			r.mu.RLock()
			if r.closed {
				r.mu.RUnlock()
				return
			}
			select {
			case r.ch <- psMsg:
			default:
				log.Printf("[ws] redis pubsub channel full, message dropped")
			}
			r.mu.RUnlock()
		}
	}
}

// Publish 发布消息到 Redis 频道。
//
// 将 PubSubMessage 序列化为 JSON 后，通过 Redis PUBLISH 命令广播到所有订阅者。
// 所有实例（包括发送者自身）都会收到该消息，各实例在 subscribeLoop 中检查本地连接并按需推送。
//
// 参数：
//   - ctx: 上下文，用于控制发布操作的超时和取消
//   - msg: 要发布的消息结构
//
// 返回值：
//   - error: Redis 发布失败时返回错误，成功时返回 nil
//
// 注意事项：
//   - 消息会被所有订阅该频道的实例收到，包括发送者自身
//   - Redis Pub/Sub 是 fire-and-forget 模式，不持久化消息，断线期间的消息会丢失
func (r *RedisPubSub) Publish(ctx context.Context, msg PubSubMessage) error {
	r.mu.RLock()
	if r.closed {
		r.mu.RUnlock()
		return nil
	}
	r.mu.RUnlock()

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return r.client.Publish(ctx, r.channel, data).Err()
}

// Subscribe 订阅消息，返回内部 channel 的只读引用。
//
// Manager.Init 时调用，返回的 channel 会被 subscribeLoop 持续监听。
// 该 channel 由 listen() goroutine 写入，包含从 Redis 反序列化后的消息。
//
// 注意：RedisPubSub 是单订阅者模型，多次调用 Subscribe 返回同一个 channel。
func (r *RedisPubSub) Subscribe(ctx context.Context) (<-chan PubSubMessage, error) {
	return r.ch, nil
}

// Close 关闭消息总线，释放资源。
//
// 通过 sync.Once 保证只执行一次：
//  1. 调用 cancel 通知 listen() goroutine 退出
//  2. 关闭 Redis Pub/Sub 订阅，释放 Redis 连接资源
//  3. 关闭内部 channel，使 subscribeLoop 的 range 循环正常退出
func (r *RedisPubSub) Close() error {
	r.closeOnce.Do(func() {
		// 先设置 closed 标志，防止 listen() 和 Publish() 写入已关闭的 channel
		r.mu.Lock()
		r.closed = true
		r.mu.Unlock()

		r.cancel()
		if r.pubsub != nil {
			r.pubsub.Close()
		}
		// 关闭内部 channel，使 subscribeLoop 的 range 循环正常退出
		close(r.ch)
	})
	return nil
}
