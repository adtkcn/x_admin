package ws

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"x_admin/core/pubsub"
)

// 在线用户统计相关常量。
//
// 集群模式下，每个实例维护一个 Redis SET 存储本实例的在线 uid：
//
//	SET {prefix}ws:online:{nodeID}
//	  - value: 该实例上所有在线 uid
//	  - TTL: 30s，由心跳 goroutine 每 10s 刷新
//
// 实例崩溃后 SET 自动过期，不会导致在线人数虚高。
// GetOnlineCount 通过 SUNION 所有实例的 SET 获取全局唯一在线用户数。
const (
	onlineKeyPrefix       = "ws:online:"
	onlineTTL             = 30 * time.Second
	onlineHeartbeatPeriod = 10 * time.Second
)

// Manager WebSocket 连接管理器，负责管理所有 WebSocket 连接的生命周期。
//
// 核心职责：
//   - 连接注册/注销：通过 Register channel 异步注册，closeConn 直接调用 unregisterClient
//   - 消息推送：支持按用户（SendToUser）、按房间（SendToRoom）、全局广播（SendToAll）
//   - 集群支持：通过 PubSub 接口实现跨实例消息广播，通过 Redis SET + TTL 心跳实现全局在线计数
//
// 房间数据仅存本地内存，集群下跨实例通信完全通过 PubSub 完成。
type Manager struct {
	// 连接存储
	clients   map[string]*Client         // UUID → *Client 映射
	uid_uuids map[string]map[string]bool // UID → UUID 集合映射（支持按用户查找）

	// 连接注册
	Register chan *Client // 注册 channel，新连接通过此 channel 加入管理器
	mutex    sync.RWMutex // 读写锁，保护连接存储

	// 房间管理（纯本地，两种模式共用）
	roomManager *RoomManager

	// 集群支持
	nodeID          string             // 当前实例唯一标识（UUID）
	emitter         pubsub.Emitter     // 事件总线（mitt 风格，local / redis 多后端）
	redisClient     *redis.Client      // 集群模式下的 Redis 客户端（nil 表示单机模式）
	prefix          string             // Redis key 前缀
	heartbeatCancel context.CancelFunc // 心跳 goroutine 取消函数
}

// NewManager 创建 WebSocket 连接管理器。
func NewManager() *Manager {
	return &Manager{
		clients:     make(map[string]*Client),
		uid_uuids:   make(map[string]map[string]bool),
		Register:    make(chan *Client, 512),
		roomManager: NewRoomManager(),
	}
}

// Init 初始化事件总线和 Redis 客户端。
func (m *Manager) Init(em pubsub.Emitter, rdb *redis.Client, prefix string) {
	m.emitter = em
	m.redisClient = rdb
	m.prefix = prefix
	m.nodeID = uuid.NewString()

	// 按业务消息类型订阅事件，收到后反序列化并分发到本地连接
	m.subscribeEvents()

	// 集群模式：启动在线心跳
	if m.redisClient != nil {
		ctx, cancel := context.WithCancel(context.Background())
		m.heartbeatCancel = cancel
		go m.onlineHeartbeat(ctx)
	}
}

// onlineHeartbeat 定期将本实例在线 uid 同步到 Redis SET，并刷新 TTL。
// 实例崩溃后 SET 自动过期，不会导致在线人数虚高。
func (m *Manager) onlineHeartbeat(ctx context.Context) {
	m.refreshOnlineSet()
	ticker := time.NewTicker(onlineHeartbeatPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.refreshOnlineSet()
		}
	}
}

// refreshOnlineSet 将本实例当前在线 uid 列表写入 Redis SET 并刷新 TTL。
func (m *Manager) refreshOnlineSet() {
	m.mutex.RLock()
	uids := make([]any, 0, len(m.uid_uuids))
	for uid := range m.uid_uuids {
		uids = append(uids, uid)
	}
	m.mutex.RUnlock()

	ctx := context.Background()
	key := m.prefix + onlineKeyPrefix + m.nodeID
	pipe := m.redisClient.Pipeline()
	pipe.Del(ctx, key)
	if len(uids) > 0 {
		pipe.SAdd(ctx, key, uids...)
	}
	pipe.Expire(ctx, key, onlineTTL)
	if _, err := pipe.Exec(ctx); err != nil {
		log.Printf("[ws] refresh online set error: %v", err)
	}
}

