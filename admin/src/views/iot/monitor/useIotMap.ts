import { shallowRef, type Ref } from 'vue'
import { Map as MapLibreMap, NavigationControl, Marker } from 'maplibre-gl'
import type { StyleSpecification } from 'maplibre-gl'
import {
    SeverityMap,
    SeverityColor,
    type type_iot_device_space,
    type type_device_map_item
} from '@/api/iot'

// 大屏地图组合式函数：只负责 MapLibre 实例、样式/图层/数据源、区域面与设备点的渲染、
// 点击命中路由、popup 呈现与视野适配；设备数据(在线状态/告警级别)与最新值由调用方以 ref 传入，
// 区域选择与「点开设备回读存量值」等副作用通过回调交还调用方。

export interface UseIotMapOptions {
    container: Ref<HTMLElement | undefined>
    points: Ref<type_device_map_item[]>
    spaces: Ref<type_iot_device_space[]>
    currentSpace: Ref<string>
    // popup 取该设备当前实时值（无则调用方回读快照后经 syncPopup 重渲）
    getPopupValues: (deviceId: string) => Record<string, any> | undefined
    // 点中区域轮廓 → 选中该区域；点开设备 popup → 通知调用方回读存量值
    onSelectSpace: (spaceId: string) => void
    onPopupOpen: (deviceId: string) => void
}

