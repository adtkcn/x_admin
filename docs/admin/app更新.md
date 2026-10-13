# App 更新调用说明

App 更新覆盖 **三端**：管理端（`admin`）负责发布/回滚/灰度，服务端（`fabu` 模块）提供检查/计数/短链接口，App 端（`x_admin_app`）用 SDK 触发弹窗与安装。文档串起三端调用面与协议。

## 一、数据模型

```
App（应用，按 bundle_id + platform 唯一定位）
  └── Version（全量安装包 .apk/.ipa，一个 App 多版本，仅一个 released=1）
        └── Wgt（uni-app 热更资源包，挂在具体 Version 下，同版本仅一个 released=1）
```

- `App.short_url`：下载短链标识，对外的落地页 URL 用它。
- `Version.update_mode`：`0` 普通 / `1` 强制（`2` 静默已废弃，客户端一律按普通处理）。
- `Version.gray`：灰度标志，检查更新接口的过滤条件之一。
- `Wgt.version_id`：热更包归属的具体全量版本；SDK 检查时按版本 name 分段比较。

## 二、管理端 API（`admin/src/api/fabu.ts`）

| 函数 | 方法 & 路径 | 用途 |
| --- | --- | --- |
| `fabuAppLists` | GET `/fabu/app/list` | 应用列表（下拉/表格用） |
| `fabuAppDel` | POST `/fabu/app/del` | 删除应用 |
| `fabuVersionLists` | GET `/fabu/version/list` | 按 `app_id` 分页取版本 |
| **`fabuVersionUpload`** | POST `/fabu/version/upload` | **分片上传完成后**，用 `{ file_hash_id, file_name }` 通知后端解析 manifest 并登记版本（自动按 BundleId 建应用/追加版本） |
| `fabuVersionRelease` | POST `/fabu/version/release` | 发布为线上版本 |
| `fabuVersionCancel` | POST `/fabu/version/cancel` | 取消发布（回滚） |
| `fabuVersionGray` | POST `/fabu/version/gray` | 灰度开/关 |
| `fabuVersionUpdateMode` | POST `/fabu/version/updateMode` | 切换 `update_mode`（普通/强制） |
| `fabuVersionDel` | POST `/fabu/version/del` | 删除版本 |
| `fabuWgtLists` | GET `/fabu/wgt/list` | 按 `version_id` 分页取热更包 |
| **`fabuWgtUpload`** | POST `/fabu/wgt/upload` | 分片上传完成后，用 `{ version_id, file_hash_id, file_name }` 登记 wgt 记录 |
| `fabuWgtRelease` | POST `/fabu/wgt/release` | 发布 / 取消发布热更包（`released: boolean`；同版本下其它已发布包会被后端自动取消） |
| `fabuWgtDel` | POST `/fabu/wgt/del` | 删除热更包 |

> 上传落盘本身走分片上传组件（见 [分片上传](./分片上传.md)），本节所有 `upload` 接口只做"**基于 `file_hash_id` 的解析与登记**"，不再传文件。

## 三、管理端页面与操作流

| 路由 | 页面 | 关键操作 |
| --- | --- | --- |
| `views/fabu/app/index.vue` | 应用列表 | 选应用 → 跳"版本"；删除 |
| `views/fabu/version/index.vue` | 版本列表 | 上传安装包 / 发布 / 取消发布 / 切换灰度 / 切换更新模式 / 下载链接复制 / 热更新弹窗 / 删除 |
| `views/fabu/version/wgt.vue` | 热更新弹窗 | 上传 `.wgt` → 发布 / 取消 / 复制链接 / 删除 |
| `views/fabu/components/upload.vue` | 上传安装包弹窗 | 承载 `upload-chunk`，`ext=['ipa','apk']` |
| `views/fabu/download/index.vue` | 落地下载页（面向终端用户） | 按 `short_url` 拉信息 → iOS 走 `itms-services://` / Android 走计数下载 |

**典型发布流程**：

