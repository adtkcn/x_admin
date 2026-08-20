// Package ws 提供 WebSocket 连接管理，支持单机（local）和集群（redis）两种部署模式。
//
// 集群模式下，多个服务实例通过 core/pubsub 事件总线（mitt 风格 Emitter）同步连接状态。
// pubsub 包只负责按事件类型透明传输 []byte 载荷，不感知业务语义；本包在此之上定义
// WebSocket 业务消息（WsMessageType / WsMessage），发送时以消息类型为事件名 Emit、
// 订阅时按类型 On 并反序列化载荷。
package ws

import "encoding/json/jsontext"

// WsMessageType WebSocket 业务消息类型，决定接收端的分发逻辑。
// 属于 ws 业务语义，故定义在 ws 包而非 pubsub 包。
type WsMessageType string

const (
	// MsgTypeUser 向指定用户的所有连接推送，Target 为用户 UID
	MsgTypeUser WsMessageType = "user"
	// MsgTypeRoom 向指定房间的所有成员推送，Target 为房间 ID
	MsgTypeRoom WsMessageType = "room"
	// MsgTypeAll 向所有在线连接推送，Target 为空
	MsgTypeAll WsMessageType = "all"
	// MsgTypeCloseRoom 关闭指定房间的所有连接，Target 为房间 ID
	MsgTypeCloseRoom WsMessageType = "close_room"
	// MsgTypeCloseUser 关闭指定用户的所有连接，Target 为用户 UID
	MsgTypeCloseUser WsMessageType = "close_user"
	// MsgTypeCloseAll 关闭所有连接，Target 为空
	MsgTypeCloseAll WsMessageType = "close_all"
)

// WsMessage 跨实例 WebSocket 业务消息。
// ws 层负责将其序列化为 []byte 后，通过纯 pubsub 总线透明传输。
type WsMessage struct {
	// Type 消息类型，接收端据此决定分发逻辑
	Type WsMessageType `json:"type"`
	// Target 推送目标标识：uid（user 类型）或 roomID（room 类型），all 类型时为空字符串
	Target string `json:"target"`
	// Data 序列化后的消息体（保留原始 JSON 文本，v2 下用 jsontext.Value 等价 RawMessage）
	Data jsontext.Value `json:"data"`
	// NodeID 发送节点的唯一标识，用于调试和追踪
	NodeID string `json:"node_id"`
}

// WsResponse WebSocket 消息响应格式，作为消息体推送到客户端。
type WsResponse struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
