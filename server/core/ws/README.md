# WebSocket 连接管理模块

支持**单机**和**集群**两种部署模式的 WebSocket 连接管理模块。

## 架构概览

```
                          ┌─────────────────────────────────────────┐
                          │             Manager                      │
                          │                                          │
   Client A ──ws──►       │  clients     (UUID → *Client)            │
   Client B ──ws──►       │  uid_uuids   (UID  → Set<UUID>)          │
   Client C ──ws──►       │  rooms       (RoomID → Set<UUID>)  本地   │
                          │  uuid_rooms  (UUID → Set<RoomID>)  本地   │
                          │                                          │
                          │  SendToUser / SendToRoom / SendToAll     │
                          │      │ 自动包装为 WsResponse{type,data}   │
                          │      ▼                                   │
                          │  PubSub.Publish()                        │
                          └──────┬───────────────────────────────────┘
                                 │
                 ┌───────────────┼───────────────┐
                 ▼                               ▼
         ┌──────────────┐               ┌──────────────┐
         │ LocalPubSub  │               │ RedisPubSub  │
         │  (单机模式)   │               │  (集群模式)    │
         │              │               │              │
         │ Go channel   │               │ Redis Pub/Sub│
         │ 零序列化开销   │               │ 跨实例广播     │
         └──────────────┘               └──────────────┘
```

> **房间数据仅存本地内存**，两种模式共用同一套 `RoomManager` 实现。
> 集群下跨实例通信完全通过 PubSub 完成，不依赖 Redis SET。

### 分层架构

```
┌─ Controller 层 (app/controller/ws.go) ─────────────────────┐
│  handleWsMessage()                                         │
│  - 解析 JSON 业务消息: {"type":"join_room","data":"room1"}  │
│  - 路由 join_room / leave_room / ...                       │
│  - 调用 Manager.JoinRoom / LeaveRoom                       │
│  - 可扩展更多业务消息类型                                     │
└───────────┬────────────────────────────────────────────────┘
            │ Client.OnMessage 回调
            ▼
┌─ 基础设施层 (core/ws) ──────────────────────────────────────┐
│  Client.Read()                                              │
│  - 心跳 ping→pong（内部处理）                                │
│  - 其他消息 → 调用 OnMessage 回调上抛                        │
│                                                             │
│  Manager                                                    │
│  - 连接注册/注销                                             │
│  - JoinRoom / LeaveRoom（房间管理，纯本地）                   │
│  - SendToUser / SendToRoom / SendToAll（推送原语）           │
│  - PubSub（集群消息广播抽象）                                 │
└─────────────────────────────────────────────────────────────┘
```

### 消息流转

```
SendToRoom("room1", "chat", data)
  → publish(MsgTypeRoom, "room1", WsResponse{type:"chat", data})
  → PubSub.Publish(PubSubMessage)
  → 所有实例的 subscribeLoop 收到消息
  → localSendToRoom("room1", jsonMsg)
  → 查本地 rooms["room1"] → 有连接则推送，无则忽略

CloseRoom("room1")
  → publish(MsgTypeCloseRoom, "room1")
  → 所有实例的 subscribeLoop 收到消息
  → localCloseRoom("room1") → 关闭本地 rooms["room1"] 中的连接
  → 各连接 unregisterClient → LeaveAll 自然清理房间数据
```

### 统一消息格式

所有通过 `SendTo*` 发送的消息自动包装为 `WsResponse{type, data}` 格式：

```json
{"type": "onlineCount", "data": {"count": 5}}
{"type": "notice",       "data": {"id": 1, "title": "...", "content": "..."}}
```

## 文件说明

| 文件 | 说明 |
|------|------|
| `client.go` | WebSocket 客户端连接，管理单个连接的读写、心跳、OnMessage 回调 |
| `manager.go` | 连接管理器，管理连接注册/注销、消息推送、集群在线计数 |
| `interface.go` | PubSub 接口定义、PubSubMessage / WsResponse 消息结构 |
| `pubsub_local.go` | 本地 channel 实现（单机模式） |
| `pubsub_redis.go` | Redis Pub/Sub 实现（集群模式） |
| `room.go` | 房间管理器（纯本地内存，两种模式共用） |

## 配置

在 `.env.yaml` 中配置集群模式：

