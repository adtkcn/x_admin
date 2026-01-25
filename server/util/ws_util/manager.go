package ws_util

import (
	"encoding/json"
	"sync"
)

type Manager struct {
	clients    map[string]*Client         // 所有客户端：key=clientID
	uid_uuids  map[string][]string        // uid -> [uuid1, uuid2, ...]
	rooms      map[string]map[string]bool // 群组：roomID -> {clientID: true}
	Register   chan *Client
	UnRegister chan *Client
	mutex      sync.RWMutex
}

func NewManager() *Manager {
	var manager = Manager{
		clients:    make(map[string]*Client),
		uid_uuids:  make(map[string][]string),
		rooms:      make(map[string]map[string]bool),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
	go manager.Start()

	return &manager
}

// 启动管理器协程
func (m *Manager) Start() {
	for {
		select {
		case client := <-m.Register:
			m.mutex.Lock()
			m.clients[client.UUID] = client
			// 加入用户
			if client.Uid != "" {
				m.uid_uuids[client.Uid] = append(m.uid_uuids[client.Uid], client.UUID)
			}

			// 加入房间
			if client.RoomID != "" {
				if _, ok := m.rooms[client.RoomID]; !ok {
					m.rooms[client.RoomID] = make(map[string]bool)
				}
				m.rooms[client.RoomID][client.UUID] = true
			}
			m.mutex.Unlock()

		case client := <-m.UnRegister:
			m.mutex.Lock()
			if _, ok := m.clients[client.UUID]; ok {
				// 删除客户端
				delete(m.clients, client.UUID)
				// 删除用户对应的UUID
				if client.Uid != "" {
					for i, uuid := range m.uid_uuids[client.Uid] {
						if uuid == client.UUID {
							m.uid_uuids[client.Uid] = append(m.uid_uuids[client.Uid][:i], m.uid_uuids[client.Uid][i+1:]...)
							break
						}
					}
				}

				// 退出房间
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

// 单推：向指定 clientID 发送消息
func (m *Manager) SendToUser(uid string, message any) {
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return
	}

	m.mutex.RLock()
	clientIDs, exists := m.uid_uuids[uid]
	m.mutex.RUnlock()

	if exists && clientIDs != nil {
		for _, clientID := range clientIDs {
			if client, ok := m.clients[clientID]; ok {
				select {
				case client.send <- jsonMsg:
				default:
					// 通道满，丢弃或记录日志
					close(client.send)
				}
			}
		}
	}
}

// 群推：向指定 roomID 的所有成员发送消息
func (m *Manager) SendToRoom(roomID string, message any) {
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return
	}

	m.mutex.RLock()
	room, exists := m.rooms[roomID]
	if !exists {
		m.mutex.RUnlock()
		return
	}

	var clientsToSend []*Client
	for clientID := range room {
		if client, ok := m.clients[clientID]; ok {
			clientsToSend = append(clientsToSend, client)
		}
	}
	m.mutex.RUnlock()

	for _, client := range clientsToSend {
		select {
		case client.send <- jsonMsg:
		default:
			close(client.send)
		}
	}
}

// 全推：向所有在线用户发送消息
func (m *Manager) SendToAll(message any) {
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		return
	}
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, client := range m.clients {
		select {
		case client.send <- jsonMsg:
		default:
			close(client.send)
		}
	}
}

// 获取当前在线用户数
func (m *Manager) GetOnlineCount() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return len(m.clients)
}

// 获取房间内的在线用户数
func (m *Manager) GetRoomOnlineCount(roomID string) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if room, ok := m.rooms[roomID]; ok {
		return len(room)
	}
	return 0
}

// 关闭所有连接
func (m *Manager) CloseAll() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, client := range m.clients {
		client.Close()
	}
}

// 关闭指定房间的所有连接
func (m *Manager) CloseRoom(roomID string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if room, ok := m.rooms[roomID]; ok {
		for clientID := range room {
			if client, ok := m.clients[clientID]; ok {
				client.Close()
			}
		}
		delete(m.rooms, roomID)
	}
}

// 关闭指定用户的连接
func (m *Manager) CloseUser(clientID string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if client, ok := m.clients[clientID]; ok {
		client.Close()
		delete(m.clients, clientID)
	}
}