```
1) 版本页 → 「上传版本」 → upload-chunk 分片上传 .apk/.ipa → fabuVersionUpload({file_hash_id, file_name})
2) 表格出现新版本（默认未发布）→ 点「发布」置为线上版本
3) 若需热修复：该行点「热更新」 → wgt.vue → upload-chunk 上传 .wgt → fabuWgtUpload({version_id, file_hash_id, file_name}) → 点「发布」
4) 需要强制更新：改「更新模式」下拉选「强制」
5) 需要灰度：切换「灰度」开关（后端检查接口按 gray 过滤）
```

## 四、App 端 SDK 集成（`x_admin_app/sdk/upgrade/`）

SDK 分层：**协议 → 决策 → 桥接 → UI**；当前提供 `uni-app vue3` 与 `uni-app x` 两端实现。

### 服务端协议（App 主动调用）

```
GET {baseUrl}/api/web/fabu/version/checkupdate
    ?bundle_id=<包名>
    &platform=<ios|android>
    &version_code=<int>          # 客户端安装包 versionCode
    &wgt_version=<string>        # 可选，客户端当前 wgt 版本 name（如 1.0.3）
```

- **公开接口**，无需 Token，SDK 内直接 `uni.request`。
- 服务端优先级：先判全量包（`released=1` 且 `version_code > 客户端`），无全量更新时才下发同版本下最新已发布 wgt。
- 响应（`code=0` 信封的 `data`）三种固定形态：

```json
{ "update": true,  "type": "app", "app": { "version", "version_code", "download_url", "install_url", "update_mode" } }
{ "update": true,  "type": "wgt", "wgt": { "version", "version_code", "download_url", "md5", "size" } }
{ "update": false }
```

### uni-app vue3 集成（`sdk/upgrade/vue3/`）

```ts
// App.vue onLaunch
import { check } from '@/sdk/upgrade/vue3/config'
check()                       // 用 utils/env.ts 的 baseUrl；如需覆盖：check({ baseUrl: '...' })

// 常驻页面（如首页）放置弹窗组件（easycom 免导入）
// <x-upgrade></x-upgrade>
```

`check()` 可在任意时机调用；若 `<x-upgrade />` 尚未渲染，SDK 会用 `pendingCheck` 暂存，待组件 `registerCheckHandler` 后自动补执行；多页面重复渲染不会重复检查。

配置项（`sdk/upgrade/vue3/config.ts`）：

```ts
interface UpgradeConfig {
    baseUrl: string            // 服务端地址
    checkPath: string          // 默认 '/api/web/fabu/version/checkupdate'
    enableWgt: boolean         // 是否启用 wgt 热更（uni-app x 恒 false）
    bundleId: string           // 应用包名，由接入方运行时注入（SDK 不自动获取）
}
```

组件状态机 `Phase`：`HIDDEN → CONFIRM → DOWNLOADING → READY`，由 `flow.ts` 驱动。

### uni-app x 集成

```html
<template>
    <x-upgrade ref="upgradeRef"></x-upgrade>
</template>
<script setup lang="uts">
const upgradeRef = ref<ComponentPublicInstance | null>(null)
onReady(() => upgradeRef.value?.$callMethod('check'))
</script>
```

服务端地址在 `x-upgrade.uvue` 顶部 `DEFAULT_BASE_URL` 常量修改；组件自包含、**不支持 wgt**（收到 `type=wgt` 直接忽略）。

### 三条更新流程（vue3 端）

| 场景 | 行为 |
| --- | --- |
| 普通 `update_mode=0` | 弹窗提示升级 →「以后再说」可跳过（同 `version_code` 记本地 storage 不再提示）→ 同意后后台下载，弹窗显示进度 → 下载完拉起安装 |
| 强制 `update_mode=1` | 弹窗无关闭入口（遮罩 / 按钮均不可关），必须等待下载完成并安装 |
| wgt 热更 | **不弹窗**：静默下载 → `plus.runtime.install` 安装 → 仅弹「重启生效」确认框（取消则下次启动再生效） |

**平台差异**：

- wgt 仅 **Android + vue3** 支持（iOS 系统限制不允许热更自身 JS，uni-app x 无 wgt 概念）。
- iOS 全量包：`plus.runtime.openURL(install_url)`，`install_url` 为服务端 `itms-services` OTA 地址。
- Android apk 安装：vue3 用 `plus.runtime.install`；uni-app x 用 `uni.installApk`（HBuilderX 3.94+，需系统"安装未知应用"授权）。

