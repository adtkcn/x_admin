# core/ws WebSocket 连接管理

基于 gorilla/websocket 的 WebSocket 连接管理模块，支持单机与集群部署。

核心设计：
- **事件总线统一走 `core/pubsub`（Redis 后端）**，实现跨实例消息广播；房间数据保存在各实例本地内存，跨实例通信完全通过事件总线完成。
- 提供用户级（`SendToUser`）、房间级（`SendToRoom`）、全局（`SendToAll`）推送，以及跨实例广播的 `CloseUser`/`CloseRoom`/`CloseAll`。
- 集群模式下通过 Redis SET + TTL 心跳做全局唯一在线人数统计。

---

## 一、使用说明

### 1. 初始化（已由 core 自动完成）

初始化在 `core/ws.go` 的 `init()` 中自动完成，业务代码无需手动初始化：

```go
// core/ws.go
var Ws = ws.NewManager()

func init() {
    // 事件总线始终使用 Redis 后端（跨实例广播）
    var rdb = Redis
    var prefix = config.RedisConfig.RedisPrefix

    em, err := pubsub.New(pubsub.Config{
        Backend: pubsub.BackendRedis,
        Prefix:  prefix + "ws:", // 事件频道前缀
    }, Redis)
    if err != nil {
        log.Fatalf("[ws] init pubsub error: %v", err)
    }

    if !config.AppConfig.ClusterMode {
        rdb = nil // 单机模式不使用 Redis 做在线计数（直接统计本实例用户数）
    }

    Ws.Init(em, rdb, prefix)
    go Ws.Start() // 启动注册事件循环
}
```

> 注意：事件总线依赖 Redis，因此**单机部署同样需要 Redis 可用**。

### 2. 建立连接（Controller 层）

升级 HTTP 连接为 WebSocket，构造 `Client` 并设置业务消息回调，注册到管理器后启动读写 goroutine：

```go
func WsHandler(c *gin.Context) {
    uuid := util.ToolsUtil.MakeUuidV7()
    adminId := config.AdminConfig.GetAdminId(c)

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }

    client := ws.NewClient(uuid, adminId, conn, core.Ws)
    client.OnMessage = handleWsMessage // 心跳以外的消息上抛给业务层
    core.Ws.Register <- client

    go client.Write()
    go client.Read()
}
```

### 3. 房间管理

房间通过业务消息动态加入/离开，不在连接时指定。客户端发送 JSON 控制消息，Controller 层路由到 `Manager`：

```json
{"type": "join_room",  "data": "room1"}
{"type": "leave_room", "data": "room1"}
```

```go
func handleWsMessage(client *ws.Client, data []byte) {
    var msg clientMessage
    json.Unmarshal(data, &msg)
    switch msg.Type {
    case "join_room":
        client.Manager.JoinRoom(client.UUID, msg.Data)
    case "leave_room":
        client.Manager.LeaveRoom(client.UUID, msg.Data)
    }
}
```

房间数据仅存本地内存，集群下各实例独立维护自己的房间连接列表；
`SendToRoom`/`CloseRoom` 经事件总线广播到所有实例，各实例各自处理本地连接。

### 4. 消息推送

所有 `SendTo*` 方法自动将消息包装为 `WsResponse{type, data}` 格式：

```go
// 向指定用户推送通知（跨实例：目标用户在本实例或其他实例都能收到）
core.Ws.SendToUser("user123", "notice", map[string]any{
    "id":      1,
    "title":   "新审批",
    "content": "您有一条新的审批待处理",
})
// 前端收到: {"type":"notice","data":{"id":1,"title":"新审批","content":"..."}}

// 向指定房间推送消息
core.Ws.SendToRoom("room456", "chat", map[string]any{
    "from": "user123",
    "text": "hello room",
})
// 前端收到: {"type":"chat","data":{"from":"user123","text":"hello room"}}

// 全局广播
core.Ws.SendToAll("onlineCount", map[string]any{"count": 5})
// 前端收到: {"type":"onlineCount","data":{"count":5}}
```

### 5. 主动关闭连接

`Close*` 通过事件总线广播到所有实例，所有实例同时关闭对应连接：

```go
core.Ws.CloseUser("user123")  // 关闭指定用户的所有连接
core.Ws.CloseRoom("room456")  // 关闭指定房间的所有连接
core.Ws.CloseAll()            // 关闭所有连接
```

