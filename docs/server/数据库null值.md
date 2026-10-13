---
name: "后端-数据库Null值"
description: "x_admin 后端 x_null 可空类型的作用、可用类型与查询/模型/序列化用法示例"
---

## 为什么需要 x_null

数据库里 `int`、`string`、`datetime` 等列常常允许 `NULL`，但 Go 的原生 `int`/`string`/`time.Time` **不可空**，两者的空值语义对不上，直接映射会报错。原生零值又无法区分"用户没填"和"用户填了 0 / 空串"——这在**列表筛选**里尤其致命（`status=0` 到底是"筛选离线"还是"不筛选"？）。

`x_admin` 统一使用外部库 **`github.com/adtkcn/x_null`**（见 `server/go.mod`，当前 `v0.0.2`）来解决：它提供一组"可空 + 可区分未设置"的类型，一处声明即同时打通 **数据库读写、JSON 序列化、Gin 参数绑定、校验器、格式化输出**。

## x_null 做了什么

每个 `x_null` 类型都是一个 `{ Val, Valid }` 结构体，并一次性实现了各路框架需要的接口，业务侧无需再手写：

| 面向 | 实现的接口 | 效果 |
| --- | --- | --- |
| GORM / database/sql | `Scan(any) error`、`Value() (driver.Value, error)` | DB 的 `NULL` ↔ Go 空值自动互转 |
| encoding/json | `MarshalJSON`、`UnmarshalJSON` | 响应里未设置输出 `null`；前端可传 `null` |
| Gin 参数绑定 | `UnmarshalParam(param string) error` | 地址栏 / 表单 query 直接绑定到可空字段 |
| go-playground/validator | 配合 `RegisterCustomTypeFunc` | `binding:"required"` 等校验对可空类型依然生效 |
| fmt 等 | `String() string` | 日志/格式化输出友好 |

> 相比标准库 `sql.NullString/NullInt64`，`x_null` 额外补齐了 **JSON 双向、Gin 绑定、数字/字符串数字/null 兼容、`ValueOrZero()` 便捷取值、excel `Decode*`**，因此全项目统一用它。

## 可用类型与常用方法

```go
x_null.String   // 可空字符串
x_null.Int64    // 可空整数（支持前端传 null / 字符串数字 / 数字）
x_null.Float64  // 可空浮点
x_null.Time     // 可空时间
```

| 成员 | 说明 |
| --- | --- |
| `.Val` | 底层值（`string` / `int64` / `float64` / `time.Time`） |
| `.Valid` | 是否被设置（`false` = 未提供，与"设成零值"区分开） |
| `.ValueOrZero()` | 取值；未设置时返回对应零值，免去判 `.Valid` |
| `x_null.NewInt64(1)` | 构造器（各类型均有 `NewXxx`），显式置为某个值 |
| `x_null.DecodeString` / `DecodeInt64` / `DecodeTime` | Excel 导入列的值解码函数（见 `导入导出excel`） |

## 用法示例

### 1）模型层：可空列

时间类字段用 `x_null.Time` + `default:null`，未发生时保持 `NULL` 而非零值时间：

```go
// app/model/iot_model/device.go
type IotDevice struct {
    // ...
    Status          uint8       `gorm:"not null;default:0" json:"status"`
    LastOnlineTime  x_null.Time `gorm:"default:null;comment:'最后上线时间'" json:"last_online_time"`
    LastOfflineTime x_null.Time `gorm:"default:null;comment:'最后离线时间'" json:"last_offline_time"`
    CreateTime      x_null.Time `gorm:"autoCreateTime;not null" json:"create_time"`
    UpdateTime      x_null.Time `gorm:"autoUpdateTime;not null" json:"update_time"`
}
```

响应 DTO 同样用 `x_null.Time`，未上线的设备 `last_online_time` 自然输出 `null`（见 `iot_schema.DeviceResp`）。

### 2）查询条件："免哨兵"筛选（核心用法）

列表筛选项统一声明为 `x_null.String`：**不传 / 传空串都视为"不筛选"，不再需要 `-1` 之类哨兵值**（`app/schema/iot_schema/device_schema.go`）：

```go
// 查询条件统一用 x_null.String：不传/空串均视为"不筛选"，无需 -1 之类的哨兵值
type DeviceListReq struct {
    Name     x_null.String `json:"name" form:"name"`     // 模糊
    Status   x_null.String `json:"status" form:"status"` // 0离线 1在线（0 不再与"不筛选"混淆）
    ProductId x_null.String `json:"product_id" form:"product_id"`
    // ...
    request.PageReq
}
```

服务层用 `ValueOrZero()` 判空决定是否拼 `WHERE`（`rule_chain_service.go` 同款写法）：

```go
if v := listReq.Name.ValueOrZero(); v != "" {
    chain = chain.Where("name like ?", "%"+v+"%")
}
if v := listReq.IsStop.ValueOrZero(); v != "" {
    chain = chain.Where("is_stop = ?", v)
}
```

### 3）可空数字：程序内构造或前端传 null

整数筛选用 `x_null.Int64`，不设置字段 = 不筛选，需要显式条件时用 `NewInt64`：

```go
// app/corn/dynamic_task.go：只取启用(1)的定时任务
corn_service.SystemCornService.ListAll(schema.SystemCornListReq{
    Status: x_null.NewInt64(1),
})
```

`x_null.Int64` / `Float64` 兼容前端把值传成 `null`、字符串数字 `"12"` 或数字 `12`（`monitor_error_ctl.go` 里宽高即用它承载可空数值）。

### 4）校验器集成

启动时显式注册（`bootstrap.go` → `util/validation.Register()`），让 `x_null` 拆成底层值再交给 validator，`binding:"required"` 等标签照常生效：

```go
// util/validation/null.go
func validateValuer(field reflect.Value) any {
    switch obj := field.Interface().(type) {
    case x_null.String:  return obj.Val
    case x_null.Int64:   return obj.Val
    case x_null.Float64: return obj.Val
    case x_null.Time:    return obj.Val
    default:             return nil
    }
}
v.RegisterCustomTypeFunc(validateValuer,
    x_null.String{}, x_null.Int64{}, x_null.Float64{}, x_null.Time{})
```

## 注意事项

- **`0` 与"未设置"是两回事**：可空字段的"未填"由 `.Valid=false` 表达，别再用零值/`-1` 当哨兵；这正是查询免哨兵的基础。
- `x_null.Time` 的 `autoCreateTime/autoUpdateTime` 列可写 `not null`（GORM 总会填），而业务性可空时间（如最后离线时间）用 `default:null`。
- 新增可空类型时，需在 `validateValuer` 的 `switch` 与 `RegisterCustomTypeFunc` 的类型列表里同步登记，否则校验拿不到底层值。

## 参考

1. 标准库可空类型思路：https://pkg.go.dev/database/sql#NullInt64
2. 同类实现（guregu/null）：https://github.com/guregu/null
3. Gin 自定义参数绑定：`UnmarshalParam` 接口
