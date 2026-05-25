import request from '@/utils/request/index'
import type { Pages } from '@/utils/request'

export type type_flow_history = {
    id: string
    applyId: string | null
    templateId: string | null
    applyUserId: string | null
    applyUserNickname: string | null
    approverId: string | null
    approverNickname: string | null
    nodeId: string | null
    nodeType: string | null
    nodeLabel: string | null
    formValue: string | null
    passStatus: number | null // 通过状态：1待处理，2通过，3拒绝
    passRemark: string | null
    isShow: number | null // 是否显示：0隐藏，1显示
    createTime: string | null
    updateTime: string | null
    deleteTime: string | null
}

export type type_flow_history_query = {
    applyId?: string
    templateId?: string
    applyUserId?: string
    applyUserNickname?: string
    approverId?: string
    approverNickname?: string
    nodeId?: string
    nodeType?: string
    nodeLabel?: string
    formValue?: string
    passStatus?: number
    passRemark?: string
    isShow?: number
    createTimeStart?: string
    createTimeEnd?: string
    updateTimeStart?: string
    updateTimeEnd?: string
}

export type type_flow_history_edit = {
    id?: string
    applyId?: string
    templateId?: string
    applyUserId?: string
    applyUserNickname?: string
    approverId?: string
    approverNickname?: string
    nodeId?: string
    nodeType?: string
    nodeLabel?: string
    formValue?: string
    passStatus?: number
    passRemark?: string
}

export type type_flow_history_pass = {
    applyId: string
    nextNodeAdminId?: string
    passRemark?: string
}

export type type_flow_history_back = {
    applyId: string
    historyId: string
    remark?: string
}

export type type_flow_history_next_node = {
    applyId: string
}

export type type_flow_history_get_approver = {
    applyId: string
}

export function flow_history_list(params?: type_flow_history_query) {
    return request.get<Pages<type_flow_history>>({ url: '/flow/flow_history/list', params })
}

export function flow_history_list_all(params?: type_flow_history_query) {
    return request.get<type_flow_history[]>({ url: '/flow/flow_history/list_all', params })
}

export function flow_history_detail(id: string) {
    return request.get<type_flow_history>({ url: '/flow/flow_history/detail', params: { id } })
}

export function flow_history_add(data: type_flow_history_edit) {
    return request.post<null>({ url: '/flow/flow_history/add', data })
}

export function flow_history_edit(data: type_flow_history_edit) {
    return request.post<null>({ url: '/flow/flow_history/edit', data })
}

export function flow_history_delete(id: string) {
    return request.post<null>({ url: '/flow/flow_history/del', data: { id } })
}

export function flow_history_done_hidden(id: string) {
    return request.post<null>({ url: '/flow/flow_history/done_hidden', data: { id } })
}

export function flow_history_next_node(data: type_flow_history_next_node) {
    return request.post<any>({ url: '/flow/flow_history/next_node', data })
}

export function flow_history_get_approver(data: type_flow_history_get_approver) {
    return request.post<any>({ url: '/flow/flow_history/get_approver', data })
}

export function flow_history_pass(data: type_flow_history_pass) {
    return request.post<null>({ url: '/flow/flow_history/pass', data })
}

export function flow_history_back(data: type_flow_history_back) {
    return request.post<null>({ url: '/flow/flow_history/back', data })
}
