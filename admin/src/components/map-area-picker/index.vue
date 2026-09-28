<template>
    <el-dialog
        v-model="visible"
        title="选择区域范围"
        width="80%"
        append-to-body
        draggable
        destroy-on-close
        top="5px"
        @opened="handleOpened"
        @closed="handleClosed"
    >
        <div v-loading="loading" class="map-area-picker">
            <div class="map-area-picker__toolbar">
                <span class="map-area-picker__tips"
                    >点击左上角多边形按钮开始绘制，逐点单击添加顶点，再次点击首个顶点闭结束（Esc 取消）</span
                >
                <el-button size="small" @click="handleFit">定位到区域</el-button>
                <el-button size="small" @click="handleClear">清空</el-button>
            </div>
            <div ref="mapRef" class="map-area-picker__map" />
            <!-- <div class="map-area-picker__footer">
                <el-input
                    :model-value="geojsonStr"
                    type="textarea"
                    :rows="3"
                    readonly
                    placeholder="绘制区域后自动生成 GeoJSON"
                >
                    <template #append>
                        <el-button @click="handleCopy">
                            <el-icon><icon name="el-icon-CopyDocument" /></el-icon>
                        </el-button>
                    </template>
                </el-input>
            </div> -->
        </div>
        <template #footer>
            <el-button @click="visible = false">取消</el-button>
            <el-button type="primary" @click="handleConfirm">确定</el-button>
        </template>
    </el-dialog>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Map as MapLibreMap, NavigationControl, type StyleSpecification } from 'maplibre-gl'
import MapboxDraw from '@mapbox/mapbox-gl-draw'
import 'maplibre-gl/dist/maplibre-gl.css'
import '@mapbox/mapbox-gl-draw/dist/mapbox-gl-draw.css'

const props = defineProps<{
    // v-model 绑定的 GeoJSON 字符串(Polygon Feature)
    modelValue?: string
    // 地图初始中心 [经度, 纬度]
    center?: [number, number]
    // 初始缩放级别
    zoom?: number
    // 地图样式对象，优先级高于环境变量 VITE_MAP_TILES_URL
    style?: object
}>()

const emit = defineEmits<{
    (e: 'update:modelValue', value: string): void
    (e: 'confirm', value: string): void
}>()

const visible = ref(false)
const loading = ref(false)
const mapRef = ref<HTMLElement>()
const geojsonStr = ref('')

let map: MapLibreMap | null = null
let draw: MapboxDraw | null = null
// 防止删除冗余区域时递归触发事件
let syncing = false

/**
 * mapbox-gl-draw 内部按 mapboxgl-* 的 DOM 类名定位控件与画布，
 * 在 maplibre-gl 下必须换成 maplibregl-*，否则控件与事件绑定失效。
 * 参考 maplibre 官方示例与 mapbox-gl-draw#1357
 */
const patchDrawClasses = () => {
    const classes = MapboxDraw.constants.classes
    classes.CANVAS = 'maplibregl-canvas'
    classes.CONTROL_BASE = 'maplibregl-ctrl'
    classes.CONTROL_PREFIX = 'maplibregl-ctrl-'
    classes.CONTROL_GROUP = 'maplibregl-ctrl-group'
    classes.ATTRIBUTION = 'maplibregl-ctrl-attrib'
}
patchDrawClasses()

/**
 * mapbox-gl-draw 把 [0.2, 2] 这类裸数组直接写进 case 表达式的分支返回值，
 * maplibre 的 style-spec 校验更严格（必须写成 ["literal", [...]]），
 * 否则 gl-draw-lines 的 cold/hot 图层 addLayer 失败且只抛到控制台，
 * 表现为区域边框不显示（填充、顶点正常，容易漏测）。
 */
const patchDrawTheme = () => {
    const lines = MapboxDraw.lib.theme.find((layer) => layer.id === 'gl-draw-lines')
    const dash = lines && lines.paint && lines.paint['line-dasharray']
    if (!Array.isArray(dash) || dash[0] !== 'case') return
    // case 的第 2、3 项是两个分支的返回值
    for (const i of [2, 3]) {
        const value = dash[i]
        if (Array.isArray(value) && value[0] !== 'literal') dash[i] = ['literal', value]
    }
}
patchDrawTheme()

