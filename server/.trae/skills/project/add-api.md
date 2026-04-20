---
name: "新增接口"
description: "在已有模块中新增接口的完整指南，含 Schema/Service/Controller/Route 四件套模板"
---

# 新增接口指南

## 概述

在 x_admin 项目已有的业务模块中**新增一个 CRUD 接口**，需要修改/创建 **4 个文件**（Schema / Service / Controller / Route），严格遵循项目既定的代码模式和命名规范。

## 快速步骤总览

```
1. 扩展 Schema    → app/schema/xxx_schema.go  （新增请求/响应结构体）
2. 扩展 Service   → app/service/xxx_service.go  （新增业务方法）
3. 扩展 Controller → app/controller/admin_ctl/xxx_ctl.go  （新增 Handler 方法）
4. 注册路由       → routes/admin_route/xxx_route.go  （新增路由 + init 追加）
```

---

权限标识命名规则：假如访问路径/api/admin/system_corn/list，去掉`/api/`部分，`/`转为`:`,得到`admin:system_corn:list`，其中`admin`表示后台接口，`system_corn`表示模块名，`list`表示操作

---

---

## 第 1 步：扩展 Schema

**文件**: `app/schema/<module>_schema.go`

### Schema 字段类型规范

| 基础类型    | 使用类型                   |
| ----------- | -------------------------- |
| int 类型    | 使用 `x_null.Int64` 类型   |
| float 类型  | 使用 `x_null.Float64` 类型 |
| string 类型 | 使用 `x_null.String` 类型  |
| time 类型   | 使用 `x_null.Time` 类型    |

---

## 第 2 步：扩展 Service

**文件**: `app/service/<module>_service.go`

### Service 编码规范

| 要点     | 规则                                                            |
| -------- | --------------------------------------------------------------- |
| 错误处理 | 统一使用 `response.CheckErr()` 和 `response.CheckDBNotRecord()` |
| 类型转换 | Model ↔ Schema 用 `convert_util.Copy()`                         |

---

## 第 3 步：扩展 Controller

**文件**: `app/controller/admin_ctl/<module>_ctl.go`

### Controller 编码规范

| 要点          | 规则                                                            |
| ------------- | --------------------------------------------------------------- |
| GET 参数校验  | `util.VerifyUtil.VerifyQuery(c, &req)`                          |
| POST 参数校验 | `util.VerifyUtil.VerifyBody(c, &req)`                           |
| 响应          | 统一使用 `response.CheckAndRespWithData(c, data, err)`          |
| Swagger 注释  | 每个接口必须有 `@Summary` `@Tags` `@Param` `@Success` `@Router` |
| Tags 格式     | `<module>-<中文名称>`（与后台菜单对应）                         |

---

## 第 4 步：注册路由

**文件**: `routes/admin_route/<module>_route.go`

### 路由注册完整模板

```go 
func <ModelName>Route(rg *gin.RouterGroup) {
	handle := admin_ctl.<ModelName>Handler{}
    // 不需要鉴权的路由在rg.Group前注册
	r = rg.Group("/", middleware.PermAuth())
    // 需要鉴权的路由在rg.Group后注册
}
```

### Route 编码规范

| 要点      | 规则                                                                          |
| --------- | ----------------------------------------------------------------------------- |
| 中间件    | 需要权限校验用 `middleware.PermAuth()`，仅登录用 `middleware.LoginAuth()`     |
| 自动加载  | 必须在文件末尾写 `init()` 将路由函数追加到 `routeHandlers`                    |
| HTTP 方法 | GET 用于查询(list/detail)，POST 用于写入(add/edit/del)                        |
| 权限标识  | 注释中标明，格式为 `admin:模块名:操作` (`admin:xxx:list`, `admin:xxx:add` 等) |

## 完整示例参考

以下是一个标准 CRUD 接口的五件套文件清单（以 system_corn 为例,包含`list` / `listAll`/ `detail` / `add` / `edit` / `del`/ `del_batch` / `export_file`/ `import_file`接口）:

```
app/model/system_corn.go
app/schema/system_corn_schema.go
app/service/system_corn_service.go
app/controller/admin_ctl/system_corn_ctl.go
routes/admin_route/system_corn_route.go
```

---

## 常见快捷操作

### 仅新增一个查询接口

只需修改 4 个文件:

1. **Schema**: 新增 `XxxQueryReq` 结构体
2. **Service**: 新增 `Query()` 方法
3. **Controller**: 新增 `Query()` Handler 方法 + Swagger 注释
4. **Route**: 在现有 Route 函数中追加一行 `rg.GET("/xxx/query", handle.Query)`

### 新增非 CRUD 自定义接口

同样遵循上述四件套模式，区别在于:

- Service 方法名和逻辑自定义
- 路由路径和 HTTP 方法根据语义选择
- 权限标识自定义为 `admin:xxx:customAction`