```yaml
APP:
  ClusterMode: false   # 单机部署（默认）
  # ClusterMode: true # 集群部署，需确保 Redis 可用

REDIS:
  Url: "redis://:@127.0.0.1:6379/0"
  RedisPrefix: "x:"      # Redis 频道: "x:ws:broadcast"，在线计数: "x:ws:online:refs"
```

## 使用方式

### 初始化（自动完成）

初始化在 `core/ws.go` 的 `init()` 中自动完成，根据 `ClusterMode` 配置选择 PubSub 实现。
房间管理两种模式共用本地 `RoomManager`，无需区分：

```go
// core/ws.go
var Ws = ws.NewManager()

func init() {
    var ps ws.PubSub
    var rdb = Redis
    var prefix = config.RedisConfig.RedisPrefix
    if config.AppConfig.ClusterMode == true {
        channel := prefix + "ws:broadcast"
        ps = ws.NewRedisPubSub(Redis, channel)
    } else {
        ps = ws.NewLocalPubSub()
        rdb = nil
    }
    Ws.Init(ps, rdb, prefix)
    go Ws.Start()
}
```

### WebSocket 连接建立

在 Controller 中升级 HTTP 连接为 WebSocket，设置业务消息回调：

```go
func WsHandler(c *gin.Context) {
    uuid := util.ToolsUtil.MakeUuidV7()
    adminId := config.AdminConfig.GetAdminId(c)

    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }

    client := ws.NewClient(uuid, adminId, conn, core.Ws)
    client.OnMessage = handleWsMessage
    core.Ws.Register <- client

    go client.Write()
    go client.Read()
}
```

### 房间管理

房间通过业务消息动态加入/离开，不在连接时指定。客户端发送 JSON 控制消息：

```json
{"type": "join_room",  "data": "room1"}
{"type": "leave_room", "data": "room1"}
```

Controller 层路由到 Manager 操作：

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

房间数据仅存本地内存，集群下各实例独立维护自己的房间连接列表。
`SendToRoom` 和 `CloseRoom` 通过 PubSub 广播到所有实例，各实例各自处理本地连接。

### 消息推送

所有 `SendTo*` 方法自动将消息包装为 `WsResponse{type, data}` 格式：

```go
// 向指定用户推送通知
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

// 全局广播在线人数
core.Ws.SendToAll("onlineCount", map[string]any{
    "count": 5,
})
// 前端收到: {"type":"onlineCount","data":{"count":5}}
```

### 查询和关闭

```go
// 获取在线用户数（单机: 本地连接数，集群: 全局唯一用户数 via Redis HLEN）
count := core.Ws.GetOnlineCount()

// 关闭指定用户的所有连接（PubSub 广播，所有实例同时关闭）
core.Ws.CloseUser("user123")

// 关闭指定房间的所有连接（PubSub 广播，所有实例同时关闭）
core.Ws.CloseRoom("room456")

// 关闭所有连接（PubSub 广播，所有实例同时关闭）
core.Ws.CloseAll()
```

## 核心概念

### Client（客户端连接）

每个 WebSocket 连接对应一个 `Client` 实例：

| 字段 | 说明 | 用途 |
|------|------|------|
| `UUID` | 连接唯一标识（UUID v7） | 精确的单点推送 |
| `Uid` | 用户 ID | 按用户维度推送（一个用户可有多个连接） |
| `OnMessage` | 业务消息回调 | 心跳以外的消息上抛给 Controller 层处理 |

每个 Client 运行两个 goroutine：
- **Read**：读取客户端消息，ping→pong 内部处理，其他消息通过 `OnMessage` 回调上抛
- **Write**：向客户端发送消息，定期 Ping 保活

### Manager（连接管理器）

使用**单 goroutine 事件循环**处理连接的注册/注销，避免并发修改 map：

```
Register channel   → 添加连接到 clients/uid_uuids
                     集群模式: 首个本地连接 HINCRBY ws:online:refs +1
UnRegister channel → 移除连接，清理 uid_uuids + LeaveAll 清理房间
                     集群模式: 最后一个本地连接 HINCRBY -1（归零则 HDEL）
```

### PubSub（消息总线）

通过接口抽象实现单机/集群的无缝切换：

