import request from '@/utils/request/index'
import type { Pages } from '@/utils/request'

export type type_flow_history = {
    id?: number
    applyId?: number
    templateId?: number
    applyUserId?: number
    applyUserNickname?: string
    approverId?: number
    approverNickname?: string
    nodeId?: string
    nodeType?: string
    nodeLabel?: string
    formValue?: string
    passStatus?: number
    passRemark?: string
    createTime?: string
    updateTime?: string
    deleteTime?: string
}
// 查询
export type type_flow_history_query = {
    applyId?: number
    templateId?: number
    applyUserId?: number
    applyUserNickname?: string
    approverId?: number
    approverNickname?: string
    nodeId?: string
    nodeType?: string
    nodeLabel?: string
    formValue?: string
    passStatus?: number
    passRemark?: string
    createTimeStart?: string
    createTimeEnd?: string
    updateTimeStart?: string
    updateTimeEnd?: string
}
// 添加编辑
export type type_flow_history_edit = {
    id?: number
    applyId?: number
    templateId?: number
    applyUserId?: number
    applyUserNickname?: string
    approverId?: number
    approverNickname?: string
    nodeId?: string
    nodeType?: string
    nodeLabel?: string
    formValue?: string
    passStatus?: number
    passRemark?: string
}

// 流程历史列表
export function flow_history_list(params?: type_flow_history_query) {
    return request.get<Pages<type_flow_history>>({ url: '/flow/flow_history/list', params })
}
// 流程历史列表-所有
export function flow_history_list_all(params?: type_flow_history_query) {
    return request.get<Pages<type_flow_history>>({ url: '/flow/flow_history/listAll', params })
}

// 流程历史详情
export function flow_history_detail(id: number | string) {
    return request.get<type_flow_history>({ url: '/flow/flow_history/detail', params: { id } })
}

// 流程历史新增
export function flow_history_add(data: type_flow_history_edit) {
    return request.post<null>({ url: '/flow/flow_history/add', data })
}

// 流程历史编辑
export function flow_history_edit(data: type_flow_history_edit) {
    return request.post<null>({ url: '/flow/flow_history/edit', data })
}

// 流程历史删除
export function flow_history_delete(id: number | string) {
    return request.post<null>({ url: '/flow/flow_history/del', data: { id } })
}

// 获取下一个审批节点，中间可能有系统任务和结束节点被跳过
export function flow_history_next_node(data: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/next_node', data })
}

// 获取下一个审批节点，中间可能有系统任务和结束节点被跳过
export function flow_history_get_approver(data: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/get_approver', data })
}
export function flow_history_pass(data: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/pass', data })
}

export function flow_history_back(data: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/back', data })
}
