---
name: "后端-IoT 物联网模块"
description: "x_admin IoT 设备接入、可视化规则链处理流水线、告警状态机、在线判定、时序存储与大屏实时推送"
---

# IoT 物联网模块

本模块在 x_admin 底座上落地了一套轻量物联网平台：**设备接入 → 消息队列削峰 → 可视化规则链处理 → 时序存储 / 告警状态机 / 运行态快照 → WebSocket 大屏实时推送**。它最大程度复用现有基建（Gin、GORM、Redis、`core.Queue`、`core.Ws`、`core/pubsub`、权限体系），仅新增 IoT 业务域，不重复造轮子。

## 一、复用与新增

| 能力 | 复用/新增 | 落位 |
| --- | --- | --- |
| Web/API、鉴权权限 | ✅ 复用 Gin + `LoginAuth`/`PermAuth` | `routes`、`app/middleware` |
| 元数据 ORM | ✅ 复用 GORM + MySQL（UUIDv7 主键、软删） | `app/model/iot_model` |
| 运行态缓存 | ✅ 复用 `core.Redis`（go-redis v9） | `pipeline/online` |
| 消息队列削峰 | ✅ 复用 `core.Queue`（Redis 点对点） | `app/queue` |
| 实时推送 / 集群广播 | ✅ 复用 `core.Ws` 房间 + `core/pubsub` | `pipeline/online` |
| JS 脚本引擎 | ✅ 复用 `dop251/goja` | `pipeline/rulechain` |
| 站内/邮件/Webhook 通知 | ✅ 复用 `notice_service` + 队列 | 规则链通知节点 |
| 设备协议接入（HTTP） | 🆕 新增接入网关 + 设备凭证鉴权 | `iot_ctl/http_gateway.go`、`middleware/device_auth.go` |
| 可视化规则链引擎 | 🆕 新增树形执行器 + 节点库 | `pipeline/rulechain` |
| 时序存储 | 🆕 MySQL 暂代（预留切 TDengine） | `pipeline/store` |
| 设备地图大屏（前端） | 🆕 MapLibre | `admin/src/views/iot/monitor` |

## 二、总体架构与数据流

```
设备 ──(X-Device-Sn / X-Access-Token)──► 接入层 iot_ctl/http_gateway
   POST /api/iot/telemetry | /heartbeat        │ 只校验凭证 + 组标准信封 IngestMsg + 条数上限
                                               ▼
                                 core.Queue.Enqueue("iot:ingest")   ← 削峰解耦，请求线程不做重活
                                               │
                                               ▼  Consume(concurrency=1)
                        处理流水线 pipeline.HandleIngest
                          ├─ online.RefreshOnline(msg)      所有报文都“续命”（心跳/上报）
                          └─ rulechain.Run(msg, 产品规则链)  逐节点沿执行树下钻
                                ├─ 解码/过滤/布尔 (goja run)
                                ├─ 状态快照节点 → online.SaveLatestValues (Redis HASH + 大屏脏集合)
                                ├─ 历史存储节点 → store.Write (批写 MySQL 时序表)
                                ├─ 告警节点     → alert.EvalAlert (Redis 状态机 → 历史表 + WS)
                                └─ 通知节点     → 站内 / 邮件 / webhook（后两者走队列）
                                               │
              ┌────────────────────────────────┼────────────────────────────────┐
              ▼                                ▼                                ▼
        MySQL(元数据+时序+告警历史)      Redis(在线/快照/告警活动态)        core.Ws 房间广播
                                              │                                │
                                        app/corn 定时任务                 大屏/设备详情前端
                        (离线扫描/慢路径重建/数据清理/大屏冲刷)
```

三条**独立事件流**汇入前端：`iot:status`（上下线）、`iot:telemetry`（最新属性）、`iot:alarm`（告警触发/升级/恢复）。

## 三、目录落位