// mapbox-gl-draw 的 draw.* 事件由插件自行 fire，不在 maplibre 的 MapEventType 定义中，
// 这里统一收口做一次断言，避免到处写 ts-ignore
type DrawEvent = 'draw.create' | 'draw.update' | 'draw.delete'
const onDrawEvent = (type: DrawEvent, handler: () => void) => {
    const target = map as unknown as { on?: (t: string, fn: () => void) => void } | null
    target?.on?.(type, handler)
}

// 样式解析优先级：props.style > VITE_MAP_TILES_URL > OSM 默认
const resolveStyle = (): StyleSpecification => {
    if (props.style) return props.style as StyleSpecification
    const tiles = import.meta.env.VITE_MAP_TILES_URL as string
    return {
        version: 8,
        sources: {
            custom: {
                type: 'raster',
                tiles: [tiles || ''],
                tileSize: 256,
                attribution: ''
            }
        },
        layers: [{ id: 'custom-raster', type: 'raster', source: 'custom' }]
    }
}

// 取当前多边形列表
const getPolygons = () => {
    const features = draw?.getAll().features || []
    return features.filter((f) => f.geometry?.type === 'Polygon')
}

/**
 * 规范化 Polygon 环：删除连续重复点并保证首尾闭合。
 * 绘制时可能出现重复顶点（或在旧数据里被追加点），直接存库会产生脏 GeoJSON。
 */
const normalizePolygon = (coordinates: number[][][]) => {
    return coordinates.map((ring) => {
        const out: number[][] = []
        for (const p of ring) {
            const last = out[out.length - 1]
            if (last && last[0] === p[0] && last[1] === p[1]) continue
            out.push(p)
        }
        if (out.length > 3) {
            const first = out[0]
            const last = out[out.length - 1]
            if (first[0] !== last[0] || first[1] !== last[1]) out.push([first[0], first[1]])
        }
        return out
    })
}

const toFeatureJson = (geometry: any) => {
    if (geometry?.type !== 'Polygon') return ''
    return JSON.stringify({
        type: 'Feature',
        properties: {},
        geometry: { type: 'Polygon', coordinates: normalizePolygon(geometry.coordinates || []) }
    })
}

// 绘制内容变化时同步 geojson，并强制只保留一个区域
const syncFromDraw = () => {
    if (syncing || !draw) return
    const polys = getPolygons()
    // 多于一个时删除较早的，保证单一区域
    if (polys.length > 1) {
        syncing = true
        try {
            draw.delete(polys.slice(0, -1).map((f) => String(f.id)))
        } finally {
            syncing = false
        }
    }
    const cur = polys.length ? polys[polys.length - 1] : null
    geojsonStr.value = cur ? toFeatureJson(cur.geometry) : ''
}

const fitToPolygon = () => {
    const polys = getPolygons()
    const cur = polys[polys.length - 1]
    if (!cur || !map) return
    const coords: number[][] = (cur.geometry as any)?.coordinates?.[0] || []
    if (!coords.length) return
    const lngs = coords.map((c) => c[0])
    const lats = coords.map((c) => c[1])
    map.fitBounds(
        [
            [Math.min(...lngs), Math.min(...lats)],
            [Math.max(...lngs), Math.max(...lats)]
        ],
        { padding: 60, maxZoom: 16 }
    )
}

const handleFit = () => fitToPolygon()

const handleClear = () => {
    // deleteAll 不会 fire 事件，geojsonStr 需手动清空
    draw?.deleteAll()
    geojsonStr.value = ''
}

const initMap = () => {
    if (!mapRef.value) return
    map = new MapLibreMap({
        container: mapRef.value,
        style: resolveStyle(),
        center: props.center?.length ? props.center : [104.06, 30.57],
        zoom: props.zoom ?? 4
    })
    map.addControl(new NavigationControl(), 'top-right')

    map.on('load', () => {
        draw = new MapboxDraw({
            defaultMode: 'simple_select',
            displayControlsDefault: false,
            // 只保留：绘制多边形 + 删除
            controls: { polygon: true, trash: true }
        })
        map?.addControl(draw, 'top-left')

        onDrawEvent('draw.create', syncFromDraw)
        onDrawEvent('draw.update', syncFromDraw)
        onDrawEvent('draw.delete', syncFromDraw)

        // 回显已有区域
        loadFromGeoJson(props.modelValue || '')
        loading.value = false
    })
}

