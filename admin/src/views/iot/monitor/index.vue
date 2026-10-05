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
                <div class="bs__hint">点地图区域轮廓或树节点筛选其子树设备，推送房间覆盖全部区域</div>
            </aside>

            <!-- 右侧：地图 -->
            <div class="bs__map-wrap">
                <div ref="mapContainer" class="bs__map"></div>
                <div class="bs__legend">
                    <span class="lg"><i class="dot dot--online"></i>在线</span>
                    <span class="lg"><i class="dot dot--offline"></i>离线</span>
                    <span class="lg"><i class="dot dot--alert"></i>告警中</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, computed, shallowRef, watch, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { Map as MapLibreMap, NavigationControl, Marker } from 'maplibre-gl'
import type { StyleSpecification } from 'maplibre-gl'
import 'maplibre-gl/dist/maplibre-gl.css'
import {
    deviceSpaceAll,
    deviceMap,
    type type_iot_device_space,
    type type_device_map_item
} from '@/api/iot'
import { arrayToTree } from '@/utils/util'
import { useGlobalWs, onWsMessage, type WsEvent } from '@/hooks/useGlobalWs'

defineOptions({ name: 'iotMonitor' })

const { status: wsStatus, send: wsSend } = useGlobalWs()

const spaces = ref<type_iot_device_space[]>([])
const spaceTree = ref<any[]>([])
const points = ref<type_device_map_item[]>([])
const currentSpace = ref('')

// 设备索引：WS 报文按 device_id 直接命中点位，避免每条报文 O(n) 查找
const pointIndex = new Map<string, type_device_map_item>()
const rebuildIndex = () => {
    pointIndex.clear()
    for (const p of points.value) pointIndex.set(p.id!, p)
}

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
    for (const p of points.value) {
        if (p.status === 1) online++
        if (p.alerting) alerting++
    }
    return { online, offline: points.value.length - online, alerting }
})

// ---------------- 地图 ----------------
const mapContainer = ref<HTMLElement>()
const map = shallowRef<MapLibreMap>()
const popupMarker = shallowRef<Marker>()

// 样式解析与 map-picker 同款：环境变量瓦片源优先，缺省 OSM
const resolveStyle = (): StyleSpecification | string => {
    const tiles = import.meta.env.VITE_MAP_TILES_URL as string | undefined
    if (tiles) {
        return {
            version: 8,
            sources: {
                custom: { type: 'raster', tiles: [tiles], tileSize: 256, attribution: '' }
            },
            layers: [{ id: 'custom-raster', type: 'raster', source: 'custom' }]
        }
    }
    return 'https://demotiles.maplibre.org/style.json'
}

const areasFeatureCollection = () => {
    const features: any[] = []
    for (const s of spaces.value) {
        if (!s.geojson) continue
        try {
            const g = JSON.parse(s.geojson)
            if (g?.type === 'Feature') {
                g.properties = { ...(g.properties || {}), name: s.name, space_id: s.id }
                features.push(g)
            }
        } catch {
            /* 脏数据跳过，不影响铺图 */
        }
    }
    return { type: 'FeatureCollection' as const, features }
}

// 多边形外环面积（shoelace 绝对值，仅用于同一命中点内的相对大小比较）
const ringArea = (geom: any): number => {
    const ring =
        geom?.type === 'Polygon'
            ? geom.coordinates?.[0]
            : geom?.type === 'MultiPolygon'
              ? geom.coordinates?.[0]?.[0]
              : null
    if (!Array.isArray(ring)) return Number.MAX_SAFE_INTEGER
    let sum = 0
    for (let i = 0; i < ring.length - 1; i++) {
        const [x1, y1] = ring[i]
        const [x2, y2] = ring[i + 1]
        sum += x1 * y2 - x2 * y1
    }
    return Math.abs(sum / 2)
}

const devicesFeatureCollection = () => ({
    type: 'FeatureCollection' as const,
    features: points.value
        .filter((p) => p.longitude !== 0 || p.latitude !== 0)
        .map((p) => ({
            type: 'Feature' as const,
            geometry: { type: 'Point' as const, coordinates: [p.longitude, p.latitude] },
            // status/alerting 必须带进 properties，图层 paint 靠它们区分在线/离线/告警配色
            properties: { id: p.id, status: p.status ?? 0, alerting: !!p.alerting }
        }))
})

const refreshDevices = () => {
    const src = map.value?.getSource('devices') as any
    if (src) src.setData(devicesFeatureCollection())
}

// WS 增量合并渲染：300ms 窗口内的多条报文只触发一次 setData
let flushTimer: ReturnType<typeof setTimeout> | null = null
const markDirty = () => {
    if (flushTimer) return
    flushTimer = setTimeout(() => {
        flushTimer = null
        refreshDevices()
    }, 300)
}

