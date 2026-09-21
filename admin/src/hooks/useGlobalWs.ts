import { ref, readonly, onUnmounted } from 'vue'
import { useWebSocket } from '@vueuse/core'
import { useEventBus } from '@vueuse/core'
import useUserStore from '@/stores/modules/user'

/** WS 消息事件类型，与后端 WsResponse 对应 */
export type WsEvent = {
    type: string
    data: any
}

// 全局事件总线 —— WS 消息统一通过此总线分发
const wsBus = useEventBus<WsEvent>('ws')

// 连接状态
const status = ref<'connecting' | 'open' | 'closed'>('closed')

let wsInstance: ReturnType<typeof useWebSocket> | null = null

/**
 * 全局 WebSocket 连接（单例模式）
 *
 * - initGlobalWs(): 在 layout 中初始化，建立唯一 WS 连接
 * - onWsMessage(): 在组件中订阅特定类型的 WS 消息，自动清理
 * - useGlobalWs(): 获取连接状态和发送能力
 */
export function useGlobalWs() {
    return {
        status: readonly(status),
        send: (data: string | ArrayBuffer | Blob) => wsInstance?.send(data),
        close: () => wsInstance?.close()
    }
}

/**
 * 订阅 WS 消息（按类型过滤），组件卸载时自动取消订阅。
 *
 * @param type  消息类型，如 "notice"、"onlineCount"；传 "*" 订阅所有消息
 * @param handler 消息回调
 *
 * @example
 * // 订阅通知消息
 * onWsMessage('notice', (msg) => { ... })
 *
 * // 订阅在线人数（后端推送 { onlineCount: 5 }）
 * onWsMessage('onlineCount', (msg) => { updateChart(msg.onlineCount) })
 */
export function onWsMessage(type: string, handler: (msg: WsEvent) => void) {
    const unsubscribe = wsBus.on((msg) => {
        if (type === '*' || msg.type === type) {
            handler(msg)
        }
    })
    onUnmounted(() => unsubscribe())
}

/**
 * 初始化全局 WebSocket 连接（仅在 layout 中调用一次）
 */
export function initGlobalWs() {
    if (wsInstance) return

    const userStore = useUserStore()
    const token = userStore.token
    if (!token) return

    const domain = window.location.host
    const isHttps = window.location.protocol === 'https:'
    const url = `${isHttps ? 'wss' : 'ws'}://${domain}/api/ws?token=${token}`

    status.value = 'connecting'

    wsInstance = useWebSocket(url, {
        heartbeat: {
            message: 'ping',
            interval: 10000,
            pongTimeout: 1000
        },
        autoReconnect: {
            retries: -1,
            delay: 3000
        },
        onConnected() {
            status.value = 'open'
        },
        onDisconnected() {
            status.value = 'closed'
        },
        onMessage(_ws, e) {
            if (e.data === 'pong') return
            try {
                const msg = JSON.parse(e.data)
                // 后端统一返回 { type, data } 格式（WsResponse）
                if (msg.type) {
                    wsBus.emit(msg)
                }
            } catch {}
        }
    })
}

/**
 * 销毁全局 WebSocket 连接
 */
export function destroyGlobalWs() {
    if (wsInstance) {
        wsInstance.close()
        wsInstance = null
    }
    status.value = 'closed'
}
