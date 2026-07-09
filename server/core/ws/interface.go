// Package ws 提供 WebSocket 连接管理，支持单机和集群两种部署模式。
//
// 架构概述：
//
//	Manager 负责管理所有 WebSocket 连接的生命周期（注册/注销/推送）。
//	通过 PubSub 接口抽象消息总线，实现单机与集群的无缝切换：
//	  - 单机模式（LocalPubSub）：使用 Go channel 在进程内传递消息，零额外依赖
//	  - 集群模式（RedisPubSub）：使用 Redis Pub/Sub 跨实例广播消息
//
// 消息流转：
//
//	发送方调用 SendToUser/SendToRoom/SendToAll → 消息发布到 PubSub 总线
//	→ 所有实例的 subscribeLoop 收到消息 → 各自检查本地连接并推送
//
// 配置方式（.env.yaml）：
//
//	APP:
//	  ClusterMode: "local"   # 单机部署（默认）
//	  # ClusterMode: "redis" # 集群部署，需确保 Redis 可用
package ws

import (
	"context"
	"encoding/json"
)

// PubSubMessageType 消息类型，决定消息在接收端的分发逻辑
type PubSubMessageType string

const (
	// MsgTypeUser 向指定用户的所有连接推送，Target 为用户 UID
	MsgTypeUser PubSubMessageType = "user"
	// MsgTypeRoom 向指定房间的所有成员推送，Target 为房间 ID
	MsgTypeRoom PubSubMessageType = "room"
	// MsgTypeAll 向所有在线连接推送，Target 为空
	MsgTypeAll PubSubMessageType = "all"
	// MsgTypeCloseRoom 关闭指定房间的所有连接，Target 为房间 ID
	MsgTypeCloseRoom PubSubMessageType = "close_room"
	// MsgTypeCloseUser 关闭指定用户的所有连接，Target 为用户 UID
	MsgTypeCloseUser PubSubMessageType = "close_user"
	// MsgTypeCloseAll 关闭所有连接，Target 为空
	MsgTypeCloseAll PubSubMessageType = "close_all"
)

// PubSubMessage 跨实例消息结构，在 PubSub 总线中传输。
//
// 集群模式下，消息会被 JSON 序列化后通过 Redis Pub/Sub 广播到所有实例；
// 单机模式下，消息直接通过 Go channel 传递，不经过序列化。
type PubSubMessage struct {
	// Type 消息类型，接收端据此决定调用 localSendToUser/Room/All
	Type PubSubMessageType `json:"type"`
	// Target 推送目标标识：uid（user 类型）或 roomID（room 类型），all 类型时为空字符串
	Target string `json:"target"`
	// Data 序列化后的消息体，由 json.Marshal 生成
	Data json.RawMessage `json:"data"`
	// NodeID 发送节点的唯一标识，由 Manager 初始化时通过 UUID 生成，可用于调试和追踪
	NodeID string `json:"node_id"`
}

// PubSub 消息总线接口，定义消息的发布与订阅能力。
//
// 实现该接口即可扩展新的消息传输方式（如 NATS、Kafka 等）。
// 当前提供两种实现：
//   - LocalPubSub：基于 Go channel，适用于单机部署
//   - RedisPubSub：基于 Redis Pub/Sub，适用于集群部署
type PubSub interface {
	// Publish 发布消息到总线。
	// 单机模式下为非阻塞写入 channel；集群模式下通过 Redis PUBLISH 广播。
	// 消息发布后，所有实例（包括发送者自身）都会收到该消息，
	// 各实例在 subscribeLoop 中检查本地连接并按需推送。
	Publish(ctx context.Context, msg PubSubMessage) error

	// Subscribe 订阅消息，返回只读消息 channel。
	// Manager.Init 时调用，返回的 channel 会被 subscribeLoop 持续监听。
	Subscribe(ctx context.Context) (<-chan PubSubMessage, error)

	// Close 关闭消息总线，释放资源。
	// 应保证幂等性，多次调用不产生错误。
	Close() error
}

// WsResponse 统一的 WebSocket 推送消息格式。
//
// 所有通过 SendToUser/SendToRoom/SendToAll 发送的消息都会自动包装为此格式，
// 前端统一按 {type, data} 解析，通过 type 字段区分消息类型并路由到对应处理逻辑。
//
// 示例：
//
//	{"type": "onlineCount", "data": {"count": 5}}
//	{"type": "notice", "data": {"id": 1, "title": "..."}}
type WsResponse struct {
	Type string `json:"type"` // 消息类型，如 "onlineCount"、"notice"
	Data any    `json:"data"` // 消息内容
}