const initMap = () => {
    if (!mapContainer.value || !map.value) return
    map.value.on('load', () => {
        // 区域轮廓：透明面用于点选，虚线描边用于展示
        map.value!.addSource('areas', { type: 'geojson', data: areasFeatureCollection() })
        map.value!.addLayer({
            id: 'areas-fill',
            type: 'fill',
            source: 'areas',
            paint: { 'fill-color': '#409eff', 'fill-opacity': 0.06 }
        } as any)
        map.value!.addLayer({
            id: 'areas-line',
            type: 'line',
            source: 'areas',
            paint: {
                'line-color': '#409eff',
                'line-width': 2,
                'line-dasharray': [2, 2]
            }
        } as any)
        // 设备点位：告警红(带光晕) / 在线绿 / 离线灰(淡化)，alerting 优先级最高
        map.value!.addSource('devices', { type: 'geojson', data: devicesFeatureCollection() })
        // 告警光晕：仅告警设备，红色半透明大圆制造聚焦感
        map.value!.addLayer({
            id: 'devices-halo',
            type: 'circle',
            source: 'devices',
            filter: ['==', ['get', 'alerting'], true],
            paint: {
                'circle-radius': 15,
                'circle-color': '#f56c6c',
                'circle-opacity': 0.16,
                'circle-blur': 0.7
            }
        } as any)
        map.value!.addLayer({
            id: 'devices-circle',
            type: 'circle',
            source: 'devices',
            paint: {
                'circle-radius': ['case', ['get', 'alerting'], 9, ['==', ['get', 'status'], 1], 7, 6],
                'circle-stroke-width': ['case', ['get', 'alerting'], 2.5, 1.5],
                'circle-stroke-color': ['case', ['get', 'alerting'], '#ffd9d9', '#fff'],
                'circle-color': [
                    'case',
                    ['get', 'alerting'],
                    '#f56c6c',
                    ['==', ['get', 'status'], 1],
                    '#22c55e',
                    '#64748b'
                ],
                // 离线点位略作淡化，在线/告警保持醒目
                'circle-opacity': [
                    'case',
                    ['get', 'alerting'],
                    1,
                    ['==', ['get', 'status'], 1],
                    1,
                    0.6
                ]
            }
        } as any)
        map.value!.on('click', (e: any) => {
            const hits = map.value!.queryRenderedFeatures(e.point, {
                layers: ['devices-circle', 'areas-fill']
            })
            const dev = hits.find((f: any) => f.layer.id === 'devices-circle')
            if (dev) {
                openPopupById(String(dev.properties?.id), e.lngLat)
                return
            }
            // 重叠区域取面积最小者（最具体的子区域），避免父区域面盖住子区域被误选中放大到父区域
            const areaHits = hits.filter((f: any) => f.layer.id === 'areas-fill')
            if (areaHits.length) {
                const area = areaHits.reduce((a: any, b: any) =>
                    ringArea(b.geometry) < ringArea(a.geometry) ? b : a
                )
                selectSpace(String(area.properties?.space_id || ''))
                return
            }
            closePopup()
        })
        for (const layer of ['devices-circle', 'devices-halo', 'areas-fill']) {
            map.value!.on('mouseenter', layer, () => {
                map.value!.getCanvas().style.cursor = 'pointer'
            })
            map.value!.on('mouseleave', layer, () => {
                map.value!.getCanvas().style.cursor = ''
            })
        }
    })
}

// 选中子树后把视野适配到其区域轮廓范围
const fitToSubtree = (ids: string[]) => {
    if (!map.value || !ids.length) return
    let minX = Infinity,
        minY = Infinity,
        maxX = -Infinity,
        maxY = -Infinity
    for (const s of spaces.value) {
        if (!s.id || !ids.includes(s.id) || !s.geojson) continue
        try {
            const g = JSON.parse(s.geojson)
            if (g?.geometry?.type !== 'Polygon') continue
            for (const [lng, lat] of g.geometry.coordinates[0] as number[][]) {
                if (lng < minX) minX = lng
                if (lng > maxX) maxX = lng
                if (lat < minY) minY = lat
                if (lat > maxY) maxY = lat
            }
        } catch {
            /* 脏数据跳过 */
        }
    }
    if (minX < Infinity) {
        map.value.fitBounds(
            [
                [minX, minY],
                [maxX, maxY]
            ],
            { padding: 80, duration: 600 }
        )
    }
}

// ---------------- popup ----------------
// 设备最新值缓存：telemetry 推送实时写入，popup 展示用
const latestValues = reactive<Record<string, Record<string, any>>>({})
// 当前 popup 上下文：设备点位数据变化时按此重渲内容
const popupCtx = shallowRef<{ id: string; lngLat: { lng: number; lat: number } } | null>(null)