```
server/
├─ app/model/iot_model/            # x_iot_* GORM 实体
├─ app/schema/iot_schema/          # 各接口 Req/Resp + 规则链节点契约 + 接入信封
├─ app/service/iot_service/        # 元数据 CRUD 业务（product/device/space/category/rule_chain/alert…）
│  └─ pipeline/                    # 报文处理流水线（接入解耦后的核心）
│     ├─ pipeline.go               # 消费入口 + 对外门面
│     ├─ online/                   # 运行态：在线/离线判定、属性快照、状态/遥测/大屏广播
│     ├─ rulechain/                # 可视化规则链执行引擎 + goja 脚本原语
│     ├─ alert/                    # 告警状态机（Redis 活动态 → 历史表）
│     └─ store/                    # 时序批写/回读/清理（流水线唯一数据落地出口）
├─ app/controller/iot_ctl/         # 设备侧 HTTP 接入网关
├─ app/controller/admin_ctl/iot_controller/  # 管理端 Handler
├─ app/middleware/device_auth.go   # 设备凭证鉴权
├─ app/queue/iot.go                # iot:ingest 消费回调
├─ routes/iot_route/               # 设备侧接入路由（/api/iot/*）
├─ routes/admin_route/iot_route.go # 管理端 CRUD 路由
└─ config/iot.go                   # IoT 配置
```

前端：`admin/src/views/iot/*`（页面）+ `admin/src/api/iot/*`（接口封装）+ 菜单权限 SQL（`docs/sql/x_iot_*.sql`）。

## 四、数据模型（`x_iot_` 前缀）

沿用全库约定：UUIDv7 `char(36)` 主键（`BeforeCreate` 生成）、`x_null.Time` 时间、`soft_delete.DeletedAt` 标志位软删。

| 表 | 说明 | 关键字段 |
| --- | --- | --- |
| `x_iot_product_category` | 产品分类树 | `name`、`parent_id`、`sort` |
| `x_iot_product` | 产品定义 | `category_id`、`device_type`(DEVICE/GATEWAY/EDGE)、`protocol`、`offline_timeout`、`thing_model`(物模型 JSON)、`rule_chain_id`(引用的可视化规则链)、`is_stop` |
| `x_iot_device_space` | 设备区域树 | `parent_id`、`path`(物化路径 `/uuid1/uuid2/`)、`longitude`、`latitude`、`sort` |
| `x_iot_device` | 设备 | `sn`(业务唯一)、`product_id`、`space_id`、`gateway_id`、`status`(0/1，接入层自动维护不开放编辑)、`offline_timeout`(冗余产品值)、`access_token`(接入凭证)、经纬度 |
| `x_iot_rule_chain` | 可视化规则链 | `graph_data`(LogicFlow 图，仅供绘图回填)、`tree_data`(执行树，后端按此下钻)、`is_stop` |
| `x_iot_device_data` | 设备时序数据（MySQL 暂代 TDengine） | `(device_id, ts)` 联合索引、`values`(属性 JSON 文本)、`protocol` |
| `x_iot_alert_history` | 告警事件历史 | `status`(1 告警中/0 已恢复)、`severity`(-1 故障/1 轻微/2 严重，数值即危害序)、`values`(快照)、`triggered_at`、`recovered_at` |



## 五、设备接入层

- **协议适配**：当前落地 HTTP（`iot_ctl/http_gateway.go`），MQTT/TCP/WS 适配器可复用同一信封与队列平行接入。
- **设备鉴权**：`middleware.DeviceAuth()` 校验 `X-Device-Sn` + `X-Access-Token`（凭证见 `x_iot_device.access_token`），与后台登录态无关；鉴权结果 `DeviceAuthInfo` 写 `gin.Context` 并缓存于 Redis（`DeviceCacheTTL`），高频上报不逐条查库。`OfflineTimeout` 已按 **设备 > 产品 > 兜底 300s** 回退后随信封带下。
- **信封 `IngestMsg`**：接入层唯一的协议无关契约（`MsgType`、`DeviceId`、`ProductId`、`SpaceId`、`OfflineTimeout`、`Values`、`Ts`、`ReceivedAt`…）。设备时间戳缺失/异常时用服务端接收时间兜底。
- **只投递不处理**：校验条数上限（`MaxValues` 防报文放大）后 `core.Queue.Enqueue("iot:ingest", msg)` 立即返回，接入层耗时恒定。

