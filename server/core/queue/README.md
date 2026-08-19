# core/queue —— 即时队列

一个**与业务解耦**的异步任务队列（即时队列）：消息投递后立即可被消费。统一接口、可插拔后端（当前内置 `redis`），消息体为 `[]byte`，业务层自行序列化。

- 统一 `Queue` 接口，调用方不感知底层后端
- 内置 `redis` 后端（基于 Redis List 的跨实例工作队列）
- 消息体统一为 `[]byte`，队列不关心任务类型
- 并发 worker 数通过 `Consume` 的 `int` 参数控制
- 优雅退出：关闭时停止取新、等待在途（含进行中入队）消息处理完成，数据不丢

> 广播 / 发布订阅语义请用 `core/pubsub`；本包是「点对点」工作队列（一条消息只被一个 worker 消费）。
> 延迟（定时）投递能力由独立包 `core/queue_delay` 提供，与本包互不耦合（其 `Handler` 为独立定义，签名保持一致但不依赖本包）。

---

## 一、使用方式

### 1. 创建队列

通过 `queue.New` 工厂创建，是唯一推荐入口。项目全局实例见 `core/queue.go` 的 `core.Queue`，业务层可直接复用。

```go
import "x_admin/core/queue"

// 集群 / 单机：基于 Redis List 的跨实例队列（需传入 *redis.Client）
q, _ := queue.New(queue.Config{
    Backend: queue.BackendRedis,
    Prefix:  "myproj:", // 队列键前缀，用于隔离；不可为空
}, core.Redis)
```

`Backend` 留空时默认 `redis`，因此 `queue.New(queue.Config{Prefix: "myproj:"}, core.Redis)` 等价。

### 2. 投递消息

`Enqueue` 向指定队列名投递一条 `[]byte` 消息：

```go
body, _ := json.Marshal(map[string]any{"order_id": 123})
err := q.Enqueue(ctx, "order_jobs", body)
```

### 3. 消费消息

`Consume` 启动消费者，从队列持续拉取并交给 `Handler` 处理：

```go
// 默认 1 个 worker
q.Consume(ctx, "order_jobs", func(ctx context.Context, body []byte) error {
    var job map[string]any
    if err := json.Unmarshal(body, &job); err != nil {
        return err
    }
    return process(job)
})

// 5 个并发 worker（多 worker 竞争消费，同一条消息只被一个取到）
q.Consume(ctx, "order_jobs", handler, 5)
```

- `ctx` 取消后，所有 worker 优雅退出（已在途消息仍处理完）。
- `Handler` 返回 `error` 仅被记录日志，当前**不自动重试**（见技术细节）。

---

## 二、技术细节

### 架构与接口

```go
type Handler func(ctx context.Context, body []byte) error

type Queue interface {
    Enqueue(name string, payload any) error
    Consume(ctx context.Context, name string, handler Handler, concurrency int) error
    Close() error
}
```

- `name` 是队列名，对应 Redis 的 List 键 `{Prefix}{name}`。
- 所有后端实现同一接口，调用方面向接口编程，后端可热替换。

### 后端实现（redis）

文件：`redis.go`

- **投递**：`RPUSH {Prefix}{name} body`。投递动作本身计入 `wg`（推前 Add，写成功后 Done），保证入队调用正常完成。
- **消费**：`BRPOP`（阻塞）单连接循环拉取，取到投入容量为 `concurrency` 的有界 channel，再由 `concurrency` 个 worker 并发处理；多个实例/worker 竞争同一 key，Redis 保证一条消息只被一个取到，天然分布式工作队列。仅占用 1 条 Redis 连接（取与处理解耦）。
- **网络抖动**：`BRPOP` 出错时 `sleep 100ms` 退避重试。

### 数据安全与优雅退出

- `ctx` 取消或 `Close()` **只停止「取新消息」**，已在途（已取出投入 channel）的消息一定会被 handler 处理完；worker 处理消息时不受 `ctx` 打断。
- 取出消息即 `wg.Add(1)`，worker 处理完 `wg.Done()`；handler panic 被 recover 兜底且仍 `wg.Done()`，不会卡死计数。
- `Close()`：先 `close(stopCh)` 停取新消息，再 `wg.Wait()` 阻塞等待所有在途（含进行中的 Enqueue）完成，**绝不丢失**。`stopCh` 由 `closeOnce` 保证只关一次。

### 并发与语义

- **点对点**：同一队列的多 worker 竞争消费，单条消息仅被一个 worker 处理，避免重复执行。
- **并发控制**：`Consume` 的 `concurrency int` 参数启动 `n` 个 goroutine 且决定 channel 容量；`n<=0` 回落为 1。
- **背压**：有界 channel 满时阻塞取消息协程，避免内存无限堆积。
- **投递不阻塞**：redis 端 `RPUSH` 为 O(1) 网络写。

### 扩展新后端

1. 实现 `Queue` 接口（`Enqueue` / `Consume` / `Close`）。
2. 在 `config.go` 增加 `BackendXxx` 常量，并在 `New` 的 `switch` 中加一个 `case`。

`redis` 的调用方与 `app/task` 适配层**完全不需要改动**。

### 已知限制与后续扩展

当前实现刻意保持精简，以下能力可作为后续扩展：

- **失败重试 / 死信队列（DLQ）**：当前 `Handler` 返回 error 仅记录日志即丢弃；可引入重试计数与 `{Prefix}dlq:{name}` 死信队列。
- **ACK 机制**：`BRPop` 取出即出队，处理中崩溃会丢消息；可改用 `RPOPLPUSH` + 定时回收未 ACK 消息（注：当前依靠 `Close` 等待在途规避正常退出丢消息，不防进程崩溃）。
- **多队列监听**：当前 `Consume` 一次监听一个 `name`，可扩展为批量订阅。
