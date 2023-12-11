import request from '@/utils/request'

// 流程历史列表
export function flow_history_list(params?: Record<string, any>) {
    return request.get({ url: '/flow_history/list', params })
}
// 流程历史列表-所有
export function flow_history_list_all(params?: Record<string, any>) {
    return request.get({ url: '/flow_history/listAll', params })
}

// 流程历史详情
export function flow_history_detail(params: Record<string, any>) {
    return request.get({ url: '/flow_history/detail', params })
}

// 流程历史新增
export function flow_history_add(params: Record<string, any>) {
    return request.post({ url: '/flow_history/add', params })
}

// 流程历史编辑
export function flow_history_edit(params: Record<string, any>) {
    return request.post({ url: '/flow_history/edit', params })
}

// 流程历史删除
export function flow_history_delete(params: Record<string, any>) {
    return request.post({ url: '/flow_history/del', params })
}
