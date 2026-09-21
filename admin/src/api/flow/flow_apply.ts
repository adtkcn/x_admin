import request from '@/utils/request'
import type { Pages } from '@/utils/request'

export type type_flow_apply = {
    id: string
    template_id: string | null
    apply_user_id: string | null
    apply_user_nickname: string | null
    flow_name: string | null
    flow_group: number | null
    flow_remark: string | null
    flow_form_data: string | null
    flow_process_data: string | null
    flow_process_data_list: string | null
    form_value: string | null
    status: number | null // 状态：1待提交，2审批中，3审批完成，4审批失败
    create_time: string | null
    update_time: string | null
}
// 查询
export type type_flow_apply_query = {
    template_id?: string
    apply_user_id?: string
    apply_user_nickname?: string
    flow_name?: string
    flow_group?: number
    flow_remark?: string
    flow_form_data?: string
    flow_process_data?: string
    flow_process_data_list?: string
    form_value?: string
    status?: number
    create_time_start?: string
    create_time_end?: string
    update_time_start?: string
    update_time_end?: string
}
// 添加编辑
export type type_flow_apply_edit = {
    id?: string
    template_id?: string | null
    apply_user_id?: string | null
    apply_user_nickname?: string | null
    flow_name?: string | null
    form_value?: string | null
    status?: number | null
}

// 申请流程列表
export function flow_apply_lists(params?: type_flow_apply_query) {
    return request.get<Pages<type_flow_apply>>({ url: '/flow/flow_apply/list', params })
}
// 申请流程列表-所有
export function flow_apply_list_all(params?: type_flow_apply_query) {
    return request.get<type_flow_apply[]>({ url: '/flow/flow_apply/list_all', params })
}

// 申请流程错误用户列表
export function flow_apply_error_users(params: { id: string }) {
    return request.get<any[]>({ url: '/flow/flow_apply/error_users', params })
}

// 申请流程详情
export function flow_apply_detail(params: { id: string }) {
    return request.get<type_flow_apply>({ url: '/flow/flow_apply/detail', params: params })
}

// 申请流程新增
export function flow_apply_add(data: type_flow_apply_edit) {
    return request.post<null>({ url: '/flow/flow_apply/add', data })
}

// 申请流程编辑
export function flow_apply_edit(data: type_flow_apply_edit) {
    return request.post<null>({ url: '/flow/flow_apply/edit', data })
}

// 申请流程删除
export function flow_apply_delete(id: string) {
    return request.post<null>({ url: '/flow/flow_apply/del', data: { id } })
}
