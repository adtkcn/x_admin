<template>
    <div ref="screenRef" class="bigscreen">
        <!-- 顶部：标题 + 统计 + 状态/时钟/全屏 -->
        <header class="bs__bar">
            <!-- <div class="bs__title">区域物联大屏</div> -->
            <div class="bs__stats">
                <div class="bs__stat">
                    <span class="num">{{ points.length }}</span>
                    <span class="label">设备</span>
                </div>
                <div class="bs__stat">
                    <span class="num online">{{ stats.online }}</span>
                    <span class="label">在线</span>
                </div>
                <div class="bs__stat">
                    <span class="num offline">{{ stats.offline }}</span>
                    <span class="label">离线</span>
                </div>
                <div class="bs__stat">
                    <span class="num alerting">{{ stats.alerting }}</span>
                    <span class="label">告警中</span>
                </div>
            </div>
            <div class="bs__right">
                <el-tag size="small" :type="wsStatus === 'open' ? 'success' : 'danger'" effect="dark">
                    {{ wsStatus === 'open' ? '实时推送已连接' : '实时推送未连接' }}
                </el-tag>
                <span class="bs__time">{{ timeText }}</span>
                <el-button size="small" type="primary" plain @click="toggleFullscreen">
                    {{ isFullscreen ? '退出全屏' : '全屏投屏' }}
                </el-button>
            </div>
        </header>

        <div class="bs__body">
            <!-- 左侧：区域树 -->
            <aside class="bs__side">
                <div class="bs__panel-head">
                    <span>区域</span>
                    <el-button link type="primary" size="small" @click="selectSpace('')">全部</el-button>
                </div>
                <el-scrollbar class="bs__tree">
                    <el-tree
                        :data="spaceTree"
                        node-key="id"
                        :props="{ label: 'name' }"
                        highlight-current
                        :expand-on-click-node="false"
                        @node-click="handleTreeClick"
                    />
                </el-scrollbar>
                        </aside>

            <!-- 右侧：地图 -->
            <div class="bs__map-wrap">
                <div ref="mapContainer" class="bs__map"></div>
                <div class="bs__legend">
                    <span class="lg"><i class="dot dot--online"></i>在线</span>
                    <span class="lg"><i class="dot dot--offline"></i>离线</span>
                    <span v-for="lv in SeverityLevels" :key="lv" class="lg">
                        <i class="dot dot--lv" :style="{ background: SeverityColor[lv] }"></i>
                        {{ SeverityMap[lv] }} <b>{{ stats.byLevel[lv] || 0 }}</b>
                    </span>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, computed, watch, onMounted, onBeforeUnmount, nextTick } from 'vue'
import 'maplibre-gl/dist/maplibre-gl.css'
import {
    deviceSpaceAll,
    deviceMap,
    deviceState,
    SeverityLevels,
    SeverityMap,
    SeverityColor,
    type type_iot_device_space,
    type type_device_map_item
} from '@/api/iot'
import { arrayToTree } from '@/utils/util'
import { useGlobalWs, onWsMessage, type WsEvent } from '@/hooks/useGlobalWs'
import { useIotMap } from './useIotMap'

defineOptions({ name: 'iotMonitor' })

const { status: wsStatus, send: wsSend } = useGlobalWs()

const spaces = ref<type_iot_device_space[]>([])
const spaceTree = ref<any[]>([])
const points = ref<type_device_map_item[]>([])
const currentSpace = ref('')

// 设备最新值缓存：telemetry 推送实时写入，popup 展示用
const latestValues = reactive<Record<string, Record<string, any>>>({})

// 区域子树 id（空间模型是 pid 邻接表，子树在前端收集）
const subtreeIds = (rootId: string): string[] => {
    if (!rootId) return []
    const childrenOf = new Map<string, string[]>()
    for (const s of spaces.value) {
        const pid = s.pid || ''
        childrenOf.set(pid, [...(childrenOf.get(pid) || []), s.id!])
    }
    const out: string[] = []
    const walk = (id: string) => {
        out.push(id)
        for (const c of childrenOf.get(id) || []) walk(c)
    }
    walk(rootId)
    return out
}

const stats = computed(() => {
    let online = 0
    let alerting = 0
    // 级别计数按级别值归档：有负数档（故障=-1），不能再用数组下标
    const byLevel: Record<number, number> = {}
    for (const p of points.value) {
        if (p.status === 1) online++
        if (p.severity !== 0) {
            alerting++
            byLevel[p.severity] = (byLevel[p.severity] || 0) + 1
        }
    }
    return { online, offline: points.value.length - online, byLevel, alerting }
})

