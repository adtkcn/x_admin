package ws

import (
	"context"
	"sync"
)

// LocalPubSub 本地消息总线实现（单机模式）。
//
// 使用 Go channel 在进程内传递消息，无网络开销，无序列化成本。
// 适用于单实例部署场景，是默认的 PubSub 实现。
//
// 特点：
//   - 零依赖：不依赖外部中间件
//   - 高性能：消息直接写入 channel，无 JSON 序列化/反序列化
//   - 非阻塞：Publish 在 channel 满时丢弃消息，不会阻塞调用方
//   - 单订阅者：只有一个 subscribeLoop 消费消息（与 Manager 1:1 绑定）
type LocalPubSub struct {
	ch        chan PubSubMessage // 消息通道，缓冲大小 256
	ctx       context.Context
	cancel    context.CancelFunc
	closeOnce sync.Once    // 保证 Close 只执行一次
	closed    bool         // 标记 channel 是否已关闭，防止 Publish 写入已关闭的 channel
	mu        sync.RWMutex // 保护 closed 标志
}

// NewLocalPubSub 创建本地消息总线。
//
// 初始化一个带缓冲的 channel（容量 256），作为消息传递的媒介。
// 返回后可通过 Subscribe 获取该 channel 的只读引用。
func NewLocalPubSub() *LocalPubSub {
	ctx, cancel := context.WithCancel(context.Background())
	return &LocalPubSub{
		ch:     make(chan PubSubMessage, 256),
		ctx:    ctx,
		cancel: cancel,
	}
}

// Publish 发布消息（非阻塞写入）。
//
// 将消息写入内部 channel。如果 channel 已满（256 条未消费），
// 则丢弃当前消息并立即返回，不会阻塞调用方。
// 这种设计保证了高并发下发送方不会被慢消费者拖慢。
func (l *LocalPubSub) Publish(ctx context.Context, msg PubSubMessage) error {
	l.mu.RLock()
	if l.closed {
		l.mu.RUnlock()
		return nil
	}
	defer l.mu.RUnlock()
	select {
	case l.ch <- msg:
	default:
		// channel 满，丢弃消息，防止阻塞
	}
	return nil
}

// Subscribe 订阅消息，返回内部 channel 的只读引用。
//
// Manager.Init 时调用，返回的 channel 会被 subscribeLoop 持续监听。
// 注意：LocalPubSub 是单订阅者模型，多次调用 Subscribe 返回同一个 channel。
func (l *LocalPubSub) Subscribe(ctx context.Context) (<-chan PubSubMessage, error) {
	return l.ch, nil
}

// Close 关闭消息总线，释放资源。
//
// 通过 sync.Once 保证只执行一次：
//  1. 调用 cancel 通知所有依赖 context 的 goroutine
//  2. 关闭 channel，使 subscribeLoop 的 range 循环正常退出
func (l *LocalPubSub) Close() error {
	l.closeOnce.Do(func() {
		l.mu.Lock()
		l.closed = true
		l.mu.Unlock()
		l.cancel()
		close(l.ch)
	})
	return nil
}
