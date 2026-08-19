# core/queue_delay —— 延迟队列

一个**与业务解耦**的延迟（定时）投递队列：消息投递后不会立即被消费，而是到达指定时间后才可被消费。与 `core/queue` 即时队列**互不耦合**，调用方可按需引入。

- 统一 `DelayedQueue` 接口，调用方不感知底层后端
- 内置 `redis` 后端（基于 Redis ZSET 的到期轮询）
- 消息体统一为 `[]byte`，队列不关心任务类型
- 并发 worker 数通过 `ConsumeDelayed` 的 `int` 参数控制
- 优雅退出：关闭时停止取新、等待在途（含进行中入队）消息处理完成，数据不丢

> 延迟消息采用「直接消费延迟池」模式：不经过就绪 List、不依赖即时队列的 `Consume`。
> 即时投递能力由独立包 `core/queue` 提供。

---

## 一、使用方式

### 1. 创建延迟队列

通过 `queue_delay.New` 工厂创建，是唯一推荐入口。项目全局实例见 `core/queue.go` 的 `core.QueueDelay`，业务层可直接复用。

```go
import "x_admin/core/queue_delay"

// 基于 Redis ZSET 的延迟队列（需传入 *redis.Client）
dq, _ := queue_delay.New(queue_delay.Config{
    Backend: queue_delay.BackendRedis,
    Prefix:  "myproj:", // 队列键前缀，用于隔离；不可为空
}, core.Redis)
```

`Backend` 留空时默认 `redis`。

### 2. 投递延迟消息

```go
body, _ := json.Marshal(map[string]any{"order_id": 123})

// 5 分钟后投递
dq.EnqueueDelay("order_jobs", body, 5*time.Minute)

// 或在指定绝对时间投递
dq.EnqueueAt("order_jobs", body, time.Now().Add(time.Hour))
```

### 3. 消费到期消息

`ConsumeDelayed` 直接消费延迟池中已到期的消息（不经过即时队列的 `Consume`）：

```go
// 5 个并发 worker（多 worker 竞争消费，同一条消息只被一个取到）
dq.ConsumeDelayed(ctx, "order_jobs", func(ctx context.Context, body []byte) error {
    var job map[string]any
    if err := json.Unmarshal(body, &job); err != nil {
        return err
    }
    return process(job)
}, 5)
```

- `ctx` 取消后，所有 worker 优雅退出（已在途消息仍处理完）。
- `Handler` 返回 `error` 仅被记录日志，当前**不自动重试**（见技术细节）。

---

## 二、技术细节

### 架构与接口

```go
// Handler 为本包独立定义，签名与 core/queue 的 Handler 保持一致（但不依赖 core/queue 包）
type Handler func(ctx context.Context, body []byte) error

type DelayedQueue interface {
    EnqueueAt(name string, payload any, at time.Time) error
    EnqueueDelay(name string, payload any, delay time.Duration) error
    ConsumeDelayed(ctx context.Context, name string, handler Handler, concurrency int) error
    Close() error
}
```

- `name` 是延迟队列名，对应 Redis 的 ZSET 键 `{Prefix}{name}`，member 为消息 JSON，score 为到期毫秒时间戳。
- 两包互不依赖：延迟队列的 `Handler` 为本包独立定义（签名与 `core/queue` 一致），不引用 `queue.Queue` 接口，也未 import `core/queue` 包。

### 后端实现（redis）

文件：`redis.go`

- **投递**：`ZADD {Prefix}{name} {到期毫秒} {body}`，score=到期时间戳。
- **取到期**：Lua 脚本 `takeDueScript` 原子执行「`ZRANGEBYSCORE` 取到期 + `ZREM` 删除」，多实例并发安全，避免重复消费。
- **消费**：`ConsumeDelayed` 周期性轮询到期消息，每轮最多取 `concurrency` 条（与 worker 数对齐，避免退出窗口囤积过多已删未处理消息），取出后投入容量为 `concurrency` 的有界 channel 交给 worker 池。取空后 `sleep`，空轮询指数退避 `100ms → 3s`。

### 数据安全与优雅退出

与 `core/queue` 一致：

- `ctx` 取消或 `Close()` **只停止「取新消息」**，已在途（已取出投入 channel）的消息一定会被 handler 处理完；worker 处理消息时不受 `ctx` 打断。
- 取出消息即 `wg.Add(1)`，worker 处理完 `wg.Done()`；handler panic 被 recover 兜底且仍 `wg.Done()`，不会卡死计数。
- `Close()`：先 `close(stopCh)` 停取新消息，再 `wg.Wait()` 阻塞等待所有在途（含进行中的 Enqueue）完成，**绝不丢失**。`stopCh` 由 `closeOnce` 保证只关一次。

### 并发与语义

- **点对点**：同一延迟队列的多 worker 竞争消费，单条消息仅被一个 worker 处理，避免重复执行。
- **并发控制**：`ConsumeDelayed` 的 `concurrency int` 参数启动 `n` 个 goroutine 且决定 channel 容量；`n<=0` 回落为 1。
- **背压**：有界 channel 满时阻塞轮询协程，避免内存无限堆积。
- **投递不阻塞**：redis 端 `ZADD` 为 O(1) 网络写。

### 扩展新后端

1. 实现 `DelayedQueue` 接口（`EnqueueAt` / `EnqueueDelay` / `ConsumeDelayed` / `Close`）。
2. 在 `config.go` 增加 `BackendXxx` 常量，并在 `New` 的 `switch` 中加一个 `case`。

### 已知限制与后续扩展

- **延迟消息同 payload 去重**：`EnqueueAt` 以消息 JSON 为 member，两条相同 JSON + 相同到期毫秒会相互覆盖；若业务需同 payload 多次投递，应在 member 中加唯一 ID。
- **精度**：到期轮询为秒级近似，非实时精确触发。
- **失败重试 / 死信队列（DLQ）**：当前 `Handler` 返回 error 仅记录日志即丢弃；可引入重试计数与死信队列。
- **多队列监听**：当前 `ConsumeDelayed` 一次监听一个 `name`，可扩展为批量订阅。