// ---------------- 地图（MapLibre 实例/图层/popup/视野等细节全封装在组合式函数里） ----------------
const mapContainer = ref<HTMLElement>()
const map = useIotMap({
    container: mapContainer,
    points,
    spaces,
    currentSpace,
    getPopupValues: (id) => latestValues[id],
    onSelectSpace: (id) => selectSpace(id),
    onPopupOpen: (id) => ensureValues(id)
})

// ---------------- 数据加载 / 区域选择 ----------------
const loadPoints = async () => {
    const ids = subtreeIds(currentSpace.value)
    points.value = await deviceMap({ space_ids: ids.join(',') })
    map.refreshAll()
    // 全部视图（未选区域，ids 为空）适配到所有带轮廓的区域范围，避免停留在初始 zoom
    map.fitToSubtree(ids.length ? ids : spaces.value.map((s) => s.id!).filter(Boolean))
}

const selectSpace = async (id: string) => {
    if (id === currentSpace.value) return
    currentSpace.value = id
    map.closePopup()
    await loadPoints()
}

const handleTreeClick = (node: any) => selectSpace(node.id)

// popup 的实时值仅靠 iot:telemetry 推送填充，而该推送需规则链配了「状态快照」节点标脏才有；
// 故点开设备且无存量值时按设备回读一次 Redis 快照兜底（与详情页实时 tab 同源），已有推送值则不覆盖
const ensureValues = async (id: string) => {
    const cached = latestValues[id]
    if (cached && Object.keys(cached).length) return
    try {
        const st = await deviceState({ id })
        const fresh = latestValues[id]
        if (st?.values && Object.keys(st.values).length && (!fresh || !Object.keys(fresh).length)) {
            latestValues[id] = st.values
            map.syncPopup(id)
        }
    } catch {
        /* 回读失败保持原渲染 */
    }
}

// ---------------- WS 房间 ----------------
// 后端按报文设备所在区域广播到 iot:map:{spaceId}；进入页面把所有区域房间都加入，
// 这样“全部/任意区域”视图都有实时增量，无需按选择反复进出房间
const joinAllRooms = () => {
    for (const s of spaces.value) {
        if (s.id) wsSend(JSON.stringify({ type: 'join_room', data: `iot:map:${s.id}` }))
    }
}

// 房间成员关系保存在服务端连接上：首次连接与断线自动重连后都要重新 join
watch(wsStatus, (s) => {
    if (s === 'open') {
        joinAllRooms()
        loadPoints()
    }
})

onWsMessage('iot:status', (msg: WsEvent) => {
    const p = map.getPoint(msg.data?.device_id)
    if (!p) return
    p.status = msg.data.status
    // 设备转离线时后端已收尾活动告警，但不推级别，这里同步清掉，避免离线点仍挂告警色
    if (msg.data.status !== 1) p.severity = 0
    map.markDirty()
    map.syncPopup(p.id)
})

onWsMessage('iot:telemetry', (msg: WsEvent) => {
    if (msg.data?.device_id && msg.data?.values) {
        latestValues[msg.data.device_id] = msg.data.values
        map.syncPopup(msg.data.device_id)
    }
})

onWsMessage('iot:alarm', (msg: WsEvent) => {
    const p = map.getPoint(msg.data?.device_id)
    if (!p) return
    // status=0 为恢复(清级别)；1 为触发/升级，取推送里的 severity（含合法的 -1），缺失按轻微处理不能丢标记
    p.severity = msg.data?.status === 1 ? Number(msg.data?.severity) || 1 : 0
    map.markDirty()
    latestValues[msg.data.device_id] = msg.data.values
    map.syncPopup(p.id)
})

// ---------------- 全屏 / 时钟 ----------------
const screenRef = ref<HTMLElement>()
const isFullscreen = ref(false)

const toggleFullscreen = () => {
    if (document.fullscreenElement) {
        document.exitFullscreen()
    } else {
        screenRef.value?.requestFullscreen()
    }
}

const onFullscreenChange = () => {
    isFullscreen.value = !!document.fullscreenElement
    // 容器尺寸变化后让地图重算视口
    nextTick(() => map.resize())
}

const timeText = ref('')
const pad = (n: number) => String(n).padStart(2, '0')
const tick = () => {
    const d = new Date()
    timeText.value = `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`
}
let clockTimer: ReturnType<typeof setInterval> | null = null

// ---------------- 生命周期 ----------------
onMounted(async () => {
    spaces.value = await deviceSpaceAll()
    spaceTree.value = arrayToTree<any>(spaces.value as any[], '')
    await nextTick()
    map.init()
    await loadPoints()
    // WS 可能早已连接（layout 初始化），此时 watch 不会触发，需主动进房
    if (wsStatus.value === 'open') joinAllRooms()
    tick()
    clockTimer = setInterval(tick, 1000)
    document.addEventListener('fullscreenchange', onFullscreenChange)
})

