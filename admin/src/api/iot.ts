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

// ===================== 可视化规则链（替代产品级解码/过滤/告警脚本） =====================
// graph_data 为 LogicFlow 原生图 { nodes, edges }，节点的私有属性按命名空间收纳于 properties
export type type_rule_chain = {
    id?: string
    name?: string
    remark?: string
    graph_data?: string
    is_stop?: number
    create_time?: string
    update_time?: string
}
export type type_rule_chain_list = {
    name?: string
    pageNo?: number
    pageSize?: number
}
export function ruleChainLists(params: type_rule_chain_list) {
    return request.get<Pages<type_rule_chain>>({ url: '/iot/ruleChain/list', params })
}
export function ruleChainAll() {
    return request.get<type_rule_chain[]>({ url: '/iot/ruleChain/all' })
}
export function ruleChainDetail(params: { id: string }) {
    return request.get<type_rule_chain>({ url: '/iot/ruleChain/detail', params })
}
export function ruleChainAdd(data: type_rule_chain) {
    return request.post({ url: '/iot/ruleChain/add', data })
}
export function ruleChainEdit(data: type_rule_chain) {
    return request.post({ url: '/iot/ruleChain/edit', data })
}
export function ruleChainDel(data: { id: string }) {
    return request.post({ url: '/iot/ruleChain/del', data })
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
    offline_in?: number
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
export type type_device_data = {
    ts: number
    values: Record<string, any>
    raw: string
    protocol: string
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
    alerting: boolean
}
export function deviceMap(params?: { space_ids?: string }) {
    return request.get<type_device_map_item[]>({ url: '/iot/device/map', params })
}

// ===================== 告警历史（处理层告警状态机自动写入，管理端只读/清理） =====================
// 告警级别：由产品 goja 告警脚本 alert(values) 返回值决定（0=不告警），字典供列表/表单展示
export const SeverityMap: Record<number, string> = {
    1: '轻微',
    2: '严重',
    3: '故障'
}
// tag 颜色：级别越高越醒目
export const SeverityTagType: Record<number, 'success' | 'warning' | 'danger' | 'info'> = {
    1: 'warning',
    2: 'danger',
    3: 'danger'
}
export type type_iot_alert_list = {
    device_id?: string
    sn?: string
    product_id?: string
    space_id?: string
    status?: number
    severity?: number
    start_time?: string
    end_time?: string
    pageNo?: number
    pageSize?: number
}
export type type_iot_alert = {
    id?: string
    device_id?: string
    sn?: string
    product_id?: string
    product_name?: string
    space_id?: string
    space_name?: string
    status?: number
    severity?: number
    values?: Record<string, any>
    triggered_at?: string
    recovered_at?: string
}
export function alertLists(params: type_iot_alert_list) {
    return request.get<Pages<type_iot_alert>>({ url: '/iot/alert/list', params })
}
export function alertDel(data: { ids: string[] }) {
    return request.post({ url: '/iot/alert/del', data })
}