// 全部用 textContent 构建，避免设备名/SN/属性值注入 HTML
const addLine = (box: HTMLElement, text: string, color?: string, bold?: boolean) => {
    const d = document.createElement('div')
    d.textContent = text
    if (color) d.style.color = color
    if (bold) d.style.fontWeight = '600'
    box.appendChild(d)
}

const renderPopup = () => {
    if (!map.value || !popupCtx.value) return
    const p = pointIndex.get(popupCtx.value.id)
    if (!p) {
        closePopup()
        return
    }
    const box = document.createElement('div')
    box.className = 'map-pop'
    addLine(box, p.name || p.sn || p.id!, undefined, true)
    addLine(box, `SN：${p.sn || '-'}`)
    addLine(box, `状态：${p.status === 1 ? '在线' : '离线'}`)
    if (p.alerting) addLine(box, '告警中', '#f56c6c')
    const vals = latestValues[p.id!]
    if (vals && Object.keys(vals).length) {
        for (const [k, v] of Object.entries(vals)) addLine(box, `${k} = ${String(v)}`)
    } else {
        addLine(box, '暂无实时值')
    }
    if (!popupMarker.value) {
        popupMarker.value = new Marker({ element: document.createElement('div') })
    }
    popupMarker.value.getElement().replaceChildren(box)
    popupMarker.value.setLngLat(popupCtx.value.lngLat).addTo(map.value)
}

const openPopupById = (id: string, lngLat: { lng: number; lat: number }) => {
    if (!pointIndex.has(id)) return
    popupCtx.value = { id, lngLat }
    renderPopup()
}

const closePopup = () => {
    popupCtx.value = null
    popupMarker.value?.remove()
}

// ---------------- 数据加载 / 区域选择 ----------------
const loadPoints = async () => {
    const ids = subtreeIds(currentSpace.value)
    points.value = await deviceMap({ space_ids: ids.join(',') })
    rebuildIndex()
    refreshDevices()
    // 全部视图（未选区域，ids 为空）适配到所有带轮廓的区域范围，避免停留在初始 zoom
    fitToSubtree(ids.length ? ids : spaces.value.map((s) => s.id!).filter(Boolean))
}

const selectSpace = async (id: string) => {
    if (id === currentSpace.value) return
    currentSpace.value = id
    closePopup()
    await loadPoints()
}

const handleTreeClick = (node: any) => selectSpace(node.id)

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
    const p = pointIndex.get(msg.data?.device_id)
    if (!p) return
    p.status = msg.data.status
    markDirty()
    if (popupCtx.value?.id === p.id) renderPopup()
})

onWsMessage('iot:telemetry', (msg: WsEvent) => {
    if (msg.data?.device_id && msg.data?.values) {
        latestValues[msg.data.device_id] = msg.data.values
        if (popupCtx.value?.id === msg.data.device_id) renderPopup()
    }
})

onWsMessage('iot:alarm', (msg: WsEvent) => {
    const p = pointIndex.get(msg.data?.device_id)
    if (!p) return
    p.alerting = msg.data?.status === 1
    markDirty()
    if (popupCtx.value?.id === p.id) renderPopup()
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
    nextTick(() => map.value?.resize())
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
    map.value = new MapLibreMap({
        container: mapContainer.value!,
        style: resolveStyle(),
        center: [104.0668, 30.5728],
        zoom: 4.5
    })
    map.value.addControl(new NavigationControl({ showCompass: false }), 'top-right')
    initMap()
    await loadPoints()
    // WS 可能早已连接（layout 初始化），此时 watch 不会触发，需主动进房
    if (wsStatus.value === 'open') joinAllRooms()
    tick()
    clockTimer = setInterval(tick, 1000)
    document.addEventListener('fullscreenchange', onFullscreenChange)
})

onBeforeUnmount(() => {
    if (flushTimer) clearTimeout(flushTimer)
    if (clockTimer) clearInterval(clockTimer)
    document.removeEventListener('fullscreenchange', onFullscreenChange)
    closePopup()
    popupMarker.value = undefined
    map.value?.remove()
    map.value = undefined
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
.bs__hint {
    flex: none;
    padding: 8px 12px;
    font-size: 12px;
    line-height: 1.5;
    color: #7d97bd;
    border-top: 1px solid rgba(64, 158, 255, 0.2);
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
.bs__legend .dot--alert {
    background: #f56c6c;
    box-shadow: 0 0 0 1.5px #ffd9d9 inset, 0 0 8px rgba(245, 108, 108, 0.8);
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
