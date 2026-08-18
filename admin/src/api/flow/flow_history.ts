import request from '@/utils/request/index'
import type { Pages } from '@/utils/request'

export type type_flow_history = {
    id: string
    apply_id: string | null
    template_id: string | null
    apply_user_id: string | null
    apply_user_nickname: string | null
    approver_id: string | null
    approver_nickname: string | null
    node_id: string | null
    node_type: string | null
    node_label: string | null
    form_value: string | null
    pass_status: number | null // 通过状态：1待处理，2通过，3拒绝
    pass_remark: string | null
    is_show: number | null // 是否显示：0隐藏，1显示
    create_time: string | null
    update_time: string | null
}

export type type_flow_history_query = {
    apply_id?: string
    template_id?: string
    apply_user_id?: string
    apply_user_nickname?: string
    approver_id?: string
    approver_nickname?: string
    node_id?: string
    node_type?: string
    node_label?: string
    form_value?: string
    pass_status?: number
    pass_remark?: string
    is_show?: number
    create_time_start?: string
    create_time_end?: string
    update_time_start?: string
    update_time_end?: string
}

export type type_flow_history_edit = {
    id?: string
    apply_id?: string
    template_id?: string
    apply_user_id?: string
    apply_user_nickname?: string
    approver_id?: string
    approver_nickname?: string
    node_id?: string
    node_type?: string
    node_label?: string
    form_value?: string
    pass_status?: number
    pass_remark?: string
}

export type type_flow_history_pass = {
    apply_id: string
    next_node_admin_id?: string
    pass_remark?: string
}

export type type_flow_history_back = {
    apply_id: string
    history_id: string
    remark?: string
}

export type type_flow_history_next_node = {
    apply_id: string
}

// 流程节点树返回（对应后端 FlowTree）
export type type_flow_tree = {
    id: string
    pid: string
    label: string
    type: string // bpmn:startEvent / bpmn:userTask / ...
    props: type_flow_tree_props
    children?: type_flow_tree[]
}

export type type_flow_tree_props = {
    start_event?: {
        field_auth: Record<string, number> // 表单项权限：1读写 2只读 3隐藏
    }
    user_task?: {
        field_auth: Record<string, number>
        user_type: number // 1指定部门、岗位 2用户部门负责人 3指定审批人
        user_id: string
        dept_id: string
        post_id: string
    }
    notify_task?: {
        service_type: string // site=站内消息 email=邮件 webhook=回调
        service_content: string
        receiver_id: string[]
        email_to: string[]
        webhook_url: string
    }
    exclusive_gateway?: {
        gateway: {
            id: string
            condition: string // == != >= <= include
            value: string
        }[]
    }
    end_event?: Record<string, never>
}

export type type_flow_history_get_approver = {
    apply_id: string
}

// 节点可审批用户返回（管理员信息）
export type type_flow_history_get_approver_resp = {
    id: string
    email: string
    nickname: string
    avatar: string
    dept_id: string
    dept: string
    post_id: string
    post: string
    role_ids: string[]
    role: string
    is_disable: number
    last_login_ip: string
    last_login_time: string
    create_time: string
    update_time: string
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
    return request.post<type_flow_tree[]>({ url: '/flow/flow_history/next_node', data })
}

export function flow_history_get_approver(data: type_flow_history_get_approver) {
    return request.post<type_flow_history_get_approver_resp[]>({ url: '/flow/flow_history/get_approver', data })
}

export function flow_history_pass(data: type_flow_history_pass) {
    return request.post<null>({ url: '/flow/flow_history/pass', data })
}

export function flow_history_back(data: type_flow_history_back) {
    return request.post<null>({ url: '/flow/flow_history/back', data })
}
