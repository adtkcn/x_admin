# WebSocket 实时通信

后端 WebSocket 由 `core/ws` 包实现，提供连接管理、按用户/房间/全局推送、集群跨实例广播、全局在线人数统计与连接强制关闭能力。

## 一、核心结构

| 类型 | 位置 | 说明 |
| --- | --- | --- |
| `Manager` | `core/ws/manager.go` | 连接管理器（注册/注销、推送、房间、集群） |
| `Client` | `core/ws/client.go` | 单个连接（读写协程，`sync.Once` 幂等关闭） |
| `RoomManager` | `core/ws/room.go` | 房间管理（纯本地内存） |
| `WsMessage` / `WsResponse` | `core/ws/types.go` | 内部传输结构 / 推送给前端的响应结构 |

全局实例：`core/ws.go` 中 `var Ws = ws.NewManager()`。

## 二、初始化

```go
// core/ws.go 启动流程（示意）
core.Ws = ws.NewManager()
core.Ws.Init(emitter, redisClient, prefix)  // emitter 来自 core/pubsub，redisClient 为 nil 时单机
go core.Ws.Start()                          // 启动注册循环
```

- `NewManager()`：创建空管理器（含 `RoomManager`）。
- `Init(em pubsub.Emitter, rdb *redis.Client, prefix string)`：
  - 通过 `emitter` 订阅事件总线实现跨实例广播；
  - `rdb != nil` 进入**集群模式**（启动在线心跳）；`rdb == nil` 为**单机模式**。
- `Start()`：阻塞消费 `Register` channel，完成连接注册（注销由 `closeConn` 直接调用，不绕 channel）。

## 三、消息类型

`WsMessageType`（内部事件）：`MsgTypeUser` / `MsgTypeRoom` / `MsgTypeAll` / `MsgTypeCloseRoom` / `MsgTypeCloseUser` / `MsgTypeCloseAll`。`Manager` 在 `Init` 时按这些类型订阅事件总线，收到后本地分发（跨实例由 pubsub 广播到所有实例）。

- `WsMessage{Type, Target, Data []byte, NodeID}`：pubsub 传输载荷（JSON）。
- `WsResponse{Type string, Data any}`：最终推送给前端的消息包装。

## 四、推送 API

```go
func (m *Manager) SendToUser(uid, msgType string, data any)   // 指定用户所有连接
func (m *Manager) SendToRoom(roomID, msgType string, data any) // 房间所有成员
func (m *Manager) SendToAll(msgType string, data any)          // 全局广播
```

- 调用即 `publish`：先 `json.Marshal` 为 `WsResponse`，再包装成 `WsMessage` 经 `emitter.Emit` 发出。
- **集群语义**：`Emit` 经 pubsub 广播到所有实例（含本实例），各实例本地只推自己持有的连接；不在本实例的用户/房间由对应实例推送。因此对调用方而言是「发一次、全局送达」。

示例（定时任务广播在线人数）：

```go
core.Ws.SendToAll("onlineCount", map[string]any{"count": count})
```

## 五、房间与连接控制

```go
func (m *Manager) JoinRoom(uuid, roomID string)
func (m *Manager) LeaveRoom(uuid, roomID string)
func (m *Manager) CloseAll()    // 广播关闭所有实例的连接
func (m *Manager) CloseRoom(roomID string)
func (m *Manager) CloseUser(uid string)
func (m *Manager) Close()
```

- `JoinRoom/LeaveRoom` 仅操作本地 `RoomManager`（房间数据不跨实例）。
- `CloseAll/CloseRoom/CloseUser` 经 pubsub 广播，所有实例同时执行（跨实例关闭）。
- `Close()`：停止心跳、删除本实例在线 SET、关闭 emitter。

## 六、在线人数统计

```go
func (m *Manager) GetOnlineCount() int  // 返回唯一用户数（非连接数）
```

- **单机模式**：直接返回本实例 `len(uid_uuids)`。
- **集群模式**：`SCAN` 所有实例在线 SET（`{prefix}ws:online:{nodeID}`，TTL 30s、心跳 10s 刷新），`SUNION` 求并集得到全局唯一在线用户数。实例崩溃后 SET 自动过期，不会虚高。

## 七、HTTP 路由

`routes/api.go` 中通过 `upgrader`（如 `github.com/gorilla/websocket`）将 `/ws` 升级为 WebSocket，握手后创建 `Client` 并写入 `Manager.Register`。同一 uid 可存在多个连接（uuid 维度区分）。

## 八、扩展

- 新增业务推送：直接 `core.Ws.SendToUser/SendToRoom/SendToAll(msgType, data)`，前端按 `msgType` 区分处理。
- 新增集群广播事件：在 `subscribeEvents` 中 `m.emitter.On(string(MsgTypeX), m.onEvent(...))` 注册分发逻辑，并新增对应 `publish` 入口。

> 依赖 `core/pubsub`（事件总线）与 Redis（集群模式）。详见 [事件总线 pubsub](./事件总线pubsub.md)。
