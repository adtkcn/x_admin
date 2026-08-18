// Package pubsub 提供 mitt 风格的事件发布/订阅（Emitter），支持可插拔的多种后端。
//
// 设计目标：
//   - 统一 API：无论单机还是集群，调用方只面对同一套 mitt 风格接口
//     （On / Off / Emit，按事件类型分发，支持通配符 "*"）。
//   - 多后端：底层传输可插拔，当前内置 redis（Redis Pub/Sub 跨实例广播）；
//     新增后端（NATS/Kafka 等）只需实现 Emitter 接口并在 New 中注册分支。
//
// 关于 payload 类型：
//
//	为让各后端语义一致（跨进程必须可序列化），Emit/Handler
//	统一使用 []byte 作为载荷。本包只做透明传输，不定义任何业务语义；
//	载荷的编码/解码（JSON、protobuf 等）完全由调用方负责。
package pubsub

// All 通配事件类型。订阅它可收到任意类型的事件。
const All = "*"

// Handler 事件处理函数。payload 为透明字节载荷，由订阅方自行解码。
type Handler func(payload []byte)

// Emitter mitt 风格事件发布/订阅器，API 与语义参考 mitt.js。
//
//   - On(type, handler)：订阅某类型事件，返回取消订阅函数
//   - Off(type)：移除该类型下的全部订阅
//   - Emit(type, payload)：发布事件（redis 经总线广播到所有实例，含本实例）
//   - 通配符 "*"：订阅后可收到全部事件
//
// 当前内置 redis 后端行为：
//   - Emit 发布到 Redis，由各实例（含自身）的订阅循环收到后本地分发，
//     从而实现跨实例广播；新增后端实现 Emitter 接口即可替换。
type Emitter interface {
	// On 订阅指定类型事件，eventType 为 All("*") 时监听全部事件。
	// 返回取消订阅函数，调用即移除该 handler。
	On(eventType string, h Handler) (cancel func())

	// Off 移除指定类型下的全部订阅。
	Off(eventType string)

	// Emit 发布事件。payload 为透明字节载荷。
	// 返回后端传输层错误（local 恒为 nil）。
	Emit(eventType string, payload []byte) error

	// Close 关闭发射器并释放底层资源，应保证幂等。
	Close() error
}
