import request from '@/utils/request/index'

// 流程历史列表
export function flow_history_list(params?: Record<string, any>) {
    return request.get({ url: '/flow/flow_history/list', params })
}
// 流程历史列表-所有
export function flow_history_list_all(params?: Record<string, any>) {
    return request.get({ url: '/flow/flow_history/listAll', params })
}

// 流程历史详情
export function flow_history_detail(params: Record<string, any>) {
    return request.get({ url: '/flow/flow_history/detail', params })
}

// 流程历史新增
export function flow_history_add(data: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/add', data })
}

// 流程历史编辑
export function flow_history_edit(data: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/edit', data })
}

// 流程历史删除
export function flow_history_delete(data: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/del', data })
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
