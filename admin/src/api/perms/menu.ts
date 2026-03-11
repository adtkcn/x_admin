import request from '@/utils/request'

// 菜单详情参数
export type type_system_menu_detail = {
    id: string
}

// 菜单添加参数
export type type_system_menu_add = {
    pid?: string
    menuType: string
    menuName: string
    menuIcon?: string
    menuSort?: number
    perms?: string
    paths?: string
    component?: string
    selected?: string
    params?: string
    isCache?: number
    isShow?: number
    isDisable?: number
}

// 菜单编辑参数
export type type_system_menu_edit = {
    id: string
    pid?: string
    menuType: string
    menuName: string
    menuIcon?: string
    menuSort?: number
    perms?: string
    paths?: string
    component?: string
    selected?: string
    params?: string
    isCache?: number
    isShow?: number
    isDisable?: number
}

// 菜单删除参数
export type type_system_menu_del = {
    id: string
}

// 菜单返回信息
export type type_system_menu_resp = {
    id: string
    pid: string
    menuType: string
    menuName: string
    menuIcon: string
    menuSort: number
    perms: string
    paths: string
    component: string
    selected: string
    params: string
    isCache: number
    isShow: number
    isDisable: number
    createTime: string
    updateTime: string
    children?: type_system_menu_resp[]
}

// 菜单列表
export function menuLists(params?: any) {
    return request.get<type_system_menu_resp[]>({ url: '/system/menu/list', params })
}

// 菜单详情
export function menuDetail(params: type_system_menu_detail) {
    return request.get<type_system_menu_resp>({ url: '/system/menu/detail', params })
}

// 添加菜单
export function menuAdd(data: type_system_menu_add) {
    return request.post({ url: '/system/menu/add', data })
}

// 编辑菜单
export function menuEdit(data: type_system_menu_edit) {
    return request.post({ url: '/system/menu/edit', data })
}

// 菜单删除
export function menuDelete(data: type_system_menu_del) {
    return request.post({ url: '/system/menu/del', data })
}
