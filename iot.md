
# IoT 物联网平台 —— 基于 x_admin 的项目化落地规划

> 本文在原 IoT 架构设想基础上，**对齐现有 `x_admin`（Go 1.27 + Gin 1.12 + GORM + MySQL）工程**，
> 明确「哪些能力直接复用现有基础设施、哪些需要新增」，作为各阶段开发的落地说明书。

## 〇、与现有工程的关系（先看这一节）

x_admin 已具备 IoT 平台所需的多数底座，本方案**优先复用，不重复造轮子**：

| IoT 所需能力 | 现有工程是否已具备 | 落位 |
|:---|:---|:---|
| Web/API 框架 | ✅ Gin，统一 `response.JSON` 封装 | `core/response`、`routes` |
| 元数据 ORM | ✅ GORM + MySQL，表前缀 `x_`、UUIDv7 主键、`x_null.Time`、软删 | `core/db.go`、`app/model` |
| Redis 客户端 | ✅ `core.Redis`（go-redis v9），带 key 前缀 | `core/redis.go`、`config/redis.go` |
| WebSocket 推送 | ✅ `core.Ws` 管理器，支持 **房间/用户/全员** 广播，集群下经 `core/pubsub` 跨实例 | `core/ws`、`core/pubsub` |
| 异步队列削峰 | ✅ `core.Queue`（Redis，点对点）+ `core/queue_delay`（延迟投递） | `core/queue`、`core/queue_delay` |
| 集群模式开关 | ✅ `config.AppConfig.ClusterMode` | `config/app.go` |
| 鉴权/权限 | ✅ `middleware.LoginAuth()` + `middleware.PermAuth()` | `app/middleware` |
| 路由自动加载 | ✅ 各模块 `init()` 追加 `routeHandlers`，免手工挂载 | `routes/admin_route` |
| CRUD 代码生成 | ✅ 内置代码生成器，可按表一键生成前后端 | `app/service/generator_service` |
| **MQTT/TCP 协议接入** | ❌ 需新增 | 新建 `iot` 接入层 |
| **goja JS 引擎** | ✅ 已引入 dop251/goja | 产品级解码/过滤/告警脚本 |
| **TDengine 时序库** | ❌ 需新增依赖与连接 | 新建 `iot` 时序层 |
| **MapLibre 设备地图（前端）** | ❌ 前端新增页面 | `admin/src/views/iot` |

**结论**：IoT 相关代码统一落在新的 `iot` 业务域（对齐 `fabu`/`system` 的目录惯例），API 层、存储层、推送、告警运行态尽量直接调用现有 `core.*`。

---

## 一、总体架构（对齐 x_admin 分层）

```
┌─────────────────────────────────────────────────────────────┐
│                     API 层 (Gin / admin_route)               │
│  产品分类 / 产品 / 设备区域 / 设备 / 告警规则 / 设备地图        │
│  —— 复用 response.JSON、LoginAuth+PermAuth、routeHandlers     │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                  接入层 iot/gateway（本方案新增）              │
│  MQTT Gateway │ TCP Gateway │ HTTP Gateway │ WebSocket GW     │
└─────────────────────────────────────────────────────────────┘
                              │  投递到 core.Queue 削峰
┌─────────────────────────────────────────────────────────────┐
│              数据处理流水线 iot/pipeline（新增）              │
│  解码脚本 → 过滤脚本 → 存储 → 告警脚本+告警状态机 (goja)  │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                       存储与推送层                            │
│  MySQL(元数据/core.GetDB) │ Redis(运行态/core.Redis)          │
│  TDengine(时序,新增) │ WebSocket 大屏推送(core.Ws 复用)        │
└─────────────────────────────────────────────────────────────┘
```

---

## 二、技术栈选型