// 回显：用 set 原子替换全部内容
// 注意：draw.set()/deleteAll() 是 API 调用，默认 suppressAPIEvents=true 不会 fire draw.* 事件，
// 所以 geojsonStr 必须在这里手动同步，不能指望事件
const loadFromGeoJson = (str: string) => {
    if (!draw) return
    draw.deleteAll()
    geojsonStr.value = ''
    if (!str) return
    try {
        const data = JSON.parse(str)
        if (data?.geometry?.type === 'Polygon') {
            const feature = JSON.parse(toFeatureJson(data.geometry))
            draw.set({ type: 'FeatureCollection', features: [feature] })
            geojsonStr.value = JSON.stringify(feature)
            setTimeout(fitToPolygon, 100)
        }
    } catch (e) {
        console.warn('解析 GeoJSON 失败', e)
        ElMessage.warning('已有区域不是合法的 GeoJSON，已跳过回显')
    }
}

const handleOpened = () => {
    loading.value = true
    initMap()
}

const handleClosed = () => {
    draw = null
    if (map) {
        map.remove()
        map = null
    }
}

const handleCopy = async () => {
    if (!geojsonStr.value) return ElMessage('暂无可复制的 GeoJSON')
    try {
        await navigator.clipboard.writeText(geojsonStr.value)
        ElMessage.success('已复制到剪贴板')
    } catch {
        ElMessage.error('复制失败，请手动选择文本复制')
    }
}

const handleConfirm = () => {
    if (!geojsonStr.value) return ElMessage('请先在地图上绘制区域(至少 3 个点)')
    emit('update:modelValue', geojsonStr.value)
    emit('confirm', geojsonStr.value)
    visible.value = false
}

const open = () => {
    visible.value = true
}

defineExpose({ open })
</script>

<style scoped lang="scss">
.map-area-picker {
    .map-area-picker__toolbar {
        display: flex;
        align-items: center;
        gap: 8px;
        margin-bottom: 10px;

        .map-area-picker__tips {
            margin-right: auto;
            font-size: 13px;
            color: var(--el-text-color-secondary);
        }
    }

    .map-area-picker__map {
        width: 100%;
        height: 460px;
        border: 1px solid var(--el-border-color);
        border-radius: 6px;
        overflow: hidden;
    }

    .map-area-picker__footer {
        margin-top: 10px;
    }
}
</style>

<!-- mapbox-gl-draw 的光标样式类名挂在 maplibre 生成的容器上，scoped 样式无法命中，故全局声明；
     选择器已用 .map-area-picker__map 限定，不会污染其他地图 -->
<style lang="scss">
.map-area-picker__map.maplibregl-map {
    &.mouse-pointer .maplibregl-canvas-container.maplibregl-interactive {
        cursor: pointer;
    }
    &.mouse-move .maplibregl-canvas-container.maplibregl-interactive {
        cursor: move;
    }
    &.mouse-add .maplibregl-canvas-container.maplibregl-interactive {
        cursor: crosshair;
    }
    &.mouse-move.mode-direct_select .maplibregl-canvas-container.maplibregl-interactive {
        cursor: grab;
    }
    &.mode-direct_select.feature-vertex.mouse-move .maplibregl-canvas-container.maplibregl-interactive {
        cursor: move;
    }
    &.mode-direct_select.feature-midpoint.mouse-pointer .maplibregl-canvas-container.maplibregl-interactive {
        cursor: cell;
    }
    &.mode-direct_select.feature-feature.mouse-move .maplibregl-canvas-container.maplibregl-interactive {
        cursor: move;
    }
    &.mode-static.mouse-pointer .maplibregl-canvas-container.maplibregl-interactive {
        cursor: grab;
    }
}
</style>
