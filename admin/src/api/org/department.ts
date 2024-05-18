import request from '@/utils/request'

// 部门列表
export function deptLists(params?: any) {
    return request.get({ url: '/system/dept/list', params })
}
// 部门列表-全部
export function deptAll(params?: any) {
    return request.get({ url: '/system/dept/all', params })
}
// 部门详情
export function deptDetail(params?: any) {
    return request.get({ url: '/system/dept/detail', params })
}

// 添加部门
export function deptAdd(data: any) {
    return request.post({ url: '/system/dept/add', data })
}

// 编辑部门
export function deptEdit(data: any) {
    return request.post({ url: '/system/dept/edit', data })
}

// 删除部门
export function deptDelete(data: any) {
    return request.post({ url: '/system/dept/del', data })
}
