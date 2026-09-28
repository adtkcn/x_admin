<template>
    <div class="map-picker">
        <el-dialog
            v-model="visible"
            :title="title"
            :width="width"
            :append-to-body="true"
            :close-on-click-modal="false"
            draggable
            destroy-on-close
            top="5px"
            class="map-picker__dialog"
            @opened="handleOpened"
            @closed="handleClosed"
        >
            <div class="map-picker__toolbar">
                <span class="map-picker__tips">点击地图或拖动标记选择坐标点</span>
                <el-button :icon="Aim" size="small" @click="handleLocate">定位</el-button>
            </div>
            <div ref="mapContainer" class="map-picker__map"></div>
            <template #footer>
                <div class="map-picker__footer">
                    <span class="map-picker__current"> 当前坐标：{{ formattedCurrent }} </span>
                    <div>
                        <el-button @click="handleCancel">取消</el-button>
                        <el-button type="primary" @click="handleConfirm">确定</el-button>
                    </div>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, computed, shallowRef, onBeforeUnmount } from 'vue'
import { Map as MapLibreMap, Marker, NavigationControl } from 'maplibre-gl'
import type { StyleSpecification, MapMouseEvent } from 'maplibre-gl'
import 'maplibre-gl/dist/maplibre-gl.css'
import { Aim } from '@element-plus/icons-vue'
import feedback from '@/utils/feedback'
import type { MapCoordinate } from './type'

const props = withDefaults(
    defineProps<{
        // v-model：{ longitude, latitude }
        modelValue?: MapCoordinate
        title?: string
        width?: string
        // 无坐标时的默认中心（中国中心）
        center?: [number, number]
        // 默认缩放级别
        zoom?: number
        // 自定义 maplibre 样式；不传则用瓦片源/内置 OSM
        style?: StyleSpecification | string
    }>(),
    {
        modelValue: () => ({ longitude: 0, latitude: 0 }),
        title: '地图选点',
        width: '900px',
        center: () => [104.0668, 30.5728],
        zoom: 5,
        style: undefined
    }
)

const emit = defineEmits<{
    (e: 'update:modelValue', val: MapCoordinate): void
    (e: 'confirm', val: MapCoordinate): void
    (e: 'close'): void
}>()

const visible = ref(false)
const mapContainer = ref<HTMLElement>()
const map = shallowRef<MapLibreMap>()
const marker = shallowRef<Marker>()

const coords = reactive<MapCoordinate>({
    longitude: 0,
    latitude: 0
})

// 判定坐标是否已设置（0,0 视为未设置）
const hasCoordinate = (val?: MapCoordinate) =>
    !!val &&
    Number.isFinite(val.longitude) &&
    Number.isFinite(val.latitude) &&
    !(val.longitude === 0 && val.latitude === 0)

const formattedCurrent = computed(() => {
    if (!hasCoordinate(coords)) return '未选择'
    return `${coords.longitude.toFixed(6)}, ${coords.latitude.toFixed(6)}`
})

// 解析地图样式：props.style > 环境变量 VITE_MAP_TILES_URL > 内置 OSM
const resolveStyle = (): StyleSpecification | string => {
    if (props.style) return props.style
    const tiles = import.meta.env.VITE_MAP_TILES_URL as string | undefined

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

// 放置/移动标记
const placeMarker = (longitude: number, latitude: number) => {
    const lngLat: [number, number] = [longitude, latitude]
    if (!marker.value) {
        marker.value = new Marker({ draggable: true, color: '#409eff' })
            .setLngLat(lngLat)
            .addTo(map.value as MapLibreMap)
        marker.value.on('dragend', () => {
            const p = marker.value!.getLngLat()
            coords.longitude = +p.lng.toFixed(6)
            coords.latitude = +p.lat.toFixed(6)
        })
    } else {
        marker.value.setLngLat(lngLat)
    }
}

const initMap = () => {
    if (map.value || !mapContainer.value) return
    map.value = new MapLibreMap({
        container: mapContainer.value,
        style: resolveStyle(),
        center: props.center,
        zoom: props.zoom
    })
    map.value.addControl(new NavigationControl({ showCompass: false }), 'top-right')
    map.value.on('click', (e: MapMouseEvent) => {
        coords.longitude = +e.lngLat.lng.toFixed(6)
        coords.latitude = +e.lngLat.lat.toFixed(6)
        placeMarker(coords.longitude, coords.latitude)
    })
}

// el-dialog 打开动画结束后再初始化，保证容器已有尺寸
const handleOpened = () => {
    initMap()
    // 用当前 props 初始化选点状态
    coords.longitude = props.modelValue?.longitude ?? 0
    coords.latitude = props.modelValue?.latitude ?? 0
    if (hasCoordinate(coords)) {
        placeMarker(coords.longitude, coords.latitude)
        map.value?.flyTo({ center: [coords.longitude, coords.latitude], zoom: 14 })
    } else if (marker.value) {
        marker.value.remove()
        marker.value = undefined
    }
    map.value?.resize()
}

const handleClosed = () => {
    // destroy-on-close 会销毁容器，这里移除地图实例避免内存泄漏
    if (map.value) {
        map.value.remove()
        map.value = undefined
        marker.value = undefined
    }
    emit('close')
}

// 浏览器定位
const handleLocate = () => {
    if (!navigator.geolocation) {
        feedback.msgError('当前浏览器不支持定位')
        return
    }
    navigator.geolocation.getCurrentPosition(
        (pos) => {
            coords.longitude = +pos.coords.longitude.toFixed(6)
            coords.latitude = +pos.coords.latitude.toFixed(6)
            if (map.value) {
                placeMarker(coords.longitude, coords.latitude)
                map.value.flyTo({ center: [coords.longitude, coords.latitude], zoom: 14 })
            }
        },
        () => feedback.msgError('定位失败，请检查定位权限')
    )
}

const handleConfirm = () => {
    const val: MapCoordinate = { longitude: coords.longitude, latitude: coords.latitude }
    emit('update:modelValue', val)
    emit('confirm', val)
    visible.value = false
}

const handleCancel = () => {
    visible.value = false
}

const open = () => {
    visible.value = true
}

onBeforeUnmount(() => {
    map.value?.remove()
    map.value = undefined
})

defineExpose({ open })
</script>

<style lang="scss" scoped>
.map-picker {
    display: inline-block;
}
.map-picker__toolbar {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 10px;
    .map-picker__tips {
        margin-right: auto;
        font-size: 12px;
        color: var(--el-text-color-secondary);
    }
}
.map-picker__map {
    width: 100%;
    height: 460px;
    border-radius: 6px;
    overflow: hidden;
}
.map-picker__footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
}
.map-picker__current {
    font-size: 13px;
    color: var(--el-text-color-regular);
}
</style>

<style lang="scss">
.map-picker__dialog .el-dialog__body {
    padding-top: 10px;
}
</style>