**设备侧接口**（`/api/iot/*`，走 `DeviceAuth`）：

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| POST | `/api/iot/telemetry` | 属性/遥测上报 `{"values":{...},"ts":...}` |
| POST | `/api/iot/heartbeat` | 心跳（可空 body，仅刷新在线态） |

## 六、处理流水线

`app/queue/iot.go` 的 `IotStart` 消费 `iot:ingest`（`core.Queue.Consume(..., concurrency=1)`），转 `pipeline.HandleIngest`：

1. 解析信封，坏报文记日志后丢弃（不重试，避免死循环）；
2. `online.RefreshOnline(msg)` —— 任何报文都“续命”，并在 **离线→在线** 转换时写库 + 广播；
3. 按类型分派：`telemetry` 交给所属产品的可视化规则链 `rulechain.Run`；`heartbeat` 到第 2 步即止。

子包依赖构成单向 DAG：`online ← alert ← rulechain ← pipeline`，`store` 为叶子旁路。

## 七、可视化规则链引擎

产品挂一条规则链（`rule_chain_id`），逐报文沿**执行树**下钻。链在语义上是一棵树：触发器为唯一根，除并行节点外每节点至多一条后继；前端从 LogicFlow 图以「输入」为根展开后直传 `tree_data`，后端直接反序列化下钻（不做图→树展开）。

- **两级读缓存 + pubsub 即时失效**：产品模型与规则链执行树都遵循「本进程内存（`LocalCacheTTL`）→ Redis（`ThingCacheTTL`）→ DB」两级回源，热路径零 Redis 往返；规则链本内存层缓存的是 `CachedParsedTree`（已解析的 `*iot_schema.RuleTreeNode`），一并抹平 JSON 重解析开销。管理端 CRUD 后 `InvalidateProductCache` / `InvalidateChainCache` 除了删 Redis，还经 `core.IotBus`（频道前缀 `{RedisPrefix}iot:bus:`）广播失效事件，各实例订阅回调直接 `core.CacheDelete / CacheDeletePrefix` 清本内存，下一报文重新回源。
- **遍历语义**：节点 handler 只回答「本分支是否继续」。过滤 `false`/解码产出空 → 置 `halt` 短路整条报文；布尔 `false` 仅剪本分支；并行节点为每条子分支深拷贝独立属性负载，分支互不干扰（`halt` 位经指针共享对所有分支即时可见）。

**节点库**（注册表 `rcHandlers`，未知类型按「继续下钻」兜底）：

| 节点 | type | 副作用 | 返回契约 |
| --- | --- | --- | --- |
| 触发器 | `rule:trigger` | 按报文类型入口过滤（`msg_types` 空=全放行） | 命中放行，否则 `halt` |
| 解码 | `rule:decode` | `goja run(values,sn)` 覆盖属性 | 返对象；空则 `halt` 丢弃 |
| 数据过滤 | `rule:filter` | `goja run(values,sn)` 返布尔 | `false` → `halt` 丢整条 |
| 布尔 | `rule:bool` | `goja run(values,sn)` 返布尔 | `false` 仅剪本分支 |
| 并行 | `rule:parallel` | 无，多出口分流由树 children 表达 | 恒继续 |
| 告警 | `rule:alarm` | 级别交告警状态机收敛 | 恒继续 |
| 状态快照 | `rule:state` | 写 Redis 最新属性 + 标大屏脏集合 | 恒继续 |
| 历史存储 | `rule:history` | 经 `store.Write` 落时序表 | 恒继续 |
| 通知 | `rule:notify` | 站内/邮件/webhook（后两者走队列） | 恒继续 |

- **脚本契约**：各脚本节点统一实现全局函数 `run(values, sn)`，后端固定调 `run`。goja **编译一次缓存**（键 `链ID:节点ID:stage`），**每报文新建 Runtime**（Runtime 非并发安全），`panic`/异常 `recover` 兜住不拖垮 worker。脚本为空即透传/默认放行。
- **物模型白名单**：快照、时序、告警三处落库前统一经 `ScopeValues(values, allow)` 裁剪，只保留产品物模型已建模的 `identifier`（`allow` 为空表示未配物模型，此时不约束，避免误丢数据）。

## 八、告警状态机（report 驱动，双事件流）

