# 事件总线 pubsub

`core/pubsub` 提供 mitt 风格的事件发布/订阅（Emitter），统一 API、可插拔后端。当前内置 **Redis 后端**（基于 Redis Pub/Sub 实现跨实例广播），新增后端（NATS/Kafka 等）只需实现 `Emitter` 接口并在 `New` 中注册分支。

## 一、设计要点

- **统一 API**：调用方只面对 `On / Off / Emit`，不感知底层后端。
- **多后端**：底层传输可插拔，当前内置 `redis`。
- **payload 统一为 `[]byte`**：跨进程必须可序列化，本包只做透明传输，编码/解码（JSON、protobuf 等）由调用方负责。
- **通配符**：`All = "*"`，订阅它可收到任意类型事件。

## 二、核心接口

```go
// core/pubsub/interface.go
const All = "*"

type Handler func(payload []byte)

type Emitter interface {
    // On 订阅指定类型事件，eventType 为 All("*") 时监听全部事件。
    // 返回取消订阅函数，调用即移除该 handler。
    On(eventType string, h Handler) (cancel func())

    // Off 移除指定类型下的全部订阅。
    Off(eventType string)

    // Emit 发布事件。payload 为透明字节载荷。
    Emit(eventType string, payload []byte) error

    // Close 关闭发射器并释放底层资源（幂等）。
    Close() error
}
```

## 三、创建 Emitter

```go
// core/pubsub/config.go
func New(cfg Config, redisClient *redis.Client) (Emitter, error)
```

- 内部按 `cfg` 选择后端；当前仅 `redis` 分支：调用 `newRedisEmitter(redisClient, prefix)`，返回 `*redisEmitter`。
- `prefix`：Redis 频道前缀，用于 `PSubscribe(prefix + "*")` 一次性订阅所有事件类型，建议带业务/项目前缀以隔离。

示例（来自 `core/ws.go` 初始化）：

```go
emitter, _ := pubsub.New(pubsub.Config{Prefix: "ws:"}, core.Redis)
core.Ws.Init(emitter, redisClient, "ws:")
```

## 四、Redis 后端行为

`redisEmitter`（`core/pubsub/redis.go`）：

- **频道映射**：事件类型 `eventType` 映射到 Redis 频道 `{prefix}{eventType}`。
- **订阅**：`PSubscribe("{prefix}*")`，后台 `listen()` 循环接收消息，从频道名还原 `eventType` 后本地 `reg.dispatch`。
- **Emit 语义**：`Publish` 到 Redis，由各实例（含自身）的订阅循环收到后分发，从而实现**跨实例广播且各实例行为一致**。
- **消费容错**：`ReceiveMessage` 出错时，若非 `ctx` 取消则休眠 100ms 重试（依赖底层自动重连）。
- **幂等关闭**：`Close()` 用 `sync.Once` 取消 context、关闭 `psub`、reset 注册表。

## 五、使用示例

```go
// 订阅
cancel := emitter.On("order.paid", func(payload []byte) {
    // payload 由发布方编码，订阅方自行解码
    var evt OrderPaidEvent
    json.Unmarshal(payload, &evt)
    // ...
})
defer cancel()

// 发布（跨实例广播）
emitter.Emit("order.paid", mustJson(OrderPaidEvent{OrderId: 123}))

// 通配订阅（监听全部）
emitter.On(pubsub.All, func(payload []byte) { /* 调试/审计 */ })

// 取消某类型所有订阅
emitter.Off("order.paid")
```

## 六、与 core/queue 的区别

| 能力 | `core/pubsub` | `core/queue` |
| --- | --- | --- |
| 语义 | 广播（所有订阅者都收到） | 点对点（一条消息一个消费者） |
| 用途 | 实时事件通知（如 WebSocket 跨实例推送） | 异步作业、后台任务 |
| 载荷 | `[]byte`，调用方自解码 | `[]byte`，调用方自解码 |
| 后端 | Redis Pub/Sub | Redis List |

> queue 详见 [异步队列](./异步队列.md)；WebSocket 借助 pubsub 实现集群广播，详见 [WebSocket](./WebSocket.md)。
