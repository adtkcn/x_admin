package ws

import "sync"

// RoomManager 房间管理器，基于内存 map 管理房间与连接的映射关系。
//
// 两种部署模式共用此实现：
//   - 单机模式：房间数据仅在本实例内使用
//   - 集群模式：房间数据仍仅存本地，跨实例通信完全通过 PubSub 完成
//     （SendToRoom/CloseRoom 通过 PubSub 广播，各实例各自处理本地连接）
//
// 数据结构：
//   - rooms: RoomID → UUID 集合，用于按房间查找连接（消息推送）
//   - uuid_rooms: UUID → RoomID 集合，用于按连接查找房间（注销时清理）
type RoomManager struct {
	rooms      map[string]map[string]bool // RoomID → Set<UUID>
	uuid_rooms map[string]map[string]bool // UUID → Set<RoomID>（反向查找，注销时清理）
	mutex      sync.RWMutex
}

// NewRoomManager 创建房间管理器。
func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms:      make(map[string]map[string]bool),
		uuid_rooms: make(map[string]map[string]bool),
	}
}

// Join 将连接加入房间。
func (r *RoomManager) Join(uuid, roomID string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.rooms[roomID] == nil {
		r.rooms[roomID] = make(map[string]bool)
	}
	r.rooms[roomID][uuid] = true

	if r.uuid_rooms[uuid] == nil {
		r.uuid_rooms[uuid] = make(map[string]bool)
	}
	r.uuid_rooms[uuid][roomID] = true
}

// Leave 将连接离开房间。
func (r *RoomManager) Leave(uuid, roomID string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.leaveLocked(uuid, roomID)
}

// leaveLocked 内部离开房间实现（调用方需持锁）。
func (r *RoomManager) leaveLocked(uuid, roomID string) {
	if room, ok := r.rooms[roomID]; ok {
		delete(room, uuid)
		if len(room) == 0 {
			delete(r.rooms, roomID)
		}
	}

	if rooms, ok := r.uuid_rooms[uuid]; ok {
		delete(rooms, roomID)
		if len(rooms) == 0 {
			delete(r.uuid_rooms, uuid)
		}
	}
}

// LeaveAll 连接断开时离开所有房间。
func (r *RoomManager) LeaveAll(uuid string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rooms := r.uuid_rooms[uuid]
	for roomID := range rooms {
		r.leaveLocked(uuid, roomID)
	}
}

// GetClientUUIDs 获取房间内本地连接的 UUID 列表（用于消息推送）。
func (r *RoomManager) GetClientUUIDs(roomID string) []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	room := r.rooms[roomID]
	if room == nil {
		return nil
	}
	uuids := make([]string, 0, len(room))
	for uuid := range room {
		uuids = append(uuids, uuid)
	}
	return uuids
}