| 模块 | 方案 | 复用/新增 |
|:---|:---|:---|
| Web 框架 | Gin（现有） | 复用 |
| ORM | GORM + MySQL（现有 `core.GetDB`） | 复用 |
| 运行态缓存 | Redis（现有 `core.Redis`） | 复用 |
| 消息队列削峰 | `core.Queue` / `core/queue_delay`（现有） | 复用 |
| 实时推送 | `core.Ws` 房间广播（现有） | 复用 |
| 协议接入 | `paho.mqtt.golang`、`net.Listen`、`gorilla/websocket`（已在 go.mod） | 新增 MQTT，其余部分已有 |
| JS 引擎 | `github.com/dop2514/goja` | 新增依赖 |
| 时序存储 | TDengine `github.com/taosdata/driver-go/v3` | 新增依赖 |
| 地图前端 | MapLibre GL JS | 前端新增 |

> 新增依赖在 `server/go.mod` 引入；TDengine 建议通过独立 `core/td`（参照 `core/db.go` 写法）管理连接池，或第一阶段先不引入，见路线图。

---

## 三、目录与分层落位（对齐现有工程）

新增 `iot` 业务域，完全沿用 `fabu` 模块的分层与命名惯例：

```
server/
├─ app/
│  ├─ model/iot_model/           # GORM 实体：product_category/product/device_space/device/alert_*
│  ├─ schema/iot_schema/         # 各接口 Req/Resp 结构体（含 form/binding 标签）
│  ├─ service/iot_service/       # 业务逻辑，包级单例 var XxxService = New()，用 core.GetDB()
│  ├─ controller/admin_ctl/iot_controller/  # Handler 结构体，调用 service
│  ├─ iot/gateway/               # 【新增子域】MQTT/TCP/HTTP/WS 协议接入
│  ├─ iot/pipeline/              # 【新增子域】goja 预处理、过滤、告警状态机
│  └─ iot/td/ 或 core/td         # 【新增】TDengine 连接与超级表写入
├─ routes/admin_route/iot_route.go # init() 追加 routeHandlers，挂 LoginAuth+PermAuth
└─ config/iot.go                 # 【新增】IoT 相关配置（MQTT 端口、TDengine DSN 等）
```

前端：`admin/src/views/iot/*`（页面）+ `admin/src/api/iot/*`（接口封装）+ 菜单权限 SQL。

---

## 四、MySQL 元数据设计（对齐 GORM 模型惯例）

沿用现有模型写法：**UUIDv7 `char(36)` 主键**（`BeforeCreate` 生成）、`x_null.Time` 的 `autoCreateTime/autoUpdateTime`、`soft_delete.DeletedAt` 标志位软删、表前缀 `x_`（由 `NamingStrategy` 自动加，模型内显式 `TableName()` 返回 `x_iot_*`）。
> 注意：原方案里 `product`/`device` 等使用 `INT` 主键；为与 x_admin 全库主键风格统一，本方案调整为 **UUID 字符串主键**，关系字段（`category_id`/`product_id`/`space_id`/`gateway_id`）同样为 `char(36)`。TDengine tag 里的 `device_id` 用字符串设备标识或 sn，不必是数值。

