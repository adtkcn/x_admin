import request from '@/utils/request'

// ===================== 设备区域（树） =====================
export type type_iot_device_space = {
    id?: string
    pid?: string
    name?: string
    sort?: number
    geojson?: string
    remark?: string
    create_time?: string
    update_time?: string
    children?: type_iot_device_space[]
}
export function deviceSpaceAll() {
    return request.get<type_iot_device_space[]>({ url: '/iot/deviceSpace/all' })
}
export function deviceSpaceDetail(params: { id: string }) {
    return request.get<type_iot_device_space>({ url: '/iot/deviceSpace/detail', params })
}
export function deviceSpaceAdd(data: type_iot_device_space) {
    return request.post({ url: '/iot/deviceSpace/add', data })
}
export function deviceSpaceEdit(data: type_iot_device_space) {
    return request.post({ url: '/iot/deviceSpace/edit', data })
}
export function deviceSpaceDel(data: { id: string }) {
    return request.post({ url: '/iot/deviceSpace/del', data })
}
export function deviceSpaceSort(data: { ids: string[] }) {
    return request.post({ url: '/iot/deviceSpace/sort', data })
}
