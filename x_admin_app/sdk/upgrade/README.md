# x_admin App 升级 SDK

对接后端 `fabu` 应用发布模块的版本检查/下载/安装 SDK。当前提供两个端的实现，
并预留多端扩展位（如 react-native）。

```
sdk/upgrade/
├── README.md          # 本文件：协议、流程、架构、扩展指南
├── vue3/              # uni-app vue3 端（js/ts）
│   ├── types.ts       # 协议类型（平台无关，新端可复用定义）
│   ├── config.ts      # 配置 + 共享状态 + check() 入口
│   └── flow.ts        # 决策 + 桥接（检查/判定/下载/安装，plus.* 集中于此）
├── uniappx/           # uni-app x 端（uts）—— 组件源见下方说明
└── react-native/      # 预留，未实现
```

配套 UI 组件按「组件即入口」方式放在项目 components 目录（HBuilderX 项目不便跨目录
import easycom 组件，故两端组件同名放置，按目标工程类型只编译各自后缀文件）：

| 端 | 组件 | 说明 |
| --- | --- | --- |
| uni-app vue3 | `components/x-upgrade/x-upgrade.vue` | 弹窗 + 进度条，全功能（含 wgt 静默） |
| uni-app x | `components/x-upgrade/x-upgrade.uvue` | 自包含组件（无 wgt，仅 android 安装） |

> 在 uni-app x 工程中使用：将 `x-upgrade.uvue` 拷入其 `components/x-upgrade/` 目录即可
> （uni-app x 工程不会编译 `.vue` 后缀文件，两文件可共存于同一目录）。

## 服务端协议

```
GET {baseUrl}/api/web/fabu/version/checkupdate
    ?bundle_id=<包名>       # 应用 BundleId
    &platform=<ios|android> # 必填，安卓/iOS 包名可能相同
    &version_code=<int>     # 客户端安装包 version code
    &wgt_version=<string>   # 可选，客户端当前 wgt 资源版本 name（如 1.0.3）；缺省以宿主版本 name 为基准
```

返回（`code=0` 信封内的 `data`），三种固定形态：

```json
{ "update": true,  "type": "app", "app": { "version", "version_code", "download_url", "install_url", "update_mode" } }
{ "update": true,  "type": "wgt", "wgt": { "version", "version_code", "download_url", "md5", "size" } }
{ "update": false }
```

- `update` 只表达「有没有更新」，`type` 表达「以哪种方式更新」，数据各自成块。
- `update_mode`：`0` 普通、`1` 强制（`2` 静默已废弃，客户端一律按普通处理）。
- 服务端优先级：先判全量包（released=1 且 version_code > 客户端），无全量更新时
  才下发客户端当前版本下最新已发布 wgt。

## 三条更新流程

| 场景 | 行为 |
| --- | --- |
| 普通模式 (update_mode=0) | 弹窗提示升级 →「以后再说」可跳过（同版本 code 不再提示，落本地存储）→ 同意后后台下载，弹窗显示进度 → 下载完拉起安装 |
| 强制模式 (update_mode=1) | 弹窗无关闭入口（遮罩/按钮均不可关闭），必须等待下载并安装完成 |
| wgt 热更包 | 不弹窗：静默下载 → `plus.runtime.install` 安装 → 仅弹「重启生效」确认框（取消则下次启动再生效） |

平台差异：

- **wgt 仅 android 且仅 uni-app vue3 端支持**：uni-app x 不支持 wgt（组件收到
  `type=wgt` 直接忽略）；iOS 系统限制不允许热更自身 JS，同样跳过。
- **iOS 全量包**：vue3 端走 `plus.runtime.openURL`（`install_url` 为服务端
  itms-services OTA 地址，企业/AdHoc 分发）；uni-app x 端组件不处理 iOS，建议直连 App Store。
- **android apk 安装**：vue3 端 `plus.runtime.install`；uni-app x 端
  `uni.installApk`（HBuilderX 3.94+，需系统「安装未知应用」授权，系统会自行拉起授权页）。

## 集成

### uni-app vue3（本工程即已接好）

```ts
// App.vue onLaunch
import { check } from "@/sdk/upgrade/vue3/config";
check(); // 使用 utils/env.ts 的 baseUrl，可传 { baseUrl } 覆盖

// 常驻页面（如首页）放置弹窗组件（easycom 免导入）
// <x-upgrade></x-upgrade>
```

`check()` 可在任意时机调用（含首页加载前）：若时尚未渲染任何 `<x-upgrade />`，
SDK 内部排队，待组件 setup 注册后自动补执行；多页面重复渲染组件不会重复检查。

### uni-app x

```html
<template>
	<x-upgrade ref="upgradeRef"></x-upgrade>
</template>
<script setup lang="uts">
	const upgradeRef = ref<ComponentPublicInstance | null>(null)
	onReady(() => {
		upgradeRef.value?.$callMethod('check') // 组件自包含，无全局入口
	})
</script>
```

服务端地址在 `x-upgrade.uvue` 顶部 `DEFAULT_BASE_URL` 常量修改。

## 架构与扩展（如何新增 react-native 端）

分层原则：**协议与决策平台无关，桥接层一端一实现**。

```
types（协议类型，语义一份） → flow/决策（无更新|wgt静默|弹窗普通|弹窗强制 + 跳过记忆）
                            → 桥接（检查请求 / 下载 / 安装 / 重启 / 商店跳转）
                            → UI（弹窗组件，状态机：confirm → downloading(进度) → ready）
```

新增 react-native 端时：

1. 新建 `sdk/upgrade/react-native/`，复用 `vue3/types.ts` 的协议定义（TS 可直接 import）；
2. 决策逻辑照抄 `flow()`：`update_mode===1→force`、wgt 仅 android、跳过键
   `upgrade_skipped_version`（AsyncStorage），保证各端行为一致；
3. 桥接层换实现：检查用 `fetch`，下载用 `react-native-background-downloader` 类库，
   android 安装走 `Intent ACTION_VIEW` apk FileProvider，iOS 跳 App Store，
   RN 无 wgt 概念（收到 `type=wgt` 忽略，语义同 uni-app x）；
4. UI 按同一状态机与文案约定实现弹窗即可，服务端与后台无感知。

## 注意事项

- 检查更新为公开接口（无需 Token），SDK 内直接用 `uni.request`，不依赖项目鉴权封装。
- `download_url` 可能为相对路径（`/api/uploads/:id/:name` 文件流），SDK 内 `absUrl` 统一补全。
- wgt 比较统一用版本 **name 按段比较**（协议已废弃去点 code 方案，`1.10` 与 `1.1.0` 不再歧义）；
  客户端直接上报 `plus.runtime.version`，对 wgt 包 manifest 的 `version.code` 无格式约定。
- 后端已废弃「静默模式」（update_mode=2），本 SDK 不提供该模式。