| 表（`x_iot_` 前缀） | 关键字段 | 说明 |
|:---|:---|:---|
| `x_iot_product_category` | id(uuid), name, parent_id, sort, is_delete, create/update_time | 产品分类树 |
| `x_iot_product` | id(uuid), name, category_id, device_type, protocol, thing_model(JSON/text), decode_script, filter_script, alert_script, alert_duration, alert_silence, is_delete | 产品定义；`thing_model` 存物模型 JSON，`decode_script/filter_script/alert_script` 为产品级 goja JS 脚本；`alert_duration`(持续秒,0=命中即告)/`alert_silence`(静默秒,0=不静默) 为告警状态机参数（本项目 MySQL 无 jsonb，用 text 列） |
| `x_iot_device_space` | id(uuid), name, parent_id, path, longitude, latitude, sort, is_delete | 设备区域树，`path` 物化路径（如 `/uuid1/uuid2/`）便于子树查询 |
| `x_iot_device` | id(uuid), name, sn, product_id, space_id, gateway_id, status, longitude, latitude, access_token, is_delete | 设备；`sn` 业务唯一编号，`access_token` 接入凭证 |
| ~~`x_iot_alert_rule`~~ | — | 未建：告警条件与级别由产品级 goja 告警脚本 `alert(values)` 逐条判定（只返数字：0=不告警，非 0 即级别，数值按危害升序：-1=故障、1=轻微、2=严重，3+ 预留；故障危害最轻但优先于其他告警），持续/静默参数落在 `x_iot_product.alert_duration/alert_silence` |
| `x_iot_alert_history` | id(uuid), device_id, sn, product_id, space_id, status(1告警中/0已恢复), severity(tinyint 有符号：-1故障/1轻微/2严重，数值即危害序，3+ 预留), values(JSON快照), triggered_at, recovered_at | 告警事件历史：一条=一次触发到恢复的完整生命周期，单事件可改写（优先序更高时只更新同条 severity：危害升级或故障接管；故障消失后把记录交还给实际命中的级别）；持续时长=`recovered_at-triggered_at`（量大时按月分表/分区） |

模型示例（对齐 `fabu_model.FabuApp`）：

```go
package iot_model

import (
    "uuid"

    "github.com/adtkcn/x_null"
    "gorm.io/gorm"
    "gorm.io/plugin/soft_delete"
)

// IotProduct 产品定义
type IotProduct struct {
    ID               string                `gorm:"primarykey;type:char(36);comment:'主键'" json:"id"`
    Name             string                `gorm:"type:varchar(255);not null;default:'';comment:'产品名称'" json:"name"`
    CategoryId       string                `gorm:"type:char(36);not null;default:'';index;comment:'分类ID'" json:"category_id"`
    DeviceType       string                `gorm:"type:varchar(64);not null;default:'';comment:'设备类型'" json:"device_type"`
    Protocol         string                `gorm:"type:varchar(16);not null;default:'';comment:'接入协议:mqtt/tcp/http/ws'" json:"protocol"`
    ThingModel       string                `gorm:"type:text;comment:'物模型JSON'" json:"thing_model"`
    DecodeScript     string                `gorm:"type:text;comment:'解码脚本(goja): decode(raw)返回物模型对象'" json:"decode_script"`
    FilterScript     string                `gorm:"type:text;comment:'数据过滤脚本(goja): filter(values)返回布尔'" json:"filter_script"`
    AlertScript      string                `gorm:"type:text;comment:'数据告警脚本(goja): alert(values)返回级别数字(0=不告警,-1=故障/1=轻微/2=严重)'" json:"alert_script"`
    AlertDuration    int                   `gorm:"not null;default:0;comment:'告警持续时长(秒),0=命中即告'" json:"alert_duration"`
    AlertSilence     int                   `gorm:"not null;default:0;comment:'告警静默时长(秒),0=不静默'" json:"alert_silence"`
    IsDelete         soft_delete.DeletedAt `gorm:"not null;default:0;softDelete:flag,DeletedAtField:DeleteTime;comment:'是否删除'" json:"-"`
    CreateTime       x_null.Time           `gorm:"autoCreateTime;not null;comment:'创建时间'" json:"create_time"`
    UpdateTime       x_null.Time           `gorm:"autoUpdateTime;not null;comment:'更新时间'" json:"update_time"`
    DeleteTime       x_null.Time           `gorm:"default:null;comment:'删除时间'" json:"-"`
}

func (IotProduct) TableName() string { return "x_iot_product" }

func (m *IotProduct) BeforeCreate(tx *gorm.DB) error {
    m.ID = uuid.NewV7().String()
    return nil
}
```

建表方式：开发期用 `core.AutoMigrate(...)` 自动迁移（`core/db.go` 已提供）；生产按 `docs/sql` 惯例产出 `x_iot_*.sql` 脚本。

