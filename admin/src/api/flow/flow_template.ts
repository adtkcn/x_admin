import request from '@/utils/request'
import type { Pages } from '@/utils/request'

export type type_flow_template = {
    id?: string
    flowName?: string
    flowGroup?: number
    flowRemark?: string
    flowFormData?: string
    flowProcessData?: string
    flowProcessDataList?: string
    isDelete?: number
    createTime?: string
    updateTime?: string
    deleteTime?: string
}

export type type_flow_template_query = {
    flowName?: string
    flowGroup?: number
    flowRemark?: string
    flowFormData?: string
    flowProcessData?: string
    flowProcessDataList?: string
    createTimeStart?: string
    createTimeEnd?: string
    updateTimeStart?: string
    updateTimeEnd?: string
}

export type type_flow_template_edit = {
    id?: string
    flowName?: string
    flowGroup?: number
    flowRemark?: string
    flowFormData?: string
    flowProcessData?: string
    flowProcessDataList?: string
}

export function flow_template_lists(params?: type_flow_template_query) {
    return request.get<Pages<type_flow_template>>({ url: '/flow/flow_template/list', params })
}

export function flow_template_lists_all() {
    return request.get<type_flow_template[]>({ url: '/flow/flow_template/list_all' })
}

export function flow_template_detail(id: string) {
    return request.get<type_flow_template>({ url: '/flow/flow_template/detail', params: { id } })
}

export function flow_template_add(data: type_flow_template_edit) {
    return request.post<null>({ url: '/flow/flow_template/add', data })
}

export function flow_template_edit(data: type_flow_template_edit) {
    return request.post<null>({ url: '/flow/flow_template/edit', data })
}

export function flow_template_delete(id: string) {
    return request.post<null>({ url: '/flow/flow_template/del', data: { id } })
}
