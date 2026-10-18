import request from '@/utils/request'
import type { Pages } from '@/utils/request'
import type { NodeType } from './product'

// ===================== 设备 =====================
export type type_iot_device_list = {
    name?: string
    sn?: string
    product_id?: string
    space_id?: string
    status?: number
    start_time?: string
    end_time?: string
    pageNo?: number
    pageSize?: number
}
export type type_iot_device = {
    id?: string
    name?: string
    sn?: string
    product_id?: string
    product_name?: string
    space_id?: string
    space_name?: string
    gateway_id?: string
    gateway_name?: string
    status?: number
    offline_timeout?: number
    last_online_time?: string
    last_offline_time?: string
    product_node_type?: NodeType | ''
    longitude?: number
    latitude?: number
    access_token?: string
    remark?: string
    create_time?: string
    update_time?: string
}
export function deviceLists(params: type_iot_device_list) {
    return request.get<Pages<type_iot_device>>({ url: '/iot/device/list', params })
}
export function deviceDetail(params: { id: string }) {
    return request.get<type_iot_device>({ url: '/iot/device/detail', params })
}
export function deviceAdd(data: type_iot_device) {
    return request.post({ url: '/iot/device/add', data })
}
export function deviceEdit(data: type_iot_device) {
    return request.post({ url: '/iot/device/edit', data })
}
export function deviceDel(data: { id: string }) {
    return request.post({ url: '/iot/device/del', data })
}

// 设备运行态快照（Redis 侧：在线判定/最后报文/最新属性），供详情弹窗实时 tab
export type type_device_state = {
    device_id?: string
    sn?: string
    name?: string
    online?: boolean
    last_seen_ms?: number
    updated_at?: string
    values?: Record<string, string>
    db_status?: number
    last_online_time?: string
    last_offline_time?: string
    offline_timeout?: number
}
export function deviceState(params: { id: string }) {
    return request.get<type_device_state>({ url: '/iot/device/state', params })
}

// 时序历史区间回读（趋势图数据源，毫秒时间戳，0=不限）
// values 已由后端按产品物模型约束，只包含已建模的属性
export type type_device_data = {
    ts: number
    values: Record<string, any>
}
export function deviceData(params: { id: string; start_ms?: number; end_ms?: number; limit?: number }) {
    return request.get<type_device_data[]>({ url: '/iot/device/data', params })
}

// 大屏设备点（不分页）：space_ids 逗号分隔，选父区域时把子树 id 集齐传入，空=全部
export type type_device_map_item = {
    id: string
    sn: string
    name: string
    space_id: string
    product_id: string
    longitude: number
    latitude: number
    status: number
    // 活动告警级别：0=无告警 -1=故障 1=轻微 2=严重（数值即危害序，3+ 预留；!=0 即告警中，有故障则直接取故障）
    severity: number
}
export function deviceMap(params?: { space_ids?: string }) {
    return request.get<type_device_map_item[]>({ url: '/iot/device/map', params })
}
