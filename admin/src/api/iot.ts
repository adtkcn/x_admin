import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// ===================== 产品分类（树） =====================
export type type_iot_product_category = {
    id?: string
    pid?: string
    name?: string
    sort?: number
    remark?: string
    create_time?: string
    update_time?: string
    children?: type_iot_product_category[]
}
export function productCategoryAll() {
    return request.get<type_iot_product_category[]>({ url: '/iot/productCategory/all' })
}
export function productCategoryDetail(params: { id: string }) {
    return request.get<type_iot_product_category>({ url: '/iot/productCategory/detail', params })
}
export function productCategoryAdd(data: type_iot_product_category) {
    return request.post({ url: '/iot/productCategory/add', data })
}
export function productCategoryEdit(data: type_iot_product_category) {
    return request.post({ url: '/iot/productCategory/edit', data })
}
export function productCategoryDel(data: { id: string }) {
    return request.post({ url: '/iot/productCategory/del', data })
}
export function productCategorySort(data: { ids: string[] }) {
    return request.post({ url: '/iot/productCategory/sort', data })
}

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
    name?: string
    category_id?: string
    pageNo?: number
    pageSize?: number
}
export type type_iot_product = {
    id?: string
    name?: string
    category_id?: string
    category_name?: string
    device_type?: NodeType | ''
    protocol?: string
    thing_model?: string
    decode_script?: string
    preprocess_script?: string
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

// ===================== 设备 =====================
export type type_iot_device_list = {
    name?: string
    sn?: string
    product_id?: string
    space_id?: string
    status?: number
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