---

## 五、TDengine 时序数据（新增，第四阶段引入）

超级表建模不变，按设备类型/产品建 STable，每设备一张子表，tag 调整为字符串标识以匹配 UUID 主键体系：

```sql
CREATE STABLE sensor_data (
    ts TIMESTAMP,
    temperature FLOAT,
    humidity FLOAT,
    voltage FLOAT
) TAGS (
    device_id BINARY(64),
    product_id BINARY(64),
    space_id BINARY(64)
);
```

Go 侧通过 `database/sql` + `taosSql` 驱动，独立封装为 `core/td`（参照 `core/db.go` 的连接池写法，`maxOpen=4~8`）。
写入利用自动建子表：

```go
_, err := td.Exec(
    "INSERT INTO d_? USING sensor_data TAGS(?,?,?) VALUES(?,?,?,?)",
    deviceID, deviceID, productID, spaceID, time.Now(), temp, humi, volt)
```

**MySQL / Redis / TDengine 分工**：

| 数据类别 | 存储 | 承载 |
|:---|:---|:---|
| 元数据 | MySQL | 分类、产品（含三段 goja 脚本与告警持续/静默参数）、空间、设备 |
| 运行态 | Redis（`core.Redis`） | 设备在线状态、告警活动态/持续计时/静默、去重/限流 |
| 时序数据 | TDengine（第四阶段引入，现由 MySQL `x_iot_device_data` 暂代） | 传感器/指标历史数据 |
| 告警记录 | MySQL | 告警历史 |

---

## 六、实时推送（直接复用 `core.Ws`）

原方案的自建 `MapGateway + hub.Broadcast` **无需新建**，直接用现有 `core.Ws`：

- 设备地图大屏 = 一个 Ws **房间**（如 `iot:map:{spaceId}`），前端连 `/api/ws` 后加入房间；
- 位置更新：`core.Ws` 提供向 room/user/all 广播能力，pipeline 处理后按房间推送 `LocationMessage`；
- **集群天然支持**：`core/ws` 经 `core/pubsub`（Redis）跨实例广播，多节点部署无需额外改造；
- 前端 MapLibre 收到消息后 `source.setData` 更新 GeoJSON，逻辑与原方案一致。

坐标系：WGS84 可直接渲染；GCJ-02 在服务端纠偏后再推。离线地图可自建瓦片服务（如 Martin）。

---

## 七、消息洪峰削峰（复用 `core.Queue`）

接入层（MQTT/TCP/HTTP）收到设备报文后，**不直接同步处理**，而是 `core.Queue.Enqueue("iot:ingest", payload)` 投递；
`iot/pipeline` worker 通过 `core.Queue.Consume(ctx, "iot:ingest", handler, concurrency)` 并发消费，实现接入与处理解耦。
延迟/定时类任务（如超时下线判定）用 `core/queue_delay`。

---

## 八、告警引擎（Redis 状态机，复用 `core.Redis` + `ClusterMode`）

沿用 Redis ZSET + Hash + Lua + 分布式扫描器方案，实现「持续 N 秒才推送」状态机；规则配置与历史存 MySQL，Redis 仅承载运行态。

| Key（带 `config.RedisConfig` 前缀） | 类型 | 用途 |
|:---|:---|:---|
| `alert:{ruleID}:pending` | ZSET | member=deviceID，score=首次满足时间+Duration |
| `alert:{ruleID}:state:{deviceID}` | Hash | status、start_time、last_seen |
| `alert:{ruleID}:silence:{deviceID}` | String | 静默期，TTL 控制 |
| `alert:lock:{ruleID}` | String | 扫描器分布式锁 |

流程：条件首次满足 → Lua 写 ZSET+Hash；持续满足 → 仅更新 `last_seen`；条件中断 → `ZREM`+`DEL`；扫描器按 score 取出到期设备并校验 `last_seen` 容忍窗口 → 触发告警。所有状态变更用 Lua 保证原子性。
**集群下防重复扫描**：复用 `config.AppConfig.ClusterMode` 判定，分布式锁 `alert:lock:*` 保证同一规则同一时刻仅一个实例扫描。

