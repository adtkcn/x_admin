# WebSocket 使用说明

管理端所有实时消息（通知、在线人数、IoT 状态/遥测/告警等）都走**同一条全局 WebSocket 连接**，由 `admin/src/hooks/useGlobalWs.ts` 以「单例 + 事件总线」的形式封装。业务组件**不要自己 `new WebSocket`**，一律用 `onWsMessage` 订阅即可。

## 一、三个入口一览

| 入口 | 用途 | 调用位置 |
| --- | --- | --- |
| `initGlobalWs()` | 建立全局唯一 WS 连接（登录后调用，重复调用被短路） | 只在 `layout/default/index.vue` 的 `onMounted` 里调一次 |
| `destroyGlobalWs()` | 关闭连接并重置状态 | `layout/default/index.vue` 的 `onUnmounted` |
| `useGlobalWs()` | 拿到 `status` / `send` / `close` | 任意组件（需要主动发消息或展示连接状态时） |
| `onWsMessage(type, handler)` | 订阅指定 `type` 的消息，组件卸载时自动取消 | 任意组件 |

## 二、连接建立细节

- **URL**：`${wss|ws}://${window.location.host}/api/ws?token=${userStore.token}`（`https` 页面自动走 `wss`）。
- **鉴权**：`token` 走 query，由后端 `app/controller/ws.go` 升级处理器校验；无 token 时 `initGlobalWs` 直接短路，不发起连接。
- **心跳**：每 **30 秒** 发一次文本 `"ping"`，`pongTimeout=1000ms`；`onMessage` 里对 `data === 'pong'` 直接过滤，不会推给业务。
- **重连**：`autoReconnect: { retries: -1, delay: 3000 }`——无限次重连，间隔 3 秒。
- **状态**：`'connecting' | 'open' | 'closed'`，通过 `useGlobalWs().status` 拿到（`readonly`），可在 UI 上显示连接指示灯。

## 三、消息协议

前后端约定统一 JSON 帧：

- **服务端推送**（`core/ws.WsResponse`）：`{ "type": "<业务类型>", "data": <任意> }`。收到后 `wsBus.emit(msg)`，`onWsMessage` 按 `msg.type` 分发。
- **客户端发送**（`app/controller/ws.go` 的 `clientMessage`）：`{ "type": "<指令>", "data": "<字符串>" }`。目前内置指令：
  - `join_room` — 加入房间，`data` 为房间 ID
  - `leave_room` — 离开房间

其它自定义 `type`（如控制指令）目前会被后端忽略并打日志，需要时在后端 `handleWsMessage` 的 `switch` 里扩展。

## 四、典型用法

### 1）订阅消息（组件卸载自动取消订阅）

```ts
// 通知角标：layout/default/components/header/notice.vue
import { onWsMessage } from '@/hooks/useGlobalWs'

onWsMessage('notice', (msg) => {
    fetchUnreadCount()
    // msg.data 携带本次通知的载荷
})
```

```ts
// IoT 大屏：views/iot/monitor/index.vue —— 三种事件分别订阅
import { onWsMessage, useGlobalWs, type WsEvent } from '@/hooks/useGlobalWs'

onWsMessage('iot:status',    (msg: WsEvent) => { /* 上下线 */ })
onWsMessage('iot:telemetry', (msg: WsEvent) => { /* 属性值刷新 */ })
onWsMessage('iot:alarm',     (msg: WsEvent) => { /* 告警分级图标 */ })
```

订阅所有类型（调试用）：`onWsMessage('*', (msg) => console.log(msg))`。

### 2）加入 / 退出房间（按需订阅）

后端按房间推送的场景（比如工作台在线人数），组件挂/卸时显式 `join_room` / `leave_room`，避免不相关页面也收到推送：

```ts
// views/workbench/index.vue
const ONLINE_ROOM = 'onlineCount'
const { send, status: wsStatus } = useGlobalWs()

const joinOnlineRoom  = () => send(JSON.stringify({ type: 'join_room',  data: ONLINE_ROOM }))
const leaveOnlineRoom = () => send(JSON.stringify({ type: 'leave_room', data: ONLINE_ROOM }))

// 关键：断线重连后房间成员关系会丢，靠 status 变 open 时补一次 join
watch(wsStatus, (s) => { if (s === 'open') joinOnlineRoom() })

onMounted(joinOnlineRoom)
onActivated(joinOnlineRoom)     // keep-alive 命中时补 join
onDeactivated(leaveOnlineRoom)
onUnmounted(leaveOnlineRoom)

onWsMessage('onlineCount', (msg) => updateChart(msg.data.count))
```

> 房间加入是**连接级**状态：连接断开即离开房间。所以只要可能发生重连，就一定要 `watch(wsStatus)` 在 `open` 时补一次 `join_room`，别只在 `onMounted` 里发一次。

### 3）显示连接状态

```vue
<script setup lang="ts">
import { useGlobalWs } from '@/hooks/useGlobalWs'
const { status } = useGlobalWs()
</script>

<template>
    <el-tag :type="status === 'open' ? 'success' : 'danger'">{{ status }}</el-tag>
</template>
```

## 五、生命周期约定

- **只在 layout 建立/销毁连接**：`layout/default/index.vue` 中 `onMounted → initGlobalWs()`、`onUnmounted → destroyGlobalWs()`。业务页面**禁止**再调 `initGlobalWs` / `destroyGlobalWs`。
- **登录跳转**：登录后进入 layout 会自然建连；退出登录会走 layout 卸载路径断开。
- **组件订阅**：一律用 `onWsMessage`，内部已 `onUnmounted(() => unsubscribe())`；不要在 `setup` 之外调用（否则 `onUnmounted` 无法注册，会内存泄漏）。

## 六、常见坑

1. **重复建连**：`initGlobalWs()` 内有 `if (wsInstance) return`，多实例并存不会成环，但会浪费连接——坚持只在 layout 调用一次。
2. **忘记加入房间**：后端有些推送是按房间（如 `onlineCount`）而不是全局广播，组件不 `join_room` 就永远收不到；排查时先看后端 `SendToRoom(...)` 的房间名。
3. **在 `setup` 外调用 `onWsMessage`**：`onUnmounted` 需在组件生命周期内注册，写在 `setTimeout` / 异步回调里不会自动清理。
4. **心跳 `pong` 泄漏**：`onMessage` 里已过滤 `e.data === 'pong'`；若后端改成对象帧承载心跳，记得同步过滤条件。
5. **反代丢失 Upgrade**：nginx 需为 `/api/ws` 加 `Upgrade`/`Connection` 头，否则浏览器只能拿到 HTTP 400。参见部署章节 nginx 配置。

## 七、相关代码位置

- 前端封装：`admin/src/hooks/useGlobalWs.ts`
- 生命周期挂载：`admin/src/layout/default/index.vue`
- 典型消费点：`admin/src/layout/default/components/header/notice.vue`、`admin/src/views/workbench/index.vue`、`admin/src/views/iot/monitor/index.vue`
- 后端连接与消息路由：`server/app/controller/ws.go`（升级 + `join_room/leave_room` 分发）
- 后端推送 API：`server/core/ws/manager.go`（`SendToUser/SendToRoom/SendToAll`）与消息结构 `WsResponse`
