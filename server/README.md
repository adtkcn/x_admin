### 接口路径说明

所有接口`api/`开头,方便反向代理

后台接口`api/admin`开头，后台权限及菜单设计以这个为主

前台接口`api/*`开头，无权限设计，需要自己实现

### 目录结构

```
├── server
│   ├── admin // 后台,可以参照添加web，app等模块
│   │   ├──** //模块
│   │   ├──────/*_ctl.go // 模块控制器
│   │   ├──────/*_schema.go // 模块schema
│   │   ├──────/*_service.go // 模块服务
│   ├── config // 配置
│   ├── middleware // 中间件
│   ├── model // 数据库模型
│   ├── util // 工具包
│   ├── routers // 路由`api`,
│   ├── resources // 验证码依赖文件，待调整
│   ├── static // 静态文件，访问路由`/api/static/*`
│   ├── main.go // 入口
│   ├── .env // 配置文件，注意不提交git
```