## 五、iOS OTA / Android 直装（落地页）

`views/fabu/download/index.vue`（对外短链页，路由参数 `short_url`）：

```ts
const url =
    data.platform === 'ios'
        ? `itms-services://?action=download-manifest&url=${encodeURIComponent(origin + data.install_url)}`
        : `${origin}/api/web/fabu/count/${data.app_id}/${data.version_id}`
window.location.href = url
```

- **iOS**：`install_url` 由后端生成的 `manifest.plist` 地址，走 `itms-services://` 协议（企业 / AdHoc 分发）。
- **Android**：走 `/api/web/fabu/count/{app_id}/{version_id}` **计数下载**接口（后端记 `download_times` 后 302 到 apk 文件流）。
- 落地页应用信息通过 `fabuAppDownloadInfo(short_url)` 拉取（公开接口）。

## 六、约定与常见坑

1. **`file_hash_id` 是唯一的落库抓手**：`fabuVersionUpload` / `fabuWgtUpload` 都只传哈希 ID + 文件名，具体解析（BundleId/版本/图标/manifest）全由后端做。前端不要拼 URL、也不要冗余 `md5`。
2. **同版本仅一个已发布 wgt**：`fabuWgtRelease` 后端会自动把同 `version_id` 下其它 `released=1` 记录取消；前端 `wgt.vue` 在确认框文案里就明确写了这条规则。
3. **`update_mode=2` 静默已废弃**：SDK 不提供、也不识别，一律按 `update_mode=0` 处理。
4. **`version` 比较用 name 分段**：wgt 协议已废弃"去点 code"，客户端上报 `uni.getAppBaseInfo().appWgtVersion`；`1.10` 与 `1.1.0` 不再歧义。
5. **检查更新是公开接口**：`/api/web/fabu/version/checkupdate` 与 `/api/web/fabu/count/*` 都不带 Token；不要在 SDK 里塞鉴权头。
6. **`download_url` 可能是相对路径**（`/api/uploads/<id>/<name>`）：SDK 内 `absUrl` 会统一补 `baseUrl`；管理端下载链接复制时也手动拼 `window.location.origin + row.download_url`。
7. **首次进 App 未渲染 `<x-upgrade />` 就调 `check()`**：靠 `pendingCheck` 兜住，组件挂载时补执行；如果整个 App 都不放该组件，永远不会弹窗——务必在首页或常驻页挂上。
8. **wgt 归属当前 Version**：`fabuWgtUpload` 必须带 `version_id`（前端从 `wgt.vue` 的 `props.versionId` 传入，弹窗由"热更新"按钮打开时携带行 ID）。
9. **iOS 检查接口 `version_code` 与 Android 可能同值不同义**：服务端用 `bundle_id + platform` 双维度定位，客户端上报时 `platform` 必填。

## 七、相关代码位置

- 管理端 API 与页面：
  - `admin/src/api/fabu.ts`
  - `admin/src/api/fabu_download.ts`（落地页拉取应用信息，`fabuAppDownloadInfo`）
  - `admin/src/views/fabu/{app,version,download}/index.vue`
  - `admin/src/views/fabu/version/wgt.vue`
  - `admin/src/views/fabu/components/upload.vue`（承载 `upload-chunk`）
- App 端 SDK：
  - `x_admin_app/sdk/upgrade/README.md`（协议、流程、扩展指南）
  - `x_admin_app/sdk/upgrade/vue3/{config.ts,flow.ts,types.ts}`
  - `x_admin_app/sdk/upgrade/uniappx/`
  - 组件：`components/x-upgrade/x-upgrade.{vue,uvue}`
- 服务端接口：`server/app/controller/admin_ctl/fabu_controller/`、`server/app/controller/web_ctl/fabu_controller.go`（检查更新、计数下载、`short_url` 落地页信息）、`server/app/service/fabu/`。
- 上传通道：见 [分片上传](./分片上传.md)。
