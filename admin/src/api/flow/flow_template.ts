import request from '@/utils/request'

// 流程模板列表
export function flow_template_lists(params?: Record<string, any>) {
    return request.get({ url: '/flow/flow_template/list', params })
}

// 流程模板列表-所有
export function flow_template_lists_all(params?: Record<string, any>) {
    return request.get({ url: '/flow/flow_template/listAll', params })
}

// 流程模板详情
export function flow_template_detail(params: Record<string, any>) {
    return request.get({ url: '/flow/flow_template/detail', params })
}

// 流程模板新增
export function flow_template_add(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_template/add', params })
}

// 流程模板编辑
export function flow_template_edit(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_template/edit', params })
}

// 流程模板删除
export function flow_template_delete(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_template/del', params })
}
