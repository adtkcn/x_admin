package ws_util

import (
	"sync"
)

type Manager struct {
	clients    map[string]*Client         // 所有客户端：key=clientID
	rooms      map[string]map[string]bool // 群组：roomID -> {clientID: true}
	Register   chan *Client
	UnRegister chan *Client
	mutex      sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients:    make(map[string]*Client),
		rooms:      make(map[string]map[string]bool),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
}

// 启动管理器协程
func (m *Manager) Start() {
	for {
		select {
		case client := <-m.Register:
			m.mutex.Lock()
			m.clients[client.ID] = client

			// 加入房间
			if client.RoomID != "" {
				if _, ok := m.rooms[client.RoomID]; !ok {
					m.rooms[client.RoomID] = make(map[string]bool)
				}
				m.rooms[client.RoomID][client.ID] = true
			}
			m.mutex.Unlock()

		case client := <-m.UnRegister:
			m.mutex.Lock()
			if _, ok := m.clients[client.ID]; ok {
				// 删除客户端
				delete(m.clients, client.ID)

				// 退出房间
				if client.RoomID != "" {
					if room, ok := m.rooms[client.RoomID]; ok {
						delete(room, client.ID)
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
func (m *Manager) SendToUser(clientID string, message []byte) {
	m.mutex.RLock()
	client, exists := m.clients[clientID]
	m.mutex.RUnlock()

	if exists && client != nil {
		select {
		case client.send <- message:
		default:
			// 通道满，丢弃或记录日志
			close(client.send)
		}
	}
}

// 群推：向指定 roomID 的所有成员发送消息
func (m *Manager) SendToRoom(roomID string, message []byte) {
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
		case client.send <- message:
		default:
			close(client.send)
		}
	}
}

// 全推：向所有在线用户发送消息
func (m *Manager) SendToAll(message []byte) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, client := range m.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
		}
	}
}
