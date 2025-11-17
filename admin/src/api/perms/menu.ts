import request from '@/utils/request'

// import type { Pages } from '@/utils/request'
export interface SystemAuthMenuResp {
    id?: number
    pid: number
    menuType: string
    menuName: string
    menuIcon: string
    menuSort: number
    perms: string
    paths: string
    component: string
    selected: boolean
    params: string
    isCache: boolean
    isShow: boolean
    isDisable: boolean
    createTime: string
    updateTime: string
    children?: SystemAuthMenuResp[]
}
// 菜单列表
export function menuLists(params?: Record<string, any>) {
    return request.get<SystemAuthMenuResp[]>({ url: '/system/menu/list', params })
}
// 菜单
export function menuDetail(params: Record<string, any>) {
    return request.get<SystemAuthMenuResp>({ url: '/system/menu/detail', params })
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
