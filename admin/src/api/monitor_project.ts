import request from '@/utils/request'

// 错误项目列表
export function monitor_project_list(params?: Record<string, any>) {
    return request.get({ url: '/monitor_project/list', params })
}
// 错误项目列表-所有
export function monitor_project_list_all(params?: Record<string, any>) {
    return request.get({ url: '/monitor_project/listAll', params })
}

// 错误项目详情
export function monitor_project_detail(params: Record<string, any>) {
    return request.get({ url: '/monitor_project/detail', params })
}

// 错误项目新增
export function monitor_project_add(params: Record<string, any>) {
    return request.post({ url: '/monitor_project/add', params })
}

// 错误项目编辑
export function monitor_project_edit(params: Record<string, any>) {
    return request.post({ url: '/monitor_project/edit', params })
}

// 错误项目删除
export function monitor_project_delete(params: Record<string, any>) {
    return request.post({ url: '/monitor_project/del', params })
}
