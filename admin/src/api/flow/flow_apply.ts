import request from '@/utils/request'

// 申请流程列表
export function flow_apply_lists(params?: Record<string, any>) {
    return request.get({ url: '/flow/flow_apply/list', params })
}

// 申请流程详情
export function flow_apply_detail(params: Record<string, any>) {
    return request.get({ url: '/flow/flow_apply/detail', params })
}

// 申请流程新增
export function flow_apply_add(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_apply/add', params })
}

// 申请流程编辑
export function flow_apply_edit(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_apply/edit', params })
}

// 申请流程删除
export function flow_apply_delete(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_apply/del', params })
}
