import request from '@/utils/request'

// 菜单列表
export function menuLists(params?: Record<string, any>) {
    return request.get({ url: '/system/menu/list', params })
}
// 菜单
export function menuDetail(params: Record<string, any>) {
    return request.get({ url: '/system/menu/detail', params })
}

// 添加菜单
export function menuAdd(data: Record<string, any>) {
    return request.post({ url: '/system/menu/add', data })
}

// 编辑菜单
export function menuEdit(data: Record<string, any>) {
    return request.post({ url: '/system/menu/edit', data })
}

// 菜单删除
export function menuDelete(data: Record<string, any>) {
    return request.post({ url: '/system/menu/del', data })
}
