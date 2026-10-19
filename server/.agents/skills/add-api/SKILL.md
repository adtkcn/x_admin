---
name: "新增接口"
description: "在已有模块中新增接口的完整指南，含 Schema/Service/Controller/Route 四件套模板"
---

# 新增接口指南

## 概述

在 x_admin 项目已有的业务模块中**新增一个 CRUD 接口**，需要修改/创建 **4 个文件**（Schema / Service / Controller / Route），严格遵循项目既定的代码模式和命名规范。业务域（iot/system/flow/…）下目录均按 `<域>_schema|service|controller(子包)/route` 分目录，新文件放入对应域目录。

## 快速步骤总览

```
1. 扩展 Schema    → app/schema/<域>_schema/xxx_schema.go   （新增请求/响应结构体）
2. 扩展 Service   → app/service/<域>_service/xxx_service.go （新增业务方法）
3. 扩展 Controller → app/controller/admin_ctl/<域>_controller/xxx_ctl.go （新增 Handler 方法）
4. 注册路由       → routes/admin_route/<域>_route.go        （新增路由 + init 追加）
```

> C端接口对应 `web_ctl/` + `web_route/`（鉴权用 `UserLoginAuth()`）；IoT 设备接入网关对应 `iot_ctl/` + `iot_route/`（鉴权用 `DeviceAuth()`）。

---

权限标识命名规则：假如访问路径/api/admin/iot/device/list，去掉`/api/`部分，`/`转为`:`,得到`admin:iot:device:list`，其中`admin`表示后台接口，`iot:device`表示模块，`list`表示操作

---

---

## 第 1 步：扩展 Schema

**文件**: `app/schema/<域>_schema/<module>_schema.go`

### Schema 字段类型规范

| 基础类型    | 使用类型                   |
| ----------- | -------------------------- |
| int 类型    | 使用 `x_null.Int64` 类型   |
| float 类型  | 使用 `x_null.Float64` 类型 |
| string 类型 | 使用 `x_null.String` 类型  |
| time 类型   | 使用 `x_null.Time` 类型    |

> 查询条件统一用 `x_null` 类型：不传/空串均视为“不筛选”，无需 -1 之类哨兵值（导入 `github.com/adtkcn/x_null`）。

---

## 第 2 步：扩展 Service

**文件**: `app/service/<域>_service/<module>_service.go`

### Service 编码规范

| 要点     | 规则                                                            |
| -------- | --------------------------------------------------------------- |
| 错误处理 | 普通查询用 `response.CheckErr(err, "统计xx失败")`；查库 First 等用 `response.CheckDBErr(err, "xx不存在!", "详情获取失败")` 区分未找到/失败 |
| 业务错误 | `response.AssertArgumentError.SetMessage("网关设备不能是自己!")` 等预定义 RespType |
| 类型转换 | Model ↔ Schema 用 `convert_util.Copy()`                         |
| 分页返回 | `response.PageResp{Count, Lists}`，入参 `request.PageReq`       |

---

## 第 3 步：扩展 Controller

**文件**: `app/controller/admin_ctl/<域>_controller/<module>_ctl.go`（`type XxxHandler struct{}`）

### Controller 编码规范

| 要点          | 规则                                                            |
| ------------- | --------------------------------------------------------------- |
| GET 参数校验  | `if response.IsFail(c, util.VerifyUtil.VerifyQuery(c, &req)) { return }` |
| POST 参数校验 | `if response.IsFail(c, util.VerifyUtil.VerifyBody(c, &req)) { return }`  |
| 响应          | 统一使用 `response.JSON(c, data, err)`（内部自动处理 err 转响应）        |
| Swagger 注释  | 每个接口必须有 `@Summary` `@Tags` `@Param` `@Success` `@Router`          |
| Tags 格式     | `<module>-<中文名称>`（与后台菜单对应）                                  |
| Success 写法  | 分页：`response.Response{data=response.PageResp{lists=xx_schema.XxxResp}}`；单条：`response.Response{data=xx_schema.XxxResp}` |

---

## 第 4 步：注册路由

**文件**: `routes/admin_route/<域>_route.go`（一个域一个文件，如 `iot_route.go`）

### 路由注册完整模板

```go
// initXxxRoute Xxx 管理路由
func initXxxRoute(rg *gin.RouterGroup) {
	handle := xxx_controller.XxxHandler{}

	// 仅需登录（如表单下拉数据 all 类接口）
	look := rg.Group("/<域>", middleware.LoginAuth())
	{
		look.GET("/xxx/all", handle.All)
	}

	// 需权限认证：CRUD 写入接口
	auth := rg.Group("/<域>", middleware.PermAuth())
	{
		auth.GET("/xxx/list", handle.List)
		auth.POST("/xxx/add", middleware.RecordLog("xx新增"), handle.Add)
		auth.POST("/xxx/edit", middleware.RecordLog("xx编辑"), handle.Edit)
		auth.POST("/xxx/del", middleware.RecordLog("xx删除"), handle.Del)
	}
}

func init() {
	routeHandlers = append(routeHandlers, initXxxRoute)
}
```

### Route 编码规范

| 要点      | 规则                                                                          |
| --------- | ----------------------------------------------------------------------------- |
| 中间件    | 列表/详情/写入按权限用 `middleware.PermAuth()`；仅需登录的下拉类用 `middleware.LoginAuth()`；C端用 `UserLoginAuth()`，设备接入用 `DeviceAuth()` |
| 操作审计  | add/edit/del 等写操作包一层 `middleware.RecordLog("xx新增")` 记录操作日志       |
| 自动加载  | 必须在文件末尾写 `init()` 将路由函数追加到 `routeHandlers`                    |
| HTTP 方法 | GET 用于查询(list/detail/all)，POST 用于写入(add/edit/del)                    |
| 权限标识  | 注释中标明，格式为 `admin:<域>:<模块>:<操作>` 或按菜单层级 `admin:xxx:list`    |

## 完整示例参考

以下是一个标准 CRUD 接口的五件套文件清单（以 iot 设备为例,包含`list` / `all` / `detail` / `add` / `edit` / `del`等接口; 定时任务域为 `system_corn` 的五件套同理）:

```
app/model/iot_model/device.go
app/schema/iot_schema/device_schema.go
app/service/iot_service/device_service.go
app/controller/admin_ctl/iot_controller/device.go
routes/admin_route/iot_route.go
```

---

## 常见快捷操作

### 仅新增一个查询接口

只需修改 4 个文件:

1. **Schema**: 新增 `XxxQueryReq` 结构体（字段用 `x_null` 类型）
2. **Service**: 新增 `Query()` 方法
3. **Controller**: 新增 `Query()` Handler 方法 + Swagger 注释（`IsFail` + `response.JSON`）
4. **Route**: 在对应域 route 文件的 `auth`/`look` 分组中追加一行 `auth.GET("/xxx/query", handle.Query)`

### 新增非 CRUD 自定义接口

同样遵循上述四件套模式，区别在于:

- Service 方法名和逻辑自定义
- 路由路径和 HTTP 方法根据语义选择
- 权限标识自定义为 `admin:xxx:customAction`
