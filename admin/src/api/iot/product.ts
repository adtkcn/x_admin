import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 节点类型：网关身份由产品节点类型定义（对齐阿里云/腾讯云 IoT 惯例）
export type NodeType = 'DEVICE' | 'GATEWAY' | 'EDGE'
export const NodeTypeMap: Record<NodeType, string> = {
    DEVICE: '直连设备',
    GATEWAY: '网关设备',
    EDGE: '边缘主机'
}
// 是否具备网关能力（GATEWAY/EDGE 可挂子设备；空值为历史遗留数据，按直连设备处理）
export function nodeTypeSupportsGateway(nodeType?: string) {
    return nodeType === 'GATEWAY' || nodeType === 'EDGE'
}

// ===================== 产品 =====================
export type type_iot_product_list = {
    code?: string
    name?: string
    category_id?: string
    is_stop?: number
    pageNo?: number
    pageSize?: number
}
export type type_iot_product = {
    id?: string
    code?: string
    name?: string
    category_id?: string
    category_name?: string
    device_type?: NodeType | ''
    protocol?: string
    offline_timeout?: number
    thing_model?: string
    rule_chain_id?: string
    rule_chain_name?: string
    remark?: string
    is_stop?: number
    create_time?: string
    update_time?: string
}
export function productLists(params: type_iot_product_list) {
    return request.get<Pages<type_iot_product>>({ url: '/iot/product/list', params })
}
export function productAll() {
    return request.get<type_iot_product[]>({ url: '/iot/product/all' })
}
export function productDetail(params: { id: string }) {
    return request.get<type_iot_product>({ url: '/iot/product/detail', params })
}
export function productAdd(data: type_iot_product) {
    return request.post({ url: '/iot/product/add', data })
}
export function productEdit(data: type_iot_product) {
    return request.post({ url: '/iot/product/edit', data })
}
export function productDel(data: { id: string }) {
    return request.post({ url: '/iot/product/del', data })
}
