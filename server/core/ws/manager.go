package ws

import (
	"encoding/json"
	"sync"
)

// Manager WebSocket 连接管理器
type Manager struct {
	clients    map[string]*Client         // 所有客户端：key=UUID
	uid_uuids  map[string]map[string]bool // uid -> {uuid: true}（改用 map 提升删除效率）
	rooms      map[string]map[string]bool // 群组：roomID -> {uuid: true}
	Register   chan *Client
	UnRegister chan *Client
	mutex      sync.RWMutex
}

// NewManager 创建管理器（不自动启动，由外部显式调用 Start）
func NewManager() *Manager {
	return &Manager{
		clients:    make(map[string]*Client),
		uid_uuids:  make(map[string]map[string]bool),
		rooms:      make(map[string]map[string]bool),
		Register:   make(chan *Client, 64),
		UnRegister: make(chan *Client, 64),
	}
}

// Start 启动管理器事件循环（由 core/ws.go init 中 go 调用）
func (m *Manager) Start() {
	for {
		select {
		case client := <-m.Register:
			m.mutex.Lock()
			m.clients[client.UUID] = client

			if client.Uid != "" {
				if m.uid_uuids[client.Uid] == nil {
					m.uid_uuids[client.Uid] = make(map[string]bool)
				}
				m.uid_uuids[client.Uid][client.UUID] = true
			}

			if client.RoomID != "" {
				if m.rooms[client.RoomID] == nil {
					m.rooms[client.RoomID] = make(map[string]bool)
				}
				m.rooms[client.RoomID][client.UUID] = true
			}
			m.mutex.Unlock()

		case client := <-m.UnRegister:
			m.mutex.Lock()
			if _, ok := m.clients[client.UUID]; ok {
				delete(m.clients, client.UUID)

				if client.Uid != "" {
					if uuids, exists := m.uid_uuids[client.Uid]; exists {
						delete(uuids, client.UUID)
						if len(uuids) == 0 {
							delete(m.uid_uuids, client.Uid)
						}
					}
				}

				if client.RoomID != "" {
					if room, ok := m.rooms[client.RoomID]; ok {
						delete(room, client.UUID)
						if len(room) == 0 {
							delete(m.rooms, client.RoomID)
						}
					}
				}
			}
			m.mutex.Unlock()
		}
	}
}

// SendToUser 向指定用户的所有连接推送
func (m *Manager) SendToUser(uid string, message any) {
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return
	}

	m.mutex.RLock()
	uuids := m.uid_uuids[uid]
	// 在锁内拷贝 client 指针列表，避免释放锁后数据竞争
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

// SendToRoom 向指定房间的所有成员推送
func (m *Manager) SendToRoom(roomID string, message any) {
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return
	}

	m.mutex.RLock()
	room := m.rooms[roomID]
	if room == nil {
		m.mutex.RUnlock()
		return
	}
	var clientsToSend []*Client
	for uuid := range room {
		if c, ok := m.clients[uuid]; ok {
			clientsToSend = append(clientsToSend, c)
		}
	}
	m.mutex.RUnlock()

	for _, c := range clientsToSend {
		c.Send(jsonMsg)
	}
}

// SendToAll 向所有在线连接推送
func (m *Manager) SendToAll(message any) {
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return
	}

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

// GetOnlineCount 获取当前在线连接数
func (m *Manager) GetOnlineCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.clients)
}

// GetRoomOnlineCount 获取房间内在线人数
func (m *Manager) GetRoomOnlineCount(roomID string) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	if room, ok := m.rooms[roomID]; ok {
		return len(room)
	}
	return 0
}

// CloseAll 关闭所有连接（不持锁调用 client.Close，避免死锁）
func (m *Manager) CloseAll() {
	m.mutex.Lock()
	all := make([]*Client, 0, len(m.clients))
	for _, c := range m.clients {
		all = append(all, c)
	}
	m.mutex.Unlock()

	for _, c := range all {
		c.Close()
	}
}

// CloseRoom 关闭指定房间的所有连接
func (m *Manager) CloseRoom(roomID string) {
	m.mutex.Lock()
	room := m.rooms[roomID]
	if room == nil {
		m.mutex.Unlock()
		return
	}
	var toClose []*Client
	for uuid := range room {
		if c, ok := m.clients[uuid]; ok {
			toClose = append(toClose, c)
		}
	}
	delete(m.rooms, roomID)
	m.mutex.Unlock()

	for _, c := range toClose {
		c.Close()
	}
}

// CloseUser 关闭指定用户的所有连接
func (m *Manager) CloseUser(uid string) {
	m.mutex.Lock()
	uuids := m.uid_uuids[uid]
	if uuids == nil {
		m.mutex.Unlock()
		return
	}
	var toClose []*Client
	for uuid := range uuids {
		if c, ok := m.clients[uuid]; ok {
			toClose = append(toClose, c)
		}
	}
	m.mutex.Unlock()

	for _, c := range toClose {
		c.Close()
	}
}
