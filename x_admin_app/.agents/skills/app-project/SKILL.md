---
name: "App端项目技术文档"
description: "x_admin_app 移动端（uni-app x / UTS）导航 + 技术文档 — 目录结构、页面域、请求与升级 SDK、hb-cli 打包"
---

# x_admin_app 移动端技术文档

> 导航：uvue/UTS 编码规范见同级 [uvue skill](../uvue/SKILL.md)；后端见 [后端导航与项目技术文档](../../../../server/.agents/skills/server-project/SKILL.md)；仓库整体见 [项目总导航](../../../../.agents/skills/project-navigation/SKILL.md)。

## 项目概述

x_admin 配套移动端 App，基于 **uni-app x**（UTS 语言 + uvue 页面，Vue3 语法，编译为原生 Kotlin/Swift），对接后端 `/api/admin` 接口（登录鉴权复用后台 `Token`）。

- **appid**: `__UNI__FB29F21`（manifest.json）
- **开发工具**: HBuilderX；CLI 打包见 `hb-cli.config.js` + `@adtkcn/hb-cli`
- **接口地址**: `utils/env.ts` 中 `baseUrl`（如 `http://192.168.11.5:18000`）+ `urlPrefix: '/api/admin'`
- **UI 库**: `uni_modules/` 下 uv-ui 全家桶 + wot-design-uni（easycom 自动注册）

```bash
npm install
npm run build          # hb-cli 打包（安卓/iOS，改版本号、wifi 调试）
npm run build:prod     # hb-cli --mode prod
```

## Skills 导航

| Skill | 什么时候用 |
|-------|-----------|
| [uvue](../uvue/SKILL.md) | **写页面/组件必读**：uni-app x 的 UTS 严格类型、uvue 页面规范、条件编译、滚动容器、easycom 等编码约定 |

## 目录结构

```
x_admin_app/
├── pages/                  # 页面（pages.json 注册；tabbar：index 工作台 / my 我的）
│   ├── index/              #   首页工作台（图标资源 static/index：设备/产品/日志等入口）
│   ├── login/              #   登录
│   ├── my/                 #   我的
│   ├── monitor_project/    #   项目监控（index/details/edit/search）
│   ├── monitor_client/     #   客户端监控（index/details/edit/search，含独立 pages.json 分包配置）
│   └── video/              #   视频监控（播放源 static/monibuca）
├── api/                    # 接口模块：system/user、monitor_client、monitor_project
├── components/             # 自定义组件（x- 前缀：x-list/x-picker/x-date/x-card/x-upgrade 升级组件；verify 行为验证码）
├── hooks/                  # usePaging(分页) useDictOptions(字典) —— 与 admin 前端同名同职责
├── stores/                 # 简单状态（user.ts；注意 uni-app x 不支持 pinia，仅作模块封装）
├── methods/                # 全局方法（配合 typings/methods.d.ts）
├── utils/                  # request(uni.request封装+Token头) env(接口地址) storage perms(权限) wss(WebSocket) map emtt number_util
├── sdk/upgrade/            # App 升级 SDK：对接后端 fabu 发布模块（vue3/ 端 + uniappx/ uts 端 + react-native 预留）
├── static/                 # 静态资源（tabs 图标、index 入口图标、monibuca 播放器）
├── uni_modules/            # uv-ui 组件库全家桶 + wot-design-uni
├── typings/                # uniapp.d.ts / methods.d.ts 类型声明
├── pages.json              # 页面路由 + tabbar + 全局样式
├── manifest.json           # App 配置（appid/权限/打包）
├── androidPrivacy.json     # 安卓隐私合规弹窗
├── hb-cli.config.js        # 打包工具配置（版本/wifi 调试/产物）
└── unpackage/              # 编译产物与缓存（不提交）
```

## 关键开发约定

1. **新增页面**：`pages/` 下建目录 → `pages.json` 注册（标题、下拉刷新等）；tab 页仅 index/my。
2. **请求**：统一走 `utils/request.ts` 的 `request<T>()`（自动拼 `env.baseUrl + env.urlPrefix`，Header 带 `Token`，401 过期处理）；上传用同文件 `upload()`。
3. **编码规范**：UTS 严格类型、条件编译隔离平台代码、滚动内容必须放 scroll-view/list-view —— 细则见 [uvue skill](../uvue/SKILL.md)。
4. **App 更新**：`components/x-upgrade`（.vue/.uvue 双端同名组件）+ `sdk/upgrade/`，对接后端 fabu 模块的版本检查/下载/安装（WGT 热更与整包更新）。
5. **组件**：优先 uv-ui / wot-design-uni（easycom 免 import）；自研组件 `x-` 前缀放 `components/`。
