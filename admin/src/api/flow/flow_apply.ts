import request from '@/utils/request'
import type { Pages } from '@/utils/request'

export type type_flow_apply = {
    id?: number
    templateId?: number
    applyUserId?: number
    applyUserNickname?: string
    flowName?: string
    flowGroup?: number
    flowRemark?: string
    flowFormData?: string
    flowProcessData?: string
    flowProcessDataList?: string
    formValue?: string
    status?: number
    isDelete?: number
    createTime?: string
    updateTime?: string
    deleteTime?: string
}
// 查询
export type type_flow_apply_query = {
    templateId?: number
    applyUserId?: number
    applyUserNickname?: string
    flowName?: string
    flowGroup?: number
    flowRemark?: string
    flowFormData?: string
    flowProcessData?: string
    flowProcessDataList?: string
    formValue?: string
    status?: number
    createTimeStart?: string
    createTimeEnd?: string
    updateTimeStart?: string
    updateTimeEnd?: string
}
// 添加编辑
export type type_flow_apply_edit = {
    id?: number
    templateId?: number
    applyUserId?: number
    applyUserNickname?: string
    flowName?: string
    flowGroup?: number
    flowRemark?: string
    flowFormData?: string
    flowProcessData?: string
    flowProcessDataList?: string
    formValue?: string
    status?: number
}

// 申请流程列表
export function flow_apply_lists(params?: type_flow_apply_query) {
    return request.get<Pages<type_flow_apply>>({ url: '/flow/flow_apply/list', params })
}
// 申请流程列表-所有
export function flow_apply_list_all(params?: type_flow_apply_query) {
    return request.get<type_flow_apply[]>({ url: '/flow/flow_apply/listAll', params })
}

// 申请流程详情
export function flow_apply_detail(params: Record<string, any>) {
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
export function flow_apply_delete(id: number | string) {
    return request.post<null>({ url: '/flow/flow_apply/del', data: { id } })
}