export function useIotMap(opts: UseIotMapOptions) {
    const { container, points, spaces, currentSpace, getPopupValues, onSelectSpace, onPopupOpen } = opts

    const map = shallowRef<MapLibreMap>()
    const popupMarker = shallowRef<Marker>()
    // 当前 popup 上下文：设备点位数据变化时按此重渲内容
    const popupCtx = shallowRef<{ id: string; lngLat: { lng: number; lat: number } } | null>(null)

    // 设备索引：WS 报文与 popup 按 device_id 直接命中点位，避免每条报文 O(n) 查找
    const pointIndex = new Map<string, type_device_map_item>()
    const rebuildIndex = () => {
        pointIndex.clear()
        for (const p of points.value) pointIndex.set(p.id!, p)
    }

    // 点位颜色按级别取字典；未录入的更高级别(3+)一律取顶格红（故障是低档，不能当兜底色）
    const sevColor = (lv: number) => SeverityColor[lv] || SeverityColor[2]

    // ---------------- 样式解析（与 map-picker 同款：环境变量瓦片源优先，缺省 OSM） ----------------
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

    // ---------------- GeoJSON 组装（纯地图数据转换） ----------------
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

    const areasFeatureCollection = () => {
        // 选中区域时展示该区域及其直接子级（一层）轮廓；未选中(全部)时只展示第一级根区域
        let idSet: Set<string> | null = null
        if (currentSpace.value) {
            idSet = new Set([currentSpace.value])
            for (const s of spaces.value) {
                if (s.id && s.pid === currentSpace.value) idSet.add(s.id)
            }
        } else {
            idSet = new Set()
            for (const s of spaces.value) {
                if (s.id && !s.pid) idSet.add(s.id)
            }
        }
        const features: any[] = []
        for (const s of spaces.value) {
            if (!s.geojson) continue
            if (idSet && (!s.id || !idSet.has(s.id))) continue
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

    const devicesFeatureCollection = () => ({
        type: 'FeatureCollection' as const,
        features: points.value
            .filter((p) => p.longitude !== 0 || p.latitude !== 0)
            .map((p) => ({
                type: 'Feature' as const,
                geometry: { type: 'Point' as const, coordinates: [p.longitude, p.latitude] },
                // status/severity 必须带进 properties，图层 paint 靠它们区分在线/离线与告警级别配色
                // severity 是危害序且含负数档，0 才是“无告警”，不能当数组下标用
                properties: { id: p.id, status: p.status ?? 0, severity: p.severity ?? 0 }
            }))
    })

    const refreshDevices = () => {
        const src = map.value?.getSource('devices') as any
        if (src) src.setData(devicesFeatureCollection())
    }

    // 切换区域选择时同步刷新轮廓：只保留当前子树的区域面
    const refreshAreas = () => {
        const src = map.value?.getSource('areas') as any
        if (src) src.setData(areasFeatureCollection())
    }

    // 数据/点位变化后统一重算：重建索引 + 刷新点与面
    const refreshAll = () => {
        rebuildIndex()
        refreshDevices()
        refreshAreas()
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

    // ---------------- popup ----------------
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
        if (p.severity !== 0)
            addLine(box, `告警中·${SeverityMap[p.severity] || p.severity}`, sevColor(p.severity), true)
        const vals = getPopupValues(p.id!)
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
        onPopupOpen(id)
    }

    const closePopup = () => {
        popupCtx.value = null
        popupMarker.value?.remove()
    }

    // 仅当该设备仍是当前 popup 时重渲（供 WS 增量/异步回读值到达后调用）
    const syncPopup = (id?: string) => {
        if (id && popupCtx.value?.id !== id) return
        renderPopup()
    }

    // 选中子树后把视野适配到其区域轮廓范围
    const fitToSubtree = (ids: string[]) => {
        if (!map.value || !ids.length) return
        let minX = Infinity,
            minY = Infinity,
            maxX = -Infinity,
            maxY = -Infinity
        // 遍历 Polygon/MultiPolygon 各环求并集范围：行政区选择产出的边界多为 MultiPolygon，只认 Polygon 会整体跳过导致无法缩放
        const scan = (ring: number[][]) => {
            for (const [lng, lat] of ring || []) {
                if (lng < minX) minX = lng
                if (lng > maxX) maxX = lng
                if (lat < minY) minY = lat
                if (lat > maxY) maxY = lat
            }
        }
        for (const s of spaces.value) {
            if (!s.id || !ids.includes(s.id) || !s.geojson) continue
            try {
                const geom = JSON.parse(s.geojson)?.geometry
                if (geom?.type === 'Polygon') geom.coordinates.forEach(scan)
                else if (geom?.type === 'MultiPolygon')
                    geom.coordinates.forEach((poly: number[][][]) => poly.forEach(scan))
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

    // ---------------- 初始化与图层 ----------------
    const init = () => {
        if (!container.value) return
        map.value = new MapLibreMap({
            container: container.value,
            style: resolveStyle(),
            center: [104.0668, 30.5728],
            zoom: 4.5
        })
        map.value.addControl(new NavigationControl({ showCompass: false }), 'top-right')
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
            // 设备点位：告警按级别配色(带光晕) / 在线绿 / 离线灰(淡化)，非 0 级别优先于在线
            map.value!.addSource('devices', { type: 'geojson', data: devicesFeatureCollection() })
            // 告警光晕：仅告警设备，危害越高光晕越大越实（match 末尾是无标签的默认分支，接住 3+ 预留级）
            map.value!.addLayer({
                id: 'devices-halo',
                type: 'circle',
                source: 'devices',
                filter: ['!=', ['get', 'severity'], 0],
                paint: {
                    'circle-radius': ['match', ['get', 'severity'], -1, 12, 1, 13, 2, 16, 20],
                    'circle-color': [
                        'match',
                        ['get', 'severity'],
                        -1,
                        sevColor(-1),
                        1,
                        sevColor(1),
                        2,
                        sevColor(2),
                        sevColor(2)
                    ],
                    'circle-opacity': ['match', ['get', 'severity'], -1, 0.12, 1, 0.14, 2, 0.2, 0.3],
                    'circle-blur': 0.7
                }
            } as any)
            map.value!.addLayer({
                id: 'devices-circle',
                type: 'circle',
                source: 'devices',
                paint: {
                    'circle-radius': [
                        'case',
                        ['!=', ['get', 'severity'], 0],
                        ['match', ['get', 'severity'], -1, 7, 1, 8, 2, 9, 11],
                        ['==', ['get', 'status'], 1],
                        7,
                        6
                    ],
                    'circle-stroke-width': ['case', ['!=', ['get', 'severity'], 0], 2.5, 1.5],
                    'circle-stroke-color': '#fff',
                    'circle-color': [
                        'case',
                        ['!=', ['get', 'severity'], 0],
                        ['match', ['get', 'severity'], -1, sevColor(-1), 1, sevColor(1), 2, sevColor(2), sevColor(2)],
                        ['==', ['get', 'status'], 1],
                        '#22c55e',
                        '#64748b'
                    ],
                    // 离线点位略作淡化，在线/告警保持醒目
                    'circle-opacity': [
                        'case',
                        ['!=', ['get', 'severity'], 0],
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
                    onSelectSpace(String(area.properties?.space_id || ''))
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

    const resize = () => map.value?.resize()

    const destroy = () => {
        if (flushTimer) clearTimeout(flushTimer)
        closePopup()
        popupMarker.value = undefined
        map.value?.remove()
        map.value = undefined
    }

    return {
        init,
        refreshAll,
        refreshDevices,
        refreshAreas,
        markDirty,
        fitToSubtree,
        closePopup,
        syncPopup,
        resize,
        destroy,
        getPoint: (id: string) => pointIndex.get(id)
    }
}