---

## 九、API 与路由注册（复用现有惯例）

CRUD 接口按 `fabu_route.go` 写法，在 `routes/admin_route/iot_route.go` 中注册，`init()` 自动追加 `routeHandlers`：

```go
func initIotRoute(rg *gin.RouterGroup) {
    cate := iot_controller.IotProductCategoryHandler{}
    prod := iot_controller.IotProductHandler{}
    space := iot_controller.IotDeviceSpaceHandler{}
    dev := iot_controller.IotDeviceHandler{}

    auth := rg.Group("/iot", middleware.LoginAuth())
    {
        auth.GET("/productCategory/list", cate.List)   // 配合 PermAuth 权限标识
        auth.POST("/productCategory/add", cate.Add)
        auth.POST("/productCategory/edit", cate.Edit)
        auth.POST("/productCategory/del", cate.Del)
        // product / deviceSpace / device 同理
    }
}

func init() { routeHandlers = append(routeHandlers, initIotRoute) }
```

Controller 用 `util.VerifyUtil.VerifyQuery/VerifyBody` 校验、`response.JSON` 统一返回；Service 用包级单例 + `core.GetDB()` + `response.CheckErr`，分页返回 `{count, lists}`（与 `fabu_service` 完全一致）。
权限验证走 `middleware.PermAuth()` + `x_system_auth_menu`/`x_system_auth_perm` 菜单权限表。

---

## 十、开发路线图（结合代码生成器）

| 阶段 | 内容 | 复用点 |
|:---|:---|:---|
| **第一阶段** | MySQL 元数据 CRUD：产品分类 / 产品 / 设备区域 / 设备，四表建模 + 后端分层 + admin 前端页面 + 菜单权限 SQL | 可用内置代码生成器按表生成骨架，再补业务规则（树 path、删除级联校验、access_token 生成） |
| 第二阶段 | MQTT/HTTP 接入 + goja 预处理与基础过滤，报文入 `core.Queue` | `core.Queue` |
| 第三阶段 | TCP/WebSocket 接入，完善协议适配器（粘包：固定帧头+长度字段） | `gorilla/websocket` |
| 第四阶段 | TDengine 集成，时序写入与查询 | 新增 `core/td` |
| 第五阶段 | ✅ 已完成：告警状态机（持续 N 秒、静默、升级、恢复）——产品 goja 告警脚本逐条返回级别数字 + Redis 运行态收敛成事件级记录（单事件可升级），写 `x_iot_alert_history`、按转换点推 `iot:alarm` | `core.Redis`（report 驱动，无需到期表扫描） |
| 第六阶段 | MapLibre 设备地图 + 大屏实时推送 + 告警历史 | `core.Ws` 房间广播 |
| 第七阶段 | 权限细化、集群化压测、告警治理 | 现有权限体系 |

---

## 十一、关键难点与优化

| 难点 | 方案 |
|:---|:---|
| goja 性能 | Runtime 池化 + 脚本预编译 + 超时中断 |
| TCP 粘包 | 固定帧头 + 长度字段 |
| Redis 告警扫描 | ZSET 按 score 范围查询，Lua 原子取出，分布式锁防重复（集群复用 ClusterMode） |
| TDengine 写入 | 批量写入 + 独立连接池（maxOpen=4~8），自动建子表 |
| MySQL 性能 | 元数据查询加索引，告警历史分表/分区 |
| MapLibre 实时更新 | 复用 `core.Ws` 房间广播 + GeoJSON `setData` |
| 消息洪峰 | 接入层 → `core.Queue` → pipeline worker 削峰 |
| UUID 主键索引 | 采用 UUIDv7（时间有序），主键索引局部性优于随机 v4 |