### 6. 在线人数

```go
count := core.Ws.GetOnlineCount()
```

- **集群模式**（`ClusterMode: true`）：返回全局唯一在线用户数（各实例在线 uid 经 Redis `SUNION` 合并）。
- **单机模式**（`ClusterMode: false`）：返回本实例唯一用户数（`len(uid_uuids)`）。

### 7. 优雅关闭

```go
core.Ws.Close() // 停心跳、删本实例在线 SET、关闭事件总线
```

---

## 二、技术细节

### 架构与分层

```
Controller 层 (app/controller/ws.go)
  handleWsMessage()  ← Client.OnMessage 回调
    - 解析 JSON 业务消息，路由 join_room / leave_room / ...
    - 调用 Manager.JoinRoom / LeaveRoom
        │
        ▼
基础设施层 (core/ws)
  Client
    - Read(): 心跳 ping→pong；其他消息 → OnMessage 上抛
    - Write(): 发送消息 + 定期 Ping 保活
  Manager
    - 连接注册/注销（Register channel 单 goroutine 循环）
    - JoinRoom / LeaveRoom（房间管理，纯本地）
    - SendToUser / SendToRoom / SendToAll
    - 经 core/pubsub Emitter 广播 → 各实例订阅循环 → 本地分发
  RoomManager
    - rooms / uuid_rooms 内存映射（两种模式共用）
```

事件总线与房间的分工：**事件总线（Redis）负责跨实例广播；房间数据（本地内存）负责本实例内的连接定位**。

### 文件职责

| 文件 | 说明 |
|------|------|
| `client.go` | `Client` 连接：生命周期、心跳（ping/pong）、消息读写、`OnMessage` 上抛 |
| `interface.go` | 业务消息类型与结构（`WsMessageType` / `WsMessage` / `WsResponse`）定义 |
| `manager.go` | `Manager`：连接注册/注销、按用户/房间/全局推送、跨实例广播、在线计数 |
| `room.go` | `RoomManager`：房间 ↔ 连接的内存映射（两种部署模式共用） |

### 消息类型与分发

事件总线以**业务消息类型**为频道名，共 6 种：

| 常量 | 值 | 用途 |
|------|-----|------|
| `MsgTypeUser` | `user` | 向指定用户推送（`localSendToUser`） |
| `MsgTypeRoom` | `room` | 向指定房间推送（`localSendToRoom`） |
| `MsgTypeAll` | `all` | 全局广播（`localSendToAll`） |
| `MsgTypeCloseUser` | `closeUser` | 关闭指定用户连接（`localCloseUser`） |
| `MsgTypeCloseRoom` | `closeRoom` | 关闭指定房间连接（`localCloseRoom`） |
| `MsgTypeCloseAll` | `closeAll` | 关闭所有连接（`localCloseAll`） |

- 发布：`publish(msgType, target, message)` 把消息序列化为
  `WsMessage{Type, Target, Data, NodeID}` 再 `emitter.Emit(string(msgType), payload)`。
- 订阅：`Init` 时 `subscribeEvents()` 按上述 6 种类型调用 `emitter.On`，收到后经
  `onEvent` 反序列化 `WsMessage` 再分发给对应 `local*` 处理函数。
- 事件总线由 `core/pubsub` 提供（mitt 风格 `Emitter`：`On`/`Off`/`Emit`/`Close`），
  当前为 Redis 后端，`Emit` 经 Redis Pub/Sub 广播到所有实例（含发布方自身）。

### 推送与广播流程（以 SendToUser 为例）

```
SendToUser(uid, "chat", data)
  → publish(MsgTypeUser, uid, WsResponse{type:"chat", data})
  → emitter.Emit("user", <WsMessage 字节>)            // 经 Redis 广播到所有实例
  → 各实例 subscribeEvents 的 handler 收到
  → onEvent 反序列化 WsMessage → localSendToUser(uid, json)
  → 查本实例 uid_uuids[uid] → 命中连接则 client.Send，无则忽略
```

集群下每个实例都会收到广播，但只有「持有目标用户/房间连接」的实例才会真正下发，其余实例自然忽略。

### 集群与在线计数