| 方法 | 说明 |
|------|------|
| `Publish(msg)` | 发布消息到总线 |
| `Subscribe()` | 订阅消息，返回 channel |
| `Close()` | 关闭总线，释放资源 |

**单机模式（LocalPubSub）**：
- 使用 Go channel 传递消息
- 无序列化开销，性能最优
- 仅支持单实例

**集群模式（RedisPubSub）**：
- 使用 Redis Pub/Sub 广播
- 消息经 JSON 序列化/反序列化
- 所有实例订阅同一频道，消息广播到所有实例
- 各实例收到消息后检查本地连接，有则推送，无则忽略
- 依赖 go-redis 内置重连机制

### RoomManager（房间管理）

**两种模式共用同一实现**，纯本地内存，不依赖 Redis：

| 方法 | 说明 |
|------|------|
| `Join(uuid, roomID)` | 将连接加入房间 |
| `Leave(uuid, roomID)` | 将连接离开房间 |
| `LeaveAll(uuid)` | 连接断开时离开所有房间 |
| `GetClientUUIDs(roomID)` | 获取房间内本地连接 UUID 列表（用于消息推送） |

集群下房间消息推送完全通过 PubSub 完成：
- `SendToRoom` → PubSub 广播 → 各实例 `localSendToRoom` 查本地 rooms 推送
- `CloseRoom` → PubSub 广播 → 各实例 `localCloseRoom` 关本地连接

### WsResponse（统一消息格式）

所有 `SendTo*` 方法自动包装为 `WsResponse{type, data}` 格式：

| 字段 | 类型 | 说明 |
|------|------|------|
| `type` | string | 业务消息类型，如 "onlineCount"、"notice"、"chat" |
| `data` | any | 消息内容 |

前端通过 `type` 字段路由到对应处理逻辑，从 `data` 中取具体数据。

## 心跳机制

```
客户端                              服务端
  │                                   │
  │──── "ping" (文本消息) ──────────►│  Read goroutine 通过 send channel 回复
  │◄─── "pong" (文本消息) ───────────│  Write goroutine 统一写出
  │                                   │
  │◄─── Ping (WebSocket帧) ──────────│  Write goroutine 每54秒发送
  │──── Pong (WebSocket帧) ──────────►│  60秒未收到则断开连接
  │                                   │
```

| 参数 | 值 | 说明 |
|------|-----|------|
| `writeWait` | 10s | 写操作超时 |
| `pongWait` | 60s | 等待 Pong 超时 |
| `pingPeriod` | 54s | Ping 发送周期（pongWait × 90%） |
| `maxMessageSize` | 2048B | 客户端消息大小限制 |
| `sendChanSize` | 256 | 发送缓冲区大小 |

## 并发安全

| 操作 | 保护机制 |
|------|----------|
| 连接注册/注销 | 单 goroutine 事件循环（无锁） |
| 房间操作 | `RWMutex` 读写锁 |
| 连接查询/推送 | `RWMutex` 读写锁，拷贝指针列表后释放锁再发送 |
| 连接关闭 | `sync.Once` 保证只执行一次 |
| 消息发送 | 非阻塞 select，channel 满时丢弃 |
| PubSub 关闭 | `closed` 标志 + `RWMutex` 保护 |

## 注意事项

1. **集群模式下 GetOnlineCount 返回全局唯一在线用户数**（Redis HASH HLEN），单机模式返回本地连接数
2. **房间数据仅存本地**，不提供跨实例 `GetRoomMembers` / `GetRoomOnlineCount`，如需全局房间统计可通过 PubSub 心跳聚合
3. **Redis Pub/Sub 是 fire-and-forget 模式**，不持久化消息，Redis 断线期间的消息会丢失
4. **发送缓冲区满时会丢弃消息**，日志中会记录 `[ws] client xxx send buffer full, message dropped`
5. **Redis 频道名使用项目前缀**（`{RedisPrefix}ws:broadcast`），在线计数 key 为 `{RedisPrefix}ws:online:refs`
6. **客户端发送的所有非心跳消息都通过 OnMessage 回调上抛**，基础设施层不解析业务消息
7. **CloseRoom/CloseUser/CloseAll 通过 PubSub 广播**，所有实例同时关闭对应连接，房间数据通过 unregister 流程自然清理