告警节点脚本逐条返回级别数字（`0`=恢复/不告警，`-1`=故障，`1/2/…`=阈值级别），`alert.EvalAlert` 把它收敛成**事件级**记录：报文驱动、只在转换点写库推 WS，无需到期扫描。

- **两条独立流**：故障流(-1) 与阈值流(1/2/…) 各自维护一条「告警中」记录，互不合并。每条报文只确认其一，另一流即收恢复信号。
- **Redis 运行态**（`{前缀}iot:alert:*`，普通流无前缀、故障流 `fault:` 前缀，两组各三键）：`active`(=`历史ID:峰值:当前`，无 TTL)、`first`(持续计时起点，仅 `duration>0`)、`recover`(恢复确认起点，仅 `recover>0`)。六键一次 `MGet` 读出。
- **可配时长**（秒，节点属性）：`duration` 命中后需连续异常满时长才触发；`recover` 恢复正常后需连续满时长才收尾，确认期间切回则计时作废。
- **记录留峰值**：同流内 `severity` 只升不降（危害升级才 `UPDATE`）。
- **聚合级推送**：大屏一点只显示一条聚合级（故障优先，否则阈值当前级），仅聚合级变化才推 `iot:alarm`，避免并存期误清/降级。
- **收尾**：设备转离线经 `online` 钩子收尾两流活动告警；删设备清六键；进程重启前 `RecoverRuntime` 按库「告警中」记录回灌活动态，防 Redis 丢数据导致永久卡告警。

## 九、在线 / 离线判定（到期表 ZSET 驱动）

在线方向由报文实时驱动，离线方向由到期表驱动，运行态全在 Redis：

- **到期表** `iot:expire:zset`：`member=deviceId`，`score=离线到期毫秒`。是「谁在线 + 何时到期」的唯一事实源。每报文 `keepAlive` 刷新 `score=now+窗口`（窗口 = 设备 `offline_timeout`，兜底 300s），并记 `iot:seen:<id>`。
- **上线转换**：刷新前 `score` 缺失/已到期即判「离线→在线」，`MarkStatus(1)` 写库 + 广播；写库失败撤销刚写的 score，让下一条报文重试。
- **快路径** `ScanOffline`（corn `iotOfflineScan`，每 5 秒）：Lua 原子弹出已到期成员置离线并广播。
- **慢路径** `SweepOffline`（corn `iotOfflineSweep`，每 5 分钟）：全扫库侧在线设备按 `seen+窗口` 现算，超窗置离线、未超窗但 score 缺失/过期的回填（Redis 重启后重建的安全网）。
- **回读** `ReadDeviceState`：管理端实时状态页数据源，直读到期表 score，缺失时退回 `seen+窗口` 推算。

## 十、时序存储（批写缓冲 + 保留清理）

`store.Write` 是流水线唯一数据落地出口：属性先按物模型白名单裁剪，再投内存缓冲 `bufCh`；后台 `flusher` 协程按「攒满 `DataBatchSize` 行或每 2 秒」批量 `INSERT`。缓冲满（库堵）时退回同步单写，宁压慢 worker 也不丢数据；优雅关闭 `Flush` 冲刷剩余。`QueryRange` 供趋势图按时间区间回读；`Clean`（corn `iotDataClean`，每日 03:00）按 `DataKeepDays` 清理保留期外数据（0=永久）。切换 TDengine 只需替换 `writeSink` 实现，主链路不变。

## 十一、实时推送与大屏

复用 `core.Ws` 房间，天然支持集群（经 `core/pubsub` 跨实例扇出）：

- **房间**：单设备 `iot:device:{deviceId}`、区域大屏 `iot:map:{spaceId}`；前端连 `/api/ws` 后按需加入。
- **事件**：`iot:status`（上下线，仅转换点推）、`iot:telemetry`（最新属性快照）、`iot:alarm`（告警触发/升级/恢复，带 `alert_id` 供前端合并展示）。
- **大屏冲刷** `PushScreenTick`（corn `iotScreenPush`，每 2 秒）：无在线观众早退；否则一次取完「大屏脏集合」`iot:push:dirty`（快照更新即标脏），按 500 一批读回快照、逐台按区域房间广播 `iot:telemetry`。不限台数而受单轮 1.5 秒时间预算约束：超出预算的设备不推也不移出，顺延下个 tick（固定条数上限会把稳态吞吐钉死，脏集合清不完反而越积越多）。
- **前端大屏**：MapLibre 渲染设备点位（受 `Popup` + Vue 组件驱动），按所属产品物模型展示属性名/单位，收 WS 事件增量刷新。