onBeforeUnmount(() => {
    if (clockTimer) clearInterval(clockTimer)
    document.removeEventListener('fullscreenchange', onFullscreenChange)
    map.destroy()
})
</script>

<style scoped>
.bigscreen {
    display: flex;
    flex-direction: column;
    height: calc(100vh - 130px);
    border-radius: 8px;
    overflow: hidden;
    background: linear-gradient(160deg, #0b1c36 0%, #123058 60%, #0b1c36 100%);
    color: #d6e6ff;
}
.bigscreen:fullscreen {
    height: 100vh;
    padding: 12px 16px;
}
.bs__bar {
    display: flex;
    align-items: center;
    gap: 24px;
    padding: 10px 16px;
    flex: none;
    border-bottom: 1px solid rgba(64, 158, 255, 0.25);
}
.bs__title {
    font-size: 20px;
    font-weight: 700;
    letter-spacing: 2px;
    color: #eaf4ff;
    text-shadow: 0 0 12px rgba(64, 158, 255, 0.6);
}
.bs__stats {
    display: flex;
    gap: 28px;
}
.bs__stat {
    display: flex;
    align-items: baseline;
    gap: 6px;
}
.bs__stat .num {
    font-size: 22px;
    font-weight: 700;
    color: #5eaaff;
}
.bs__stat .num.online {
    color: #67c23a;
}
.bs__stat .num.offline {
    color: #9fb3d1;
}
.bs__stat .num.alerting {
    color: #f56c6c;
}
.bs__stat .label {
    font-size: 12px;
    color: #8fa9cc;
}
.bs__right {
    margin-left: auto;
    display: flex;
    align-items: center;
    gap: 12px;
}
.bs__time {
    font-family: Consolas, Monaco, monospace;
    font-size: 14px;
    color: #a9c4e8;
}
.bs__body {
    flex: 1;
    display: flex;
    gap: 12px;
    padding: 12px;
    min-height: 0;
}
.bs__side {
    width: 250px;
    flex: none;
    display: flex;
    flex-direction: column;
    border: 1px solid rgba(64, 158, 255, 0.25);
    border-radius: 8px;
    background: rgba(9, 25, 48, 0.55);
}
.bs__panel-head {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 8px 12px;
    font-weight: 600;
    border-bottom: 1px solid rgba(64, 158, 255, 0.2);
}
.bs__tree {
    flex: 1;
    min-height: 0;
}
 
.bs__map-wrap {
    position: relative;
    flex: 1;
    min-width: 0;
    border-radius: 8px;
    overflow: hidden;
}
.bs__map {
    position: absolute;
    inset: 0;
}
.bs__legend {
    position: absolute;
    left: 12px;
    bottom: 12px;
    z-index: 2;
    display: flex;
    gap: 14px;
    padding: 6px 12px;
    border-radius: 6px;
    border: 1px solid rgba(64, 158, 255, 0.35);
    background: rgba(9, 25, 48, 0.72);
    backdrop-filter: blur(2px);
    font-size: 12px;
    color: #cfe3ff;
    pointer-events: none;
}
.bs__legend .lg {
    display: inline-flex;
    align-items: center;
    gap: 6px;
}
.bs__legend .dot {
    width: 10px;
    height: 10px;
    border-radius: 50%;
    box-shadow: 0 0 0 1.5px #fff inset;
}
.bs__legend .dot--online {
    background: #22c55e;
}
.bs__legend .dot--offline {
    background: #64748b;
    opacity: 0.6;
}
/* 告警级别色由 SeverityColor 内联给定，这里只补与地图光晕一致的描边 */
.bs__legend .dot--lv {
    box-shadow: 0 0 0 1.5px rgba(255, 255, 255, 0.85) inset, 0 0 8px rgba(255, 255, 255, 0.25);
}
.bs__legend b {
    font-weight: 600;
    color: #eaf4ff;
}
:deep(.bs__tree .el-tree) {
    background: transparent;
    color: #cfe3ff;
    --el-tree-node-hover-bg-color: rgba(64, 158, 255, 0.15);
}
:deep(.bs__tree .el-tree-node.is-current > .el-tree-node__content) {
    background: rgba(64, 158, 255, 0.28);
}
:deep(.map-pop) {
    background: rgba(10, 28, 54, 0.94);
    border: 1px solid rgba(64, 158, 255, 0.4);
    color: #d6e6ff;
    border-radius: 6px;
    padding: 8px 10px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.45);
    max-width: 260px;
    font-size: 13px;
    line-height: 1.7;
}
</style>
