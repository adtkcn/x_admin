import request from '@/utils/request'
import type { Pages } from '@/utils/request'

export type type_flow_template = {
    id?: number
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
// 查询
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
// 添加编辑
export type type_flow_template_edit = {
    id?: number
    flowName?: string
    flowGroup?: number
    flowRemark?: string
    flowFormData?: string
    flowProcessData?: string
    flowProcessDataList?: string
}

// 流程模板列表
export function flow_template_lists(params?: type_flow_template_query) {
    return request.get<Pages<type_flow_template>>({ url: '/flow/flow_template/list', params })
}
// 流程模板列表-所有
export function flow_template_lists_all(params?: type_flow_template_query) {
    return request.get<type_flow_template[]>({ url: '/flow/flow_template/listAll', params })
}

// 流程模板详情
export function flow_template_detail(id: number | string) {
    return request.get<type_flow_template>({ url: '/flow/flow_template/detail', params: { id } })
}

// 流程模板新增
export function flow_template_add(data: type_flow_template_edit) {
    return request.post<null>({ url: '/flow/flow_template/add', data })
}

// 流程模板编辑
export function flow_template_edit(data: type_flow_template_edit) {
    return request.post<null>({ url: '/flow/flow_template/edit', data })
}

// 流程模板删除
export function flow_template_delete(id: number | string) {
    return request.post<null>({ url: '/flow/flow_template/del', data: { id } })
}