## 十二、管理端 API（`/api/iot/*`，需登录 + 权限）

- 表单下拉（仅 `LoginAuth`）：`productCategory/all`、`deviceSpace/all`、`product/all`、`ruleChain/all`。
- CRUD（`PermAuth`，写操作记审计日志）：产品分类、产品、设备区域、设备的 增删改查 + 排序；
- 运行态与链路验证：`device/state`（Redis 实时状态回读）、`device/data`（时序历史回读，供趋势图）、`device/map`（区域全部点位快照，实时增量走 WS）；
- 告警历史（处理层自动写入，管理端只读/清理）：`alert/list`、`alert/del`；
- 可视化规则链：`ruleChain/list|detail|add|edit|del`。

## 十三、前端页面

`admin/src/views/iot/`：`product_category`、`product`（含**物模型编辑器** `thing-model.vue`）、`device_space`、`device`（详情三 Tab：实时状态 / 趋势图 / 告警）、`rule_chain`（LogicFlow **规则链设计器**）、`alert`（告警历史）、`monitor`（MapLibre **设备大屏**）。接口封装在 `admin/src/api/iot/`（按实体拆分）。

## 十四、配置项（`config/iot.go`，`Iot.*`）

| 键 | 默认 | 说明 |
| --- | --- | --- |
| `Enable` | true | 接入层总开关：`false` 时不注册设备侧路由、不启动流水线 |
| `DeviceCacheTTL` | 300 | 设备鉴权信息 Redis 缓存秒数 |
| `ThingCacheTTL` | 600 | 物模型 / 规则链执行树 Redis 层缓存秒数 |
| `LocalCacheTTL` | 60 | 产品/规则链在本进程内存层 TTL（秒），pubsub 广播未达时的兜底 |
| `StateTTL` | 86400 | 最新属性快照保留秒数 |
| `MaxValues` | 50 | 单报文最大属性条数（防报文放大） |
| `PushStateChange` | true | 状态/遥测/告警是否经 `core.Ws` 广播 |
| `LogRaw` | false | 记录原始报文日志（联调用） |
| `LogRuleChain` | false | 逐节点记录规则链执行日志（联调用） |
| `DataBatchSize` | 200 | 时序批量写入单批行上限 |
| `DataKeepDays` | 30 | 时序保留天数，0=永久不清理 |

> `Normalize()` 在 viper 载入后把非正数回退默认值——这些字段直接当 Redis TTL 用，写 0 会语义反转（`MaxValues=0` 全拒、`StateTTL=0` 即失效）。`DataKeepDays`/设备 `offline_timeout` 的 0 是合法值，不受此约束。

## 十五、定时任务（`app/corn/fixed_task.go`）

| 任务 | 表达式 | 作用 |
| --- | --- | --- |
| `iotOfflineScan` | `*/5 * * * * *` | 快路径：弹到期表已到期设备置离线 |
| `iotOfflineSweep` | `0 */5 * * * *` | 慢路径：全扫复核 + 重建到期表（安全网） |
| `iotDataClean` | `0 0 3 * * *` | 清理保留期外的时序数据 |
| `iotScreenPush` | `*/2 * * * * *` | 冲刷大屏脏集合，按区域房间广播最新值 |

集群下均由 cron 分布式锁保证仅一实例执行。

## 十六、演进方向

- **多协议接入**：MQTT / TCP（粘包：固定帧头 + 长度字段）/ WS 适配器，产出同一 `IngestMsg`。
- **时序库切换**：`x_iot_device_data` → TDengine（超级表 + 自动建子表），仅改 `store` 落地实现。
- **规则链扩节点**：新增节点 = 注册 handler + schema 加属性 struct + 前端加节点/属性面板，不改遍历主逻辑。