- **事件总线**：始终 Redis Pub/Sub（`{RedisPrefix}ws:*` 频道），无论单机或集群。
- **在线计数（仅 `ClusterMode: true`）**：
  - 每个实例维护 Redis SET `{RedisPrefix}ws:online:{nodeID}`，value 为本实例在线 uid；
    TTL 30s，由 `onlineHeartbeat` 每 10s 刷新（`refreshOnlineSet`）。实例崩溃后 SET 自动过期，不会虚高。
  - `GetOnlineCount` 用 `SCAN {RedisPrefix}ws:online:*` 找到所有实例 SET，再 `SUNION` 得全局唯一用户数（用 `SCAN` 而非 `KEYS` 避免阻塞）。
  - 异常（Redis 不可用）时回退到本实例 `len(uid_uuids)`。
- **单机模式**（`rdb == nil`）：`GetOnlineCount` 直接返回 `len(uid_uuids)`。

### 心跳保活

```
客户端                               服务端
  │──── "ping" (文本消息) ─────────►│ Read goroutine 经 send channel 回复 "pong"
  │◄─── "pong" (文本消息) ──────────│ Write goroutine 统一写出
  │◄─── Ping (WebSocket 帧) ───────│ Write goroutine 每 54s 发送
  │──── Pong (WebSocket 帧) ──────►│ 60s 未收到则断开连接
```

| 参数 | 值 | 说明 |
|------|-----|------|
| `writeWait` | 10s | 写操作超时 |
| `pongWait` | 60s | 等待 Pong 超时 |
| `pingPeriod` | 54s | Ping 发送周期（pongWait × 90%） |
| `maxMessageSize` | 2048B | 客户端消息大小限制 |
| `sendChanSize` | 256 | 发送缓冲区大小 |

### 并发安全

| 操作 | 保护机制 |
|------|----------|
| 连接注册 | 单 goroutine 事件循环消费 `Register` channel（无锁添加） |
| 连接注销 | `closeConn` 经 `sync.Once` 保证只执行一次，直接调用 `unregisterClient`（不经 channel） |
| 连接/用户映射读写 | `Manager.mutex`（`RWMutex`），拷贝指针列表后再释放锁发送 |
| 房间操作 | `RoomManager.mutex`（`RWMutex`） |
| 消息发送 | 非阻塞 `select`，channel 满时丢弃并记录日志 |
| 事件总线关闭 | `core/pubsub` 内部幂等 `Close` |

### 统一消息格式

所有 `SendTo*` 推送的消息统一为 `WsResponse{type, data}`：

| 字段 | 类型 | 说明 |
|------|------|------|
| `type` | string | 业务消息类型，如 `onlineCount`、`notice`、`chat` |
| `data` | any | 消息内容 |

前端通过 `type` 路由处理逻辑，从 `data` 取具体数据。
内部总线传输用 `WsMessage{Type, Target, Data, NodeID}`（`NodeID` 用于标识来源实例）。

### 配置

```yaml
APP:
  ClusterMode: false   # 单机部署（默认）；true 为集群（需 Redis）
REDIS:
  Url: "redis://:@127.0.0.1:6379/0"
  RedisPrefix: "x:"    # 事件频道: "x:ws:*"，在线计数: "x:ws:online:*"
```

### 注意事项 / 已知限制

1. **事件总线依赖 Redis**：移除本地后端后，单机部署也需 Redis 可用。
2. **房间数据仅存本地内存**：不提供跨实例 `GetRoomMembers` / `GetRoomOnlineCount`；如需全局房间统计，可经事件总线心跳聚合。
3. **Redis Pub/Sub 是 fire-and-forget**：不持久化消息，Redis 断线期间的广播会丢失。
4. **发送缓冲区满时丢弃消息**：日志 `[ws] client xxx send buffer full, message dropped`。
5. **客户端非心跳消息全部经 `OnMessage` 上抛**，基础设施层不解析任何业务消息。
6. **`CloseUser`/`CloseRoom`/`CloseAll` 经事件总线广播**：所有实例同时关闭对应连接，房间数据随 `unregisterClient → LeaveAll` 自然清理。
7. **在线计数在集群下存在短暂一致性窗口**：基于 SET + TTL 心跳，崩溃实例最多在 TTL（30s）内仍被计入。
