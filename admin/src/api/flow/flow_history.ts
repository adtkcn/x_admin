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
export function flow_history_add(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/add', params })
}

// 流程历史编辑
export function flow_history_edit(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/edit', params })
}

// 流程历史删除
export function flow_history_delete(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/del', params })
}

// 获取下一个审批节点，中间可能有系统任务和结束节点被跳过
export function flow_history_next_node(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/next_node', params })
}

// 获取下一个审批节点，中间可能有系统任务和结束节点被跳过
export function flow_history_get_approver(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/get_approver', params })
}
export function flow_history_pass(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/pass', params })
}

export function flow_history_back(params: Record<string, any>) {
    return request.post({ url: '/flow/flow_history/back', params })
}