// subscribeEvents 按业务消息类型订阅事件总线，收到后反序列化并分发到本地连接。
// 每种类型对应一个 handler，替代原先单通道 + switch 的分发方式。
func (m *Manager) subscribeEvents() {
	m.emitter.On(string(MsgTypeUser), m.onEvent(func(msg WsMessage) {
		m.localSendToUser(msg.Target, []byte(msg.Data))
	}))
	m.emitter.On(string(MsgTypeRoom), m.onEvent(func(msg WsMessage) {
		m.localSendToRoom(msg.Target, []byte(msg.Data))
	}))
	m.emitter.On(string(MsgTypeAll), m.onEvent(func(msg WsMessage) {
		m.localSendToAll([]byte(msg.Data))
	}))
	m.emitter.On(string(MsgTypeCloseRoom), m.onEvent(func(msg WsMessage) {
		m.localCloseRoom(msg.Target)
	}))
	m.emitter.On(string(MsgTypeCloseUser), m.onEvent(func(msg WsMessage) {
		m.localCloseUser(msg.Target)
	}))
	m.emitter.On(string(MsgTypeCloseAll), m.onEvent(func(msg WsMessage) {
		m.localCloseAll()
	}))
}

// onEvent 将「反序列化 WsMessage 载荷」的样板逻辑包装为 pubsub.Handler。
func (m *Manager) onEvent(fn func(WsMessage)) pubsub.Handler {
	return func(payload []byte) {
		var msg WsMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			log.Printf("[ws] event payload unmarshal error: %v", err)
			return
		}
		fn(msg)
	}
}

// Start 启动管理器事件循环，处理连接注册。
// 注销由 closeConn 直接调用 unregisterClient 完成，无需通过 channel。
func (m *Manager) Start() {
	for client := range m.Register {
		m.registerClient(client)
	}
}

// registerClient 处理客户端注册。
func (m *Manager) registerClient(client *Client) {
	m.mutex.Lock()
	m.clients[client.UUID] = client

	if client.Uid != "" {
		if m.uid_uuids[client.Uid] == nil {
			m.uid_uuids[client.Uid] = make(map[string]bool)
		}
		m.uid_uuids[client.Uid][client.UUID] = true
	}
	m.mutex.Unlock()
}

// unregisterClient 处理客户端注销，清理所有关联数据。
// 由 closeConn 直接调用，sync.Once 保证幂等。
//
// 清理流程：
//  1. 从 clients map 中移除连接
//  2. 从 uid_uuids 中移除 UUID，如果该 UID 无其他连接则删除 key
//  3. 通过 RoomManager.LeaveAll 清理该连接所属的所有房间
func (m *Manager) unregisterClient(client *Client) {
	m.mutex.Lock()
	if _, ok := m.clients[client.UUID]; !ok {
		m.mutex.Unlock()
		return
	}
	delete(m.clients, client.UUID)

	if client.Uid != "" {
		if uuids, exists := m.uid_uuids[client.Uid]; exists {
			delete(uuids, client.UUID)
			if len(uuids) == 0 {
				delete(m.uid_uuids, client.Uid)
			}
		}
	}
	m.mutex.Unlock()

	// 清理房间关联（纯本地操作）
	m.roomManager.LeaveAll(client.UUID)
}

// JoinRoom 将指定连接加入房间。
func (m *Manager) JoinRoom(uuid, roomID string) {
	m.roomManager.Join(uuid, roomID)
}

// LeaveRoom 将指定连接离开房间。
func (m *Manager) LeaveRoom(uuid, roomID string) {
	m.roomManager.Leave(uuid, roomID)
}

// publish 以业务消息类型为事件名，将消息发布到事件总线。
func (m *Manager) publish(msgType WsMessageType, target string, message any) {
	if m.emitter == nil {
		return
	}
	data, err := json.Marshal(message)
	if err != nil {
		return
	}
	payload, err := json.Marshal(WsMessage{
		Type:   msgType,
		Target: target,
		Data:   data,
		NodeID: m.nodeID,
	})
	if err != nil {
		return
	}
	if err := m.emitter.Emit(string(msgType), payload); err != nil {
		log.Printf("[ws] emitter emit error: %v", err)
	}
}

// SendToUser 向指定用户的所有连接推送消息。
// 消息会自动包装为 WsResponse{type, data} 格式后发送。
func (m *Manager) SendToUser(uid, msgType string, data any) {
	m.publish(MsgTypeUser, uid, WsResponse{Type: msgType, Data: data})
}

// SendToRoom 向指定房间的所有成员推送消息。
func (m *Manager) SendToRoom(roomID, msgType string, data any) {
	m.publish(MsgTypeRoom, roomID, WsResponse{Type: msgType, Data: data})
}

// SendToAll 向所有在线连接推送消息（全局广播）。
func (m *Manager) SendToAll(msgType string, data any) {
	m.publish(MsgTypeAll, "", WsResponse{Type: msgType, Data: data})
}

// localSendToUser 本地推送：向本实例内指定用户的所有连接推送。
func (m *Manager) localSendToUser(uid string, jsonMsg []byte) {
	m.mutex.RLock()
	uuids := m.uid_uuids[uid]
	var clientsToSend []*Client
	for uuid := range uuids {
		if c, ok := m.clients[uuid]; ok {
			clientsToSend = append(clientsToSend, c)
		}
	}
	m.mutex.RUnlock()

	for _, c := range clientsToSend {
		c.Send(jsonMsg)
	}
}

