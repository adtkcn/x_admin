---
name: "管理端前端项目技术文档"
description: "x_admin admin 管理端前端（Vue3/Vite/Element Plus）导航 + 技术文档 — 目录结构、分层约定、业务域、组件与 hooks、开发规范"
---

# x_admin 管理端前端技术文档

> 导航：后端情况见 [后端导航与项目技术文档](../../../server/.agents/skills/server-project/SKILL.md)

## 项目概述

x_admin 后台管理前端，**Vue 3.5 + Vite 8 + TypeScript + Element Plus 2.14 + Pinia**，对接后端 `/api/admin` 接口。页面按业务域组织，与后端 admin 域（system/iot/flow/monitor/fabu…）一一对应。

- **包管理**: npm（`package-lock.json`），脚本见 `package.json`
- **入口**: `src/main.ts`；路由守卫在 `src/permission.ts`
- **环境**: `.env.development`（`VITE_APP_BASE_URL` 指向后端，如 `http://127.0.0.1:18000`；`VITE_MAP_TILES_URL` 地图瓦片）

```bash
cd admin
npm install
npm run dev          # 开发（vite）
npm run build        # 构建
npm run type-check   # vue-tsc 类型检查
npm run lint         # eslint --fix
```

## 技术栈要点

| 类别 | 技术 |
|------|------|
| UI | Element Plus + vxe-table（大数据表格）+ Tailwind CSS 4 |
| 状态 | Pinia（`stores/modules/`: user / app / setting / multipleTabs） |
| 请求 | axios 封装于 `utils/request/`（含 `Pages` 分页类型） |
| 图编辑/流程 | @logicflow/core+extension（规则链可视化）、@form-create/designer（审批流表单设计） |
| 地图 | maplibre-gl + @mapbox/mapbox-gl-draw（含 maplibre 适配补丁，见 map-picker / map-area-picker 组件） |
| 图表 | echarts 6（按需注册于 `utils/echart.ts`） |
| 上传 | @aws-sdk/client-s3 + lib-storage + spark-md5（分片直传，对应后端 `plugin/storage`） |
| 组件注册 | `unplugin-vue-components` 自动注册（`components.d.ts` 自动生成）；**AutoImport 已注释关闭，API/hook 需显式 import** |
| 图标 | `vite-plugin-svg-icons` 加载 `assets/icons/*.svg`（symbol 方式） |

## 目录结构

```
admin/src/
├── api/                  # 接口层：按域分子目录，与后端域对应
│   ├── iot/              #   device/product/product_category/device_space/rule_chain/alert
│   ├── flow/             #   审批流（apply/history/template）
│   ├── monitor/          #   监控（client/error/project）
│   ├── perms/            #   权限（admin/menu/role）
│   ├── setting/          #   设置（dict/notice/system/website）
│   ├── system/           #   corn 定时任务 / user
│   ├── org/              #   部门/岗位
│   ├── tools/            #   代码生成器
│   ├── user/             #   用户协议
│   └── *.ts              #   根级：app/file/album_cate/fabu/fabu_download 等
├── views/                # 页面：按业务域分目录（与 api 域同名）
│   ├── iot/              #   alert/device/device_space/monitor/product/product_category/rule_chain
│   ├── system/           #   admin/corn/dept/dict/menu/post/role/system/website
│   ├── fabu/             #   App 发布（含下载页）
│   ├── flow/             #   审批流
│   ├── monitor/          #   监控
│   ├── dev_tools/        #   代码生成器
│   └── workbench/ account/ material/ user/ error/
├── components/           # 全局自定义组件（kebab-case 目录，自动注册）
│   ├── map-picker/ map-area-picker/        # 地图取点/区域绘制（maplibre + mapbox-gl-draw）
│   ├── rule-chain/ flow/                   # 规则链编排（LogicFlow）/ 审批流设计（FcDesigner）
│   ├── upload/ upload-chunk/               # 上传 / S3 分片直传
│   ├── echart-component/ highlight-code/   # 图表 / 代码高亮
│   ├── verify/ material/ import-excel/     # 行为验证码 / 素材 / Excel 导入
│   └── pagination/ popup/ dict-value/ …    # 通用小组件
├── hooks/                # usePaging(分页) useDictOptions(字典下拉) useGlobalWs(全局WebSocket) useMultipleTabs useWatchRoute useLockFn useReactiveWithReset
├── stores/modules/       # Pinia：user(登录/菜单/权限) app multipleTabs setting
├── router/               # index/routes(静态路由) + routes_utils(按后端菜单动态生成路由 filterAsyncRoutes/loadRouteView)
├── utils/                # request(axios封装) auth(令牌) cache s3Client echart theme feedback validate …
├── enums/                # appEnums/requestEnums/cacheEnums/pageEnum/fileEnums
├── config/               # index(运行时配置/环境变量) setting(界面默认设置)
├── install/              # 插件注册(plugins/: element-plus/vxe-table/pinia/router/el-icon) + 指令(directives/: perms.ts → v-perms 权限指令, copy)
├── layout/               # 主布局（侧边栏/多标签 Tab）
├── styles/               # element.scss / tailwind.css / var.css / dark.css
└── assets/               # geojson(中国/省/市边界，地图区域用) icons(svg 图标库) images
```

## 关键开发约定

1. **新增页面四件套**：`views/<域>/xxx/`（页面）+ `api/<域>/xxx.ts`（接口与类型）+ 后端菜单配置（动态路由按菜单 component 路径加载 `views` 下文件）+ 权限标识（`v-perms` 指令控制按钮，与后端 `admin:域:模块:操作` 对应）。
2. **api 模块惯例**：`import request from '@/utils/request'`；类型以 `type_<域>_<模块>` 命名导出（如 `type_iot_device`），分页参数 `pageNo/pageSize` 与后端 `PageReq` 对齐。
3. **显式导入**：AutoImport 已关闭，`ref/computed/watch` 及 Element Plus API 均需从 `vue` / `element-plus` 显式 import；组件因 unplugin-vue-components 免 import。
4. **状态字段渲染**：统一走字典（`useDictOptions` + `dict-value` 组件），不在页面硬编码枚举文案。
5. **代码风格**：4 空格缩进（`.prettierrc`），提交前 `npm run lint` + `npm run type-check`。

## 相关文档导航（`docs/admin/`）

- [准备](../../../../docs/admin/准备.md) · [路由](../../../../docs/admin/路由.md) · [app更新](../../../../docs/admin/app更新.md)
- [单文件上传](../../../../docs/admin/单文件上传.md) · [分片上传](../../../../docs/admin/分片上传.md) · [WebSocket](../../../../docs/admin/WebSocket.md)
- [自定义hooks](../../../../docs/admin/自定义hooks.md) · [嵌入iframe](../../../../docs/admin/嵌入iframe.md) · [nginx配置](../../../../docs/admin/nginx配置.md)
