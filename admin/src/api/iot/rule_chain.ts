import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// ===================== 可视化规则链（替代产品级解码/过滤/告警脚本） =====================
// graph_data 为 LogicFlow 原生图 { nodes, edges }（仅供绘图）；tree_data 为同源展开的执行树（后端消费），
// 节点的私有属性按命名空间收纳于 properties
export type type_rule_chain = {
    id?: string
    name?: string
    remark?: string
    graph_data?: string
    tree_data?: string
    is_stop?: number
    create_time?: string
    update_time?: string
}
export type type_rule_chain_list = {
    name?: string
    is_stop?: number
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