// localSendToRoom 本地推送：向本实例内指定房间的所有连接推送。
func (m *Manager) localSendToRoom(roomID string, jsonMsg []byte) {
	uuids := m.roomManager.GetClientUUIDs(roomID)
	if len(uuids) == 0 {
		return
	}

	m.mutex.RLock()
	var clientsToSend []*Client
	for _, uuid := range uuids {
		if c, ok := m.clients[uuid]; ok {
			clientsToSend = append(clientsToSend, c)
		}
	}
	m.mutex.RUnlock()

	for _, c := range clientsToSend {
		c.Send(jsonMsg)
	}
}

// localSendToAll 本地推送：向本实例内所有在线连接推送。
func (m *Manager) localSendToAll(jsonMsg []byte) {
	m.mutex.RLock()
	var clientsToSend []*Client
	for _, c := range m.clients {
		clientsToSend = append(clientsToSend, c)
	}
	m.mutex.RUnlock()

	for _, c := range clientsToSend {
		c.Send(jsonMsg)
	}
}

// localCloseRoom 关闭本实例内指定房间的所有连接。
func (m *Manager) localCloseRoom(roomID string) {
	uuids := m.roomManager.GetClientUUIDs(roomID)

	m.mutex.RLock()
	var toClose []*Client
	for _, uuid := range uuids {
		if c, ok := m.clients[uuid]; ok {
			toClose = append(toClose, c)
		}
	}
	m.mutex.RUnlock()

	for _, c := range toClose {
		c.Close()
	}
}

// localCloseUser 关闭本实例内指定用户的所有连接。
func (m *Manager) localCloseUser(uid string) {
	m.mutex.RLock()
	uuids := m.uid_uuids[uid]
	var toClose []*Client
	for uuid := range uuids {
		if c, ok := m.clients[uuid]; ok {
			toClose = append(toClose, c)
		}
	}
	m.mutex.RUnlock()

	for _, c := range toClose {
		c.Close()
	}
}

// localCloseAll 关闭本实例内所有连接。
func (m *Manager) localCloseAll() {
	m.mutex.RLock()
	all := make([]*Client, 0, len(m.clients))
	for _, c := range m.clients {
		all = append(all, c)
	}
	m.mutex.RUnlock()

	for _, c := range all {
		c.Close()
	}
}

// GetOnlineCount 获取在线用户数（唯一用户数，非连接数）。
//
// 单机模式：返回本实例的唯一用户数 len(uid_uuids)。
// 集群模式：通过 Redis SCAN 找到所有实例的在线 SET，SUNION 后获取全局唯一用户数。
// 使用 SCAN 替代 KEYS，避免阻塞 Redis。
// 实例崩溃后其 SET 自动过期，不会导致计数虚高。
func (m *Manager) GetOnlineCount() int {
	if m.redisClient != nil {
		ctx := context.Background()
		pattern := m.prefix + onlineKeyPrefix + "*"
		var keys []string
		var cursor uint64
		for {
			ks, next, err := m.redisClient.Scan(ctx, cursor, pattern, 100).Result()
			if err != nil {
				log.Printf("[ws] scan online keys error: %v", err)
				return m.localOnlineCount()
			}
			keys = append(keys, ks...)
			cursor = next
			if cursor == 0 {
				break
			}
		}
		if len(keys) == 0 {
			return 0
		}
		members, err := m.redisClient.SUnion(ctx, keys...).Result()
		if err != nil {
			log.Printf("[ws] get online union error: %v", err)
			return m.localOnlineCount()
		}
		return len(members)
	}
	return m.localOnlineCount()
}

// localOnlineCount 返回本实例的唯一在线用户数。
func (m *Manager) localOnlineCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.uid_uuids)
}

// CloseAll 关闭所有连接（通过 PubSub 广播，所有实例同时关闭）。
func (m *Manager) CloseAll() {
	m.publish(MsgTypeCloseAll, "", nil)
}

// CloseRoom 关闭指定房间的所有连接（通过 PubSub 广播，所有实例同时关闭）。
func (m *Manager) CloseRoom(roomID string) {
	m.publish(MsgTypeCloseRoom, roomID, nil)
}

// CloseUser 关闭指定用户的所有连接（通过 PubSub 广播，所有实例同时关闭）。
func (m *Manager) CloseUser(uid string) {
	m.publish(MsgTypeCloseUser, uid, nil)
}

// Close 优雅关闭：停止心跳、删除本实例在线 SET、关闭 PubSub。
func (m *Manager) Close() {
	if m.heartbeatCancel != nil {
		m.heartbeatCancel()
	}
	if m.redisClient != nil {
		m.redisClient.Del(context.Background(), m.prefix+onlineKeyPrefix+m.nodeID)
	}
	if m.emitter != nil {
		m.emitter.Close()
	}
}
