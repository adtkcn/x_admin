import request from '@/utils/request'

// 角色列表
export function roleLists(params: any) {
    return request.get({ url: '/system/role/list', params })
}

// 角色列表
export function roleAll(params?: any) {
    return request.get({ url: '/system/role/all', params })
}

// 角色列表
export function roleDetail(params: any) {
    return request.get({ url: '/system/role/detail', params })
}

// 添加角色
export function roleAdd(data: any) {
    return request.post({ url: '/system/role/add', data })
}
// 编辑角色
export function roleEdit(data: any) {
    return request.post({ url: '/system/role/edit', data })
}
// 删除角色
export function roleDelete(data: any) {
    return request.post({ url: '/system/role/del', data })
}
