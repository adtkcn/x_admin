# core/pubsub

mitt 风格的事件发布/订阅（`Emitter`），提供**统一 API + 可插拔的多种后端**。

- **统一 API**：无论单机还是集群，调用方只面对同一套接口 —— `On` / `Off` / `Emit`，
  按事件类型分发，支持通配符 `*`（语义参考 [mitt.js](https://github.com/developit/mitt)）。
- **可插拔后端**：底层传输可插拔，当前内置 `redis`（Redis Pub/Sub 跨实例广播）；
  NATS / Kafka 等只需实现 `Emitter` 接口即可接入。
- **纯传输**：本包不含任何业务语义，载荷统一为 `[]byte`，编码/解码由调用方负责。

> 为让各后端语义一致（跨进程必须可序列化），`Emit`/`Handler` 统一使用
> `[]byte` 作为载荷。上层（如 `core/ws`）自行定义业务消息结构，发送时序列化、接收时反序列化。

---

## 核心概念

### Handler — 事件处理函数

```go
type Handler func(payload []byte) // payload 为透明字节载荷，订阅方自行解码
```

### Emitter — mitt 风格事件总线接口

```go
type Emitter interface {
    // 订阅某类型事件（All "*" 监听全部），返回取消订阅函数
    On(eventType string, h Handler) (cancel func())
    // 移除该类型下的全部订阅
    Off(eventType string)
    // 发布事件（redis 经总线广播到所有实例，含本实例）
    Emit(eventType string, payload []byte) error
    // 关闭并释放资源，幂等
    Close() error
}
```

`All = "*"`：订阅它可收到任意类型事件。

---

## 后端

| 后端 | 适用场景 | 行为 |
|---|---|---|
| `redis` | 单实例 / 多实例集群 | 事件类型映射到 Redis 频道 `{prefix}{eventType}`；`Emit` 发布到 Redis，各实例（含自身）经 `PSubscribe("{prefix}*")` 订阅循环收到后本地分发，实现跨实例广播 |

各后端共用内部 `registry`（`registry.go`）负责 handler 存储与按类型 + 通配符分发。

---

## 配置与创建

统一通过 `Config` + `New` 工厂创建，是推荐的构造入口。

```go
type Backend string

const (
    BackendRedis Backend = "redis"
)

const DefaultPrefix = "pubsub:" // Prefix 为空时的中性默认值

type Config struct {
    Backend Backend // 留空默认 redis
    Prefix  string  // 事件频道前缀，建议带业务/项目前缀
}

func New(cfg Config, redisClient *redis.Client) (Emitter, error)
```

`redisClient` 在 `Backend == BackendRedis` 时需要，传 `nil` 会返回错误。

---

## 使用示例

### 1. 通用用法

```go
em, err := pubsub.New(pubsub.Config{
    Backend: pubsub.BackendRedis,
    Prefix:  "myapp:ev:",
}, redisClient)
if err != nil {
    log.Fatal(err)
}
defer em.Close()

// 订阅某类型，On 返回取消订阅函数
unsub := em.On("user.login", func(payload []byte) {
    var uid string
    json.Unmarshal(payload, &uid)
    // ...
})
defer unsub()

// 通配：监听所有事件
em.On(pubsub.All, func(payload []byte) { /* 任意事件 */ })

// 发布事件（payload 自行序列化）
data, _ := json.Marshal("uid-123")
em.Emit("user.login", data)

// 移除某类型全部订阅
em.Off("user.login")
```

> 切换后端只需改 `Backend`，业务代码零改动（新增后端调用方也无需改动）。

### 2. 业务层在其上构建（以 `core/ws` 为例）

`ws` 以业务消息类型为事件名，自行完成序列化/反序列化：

```go
// 订阅：按消息类型注册 handler
em.On(string(MsgTypeUser), func(payload []byte) {
    var msg WsMessage
    json.Unmarshal(payload, &msg)
    localSendToUser(msg.Target, []byte(msg.Data))
})

// 发布：业务消息 -> []byte -> Emit
payload, _ := json.Marshal(WsMessage{Type: MsgTypeUser, Target: uid, Data: data, NodeID: nodeID})
em.Emit(string(MsgTypeUser), payload)
```

---

## 扩展新后端（NATS / Kafka 等）

1. 在包内新建文件（如 `nats.go`），实现 `Emitter` 接口：
   - `On` / `Off`：复用内部 `registry`（`newRegistry()`）。
   - `Emit`：把 payload 发到对应主题。
   - 后台订阅循环：收到消息后还原 `eventType` 并调用 `registry.dispatch(eventType, payload)`。
   - `Close`：释放连接并 `reset` 注册表，保证幂等。
2. 在 `config.go` 增加 `BackendXxx` 常量。
3. 在 `New` 的 `switch` 中增加对应 `case`。

调用方（`core/ws.go` 等）无需任何改动。

---

## 设计原则

- **统一 API**：同一套 mitt 风格接口，切换后端零业务改动。
- **纯传输**：不定义消息类型、路由、节点标识等业务字段，只传输 `[]byte`。
- **易扩展**：新增后端只需实现接口并复用 `registry`。
- **线程安全**：`registry` 以 `RWMutex` 保护，分发前快照，避免 handler 内再订阅/取消导致的死锁。
- **幂等关闭**：`Close` 可安全多次调用。
