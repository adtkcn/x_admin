import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 角色详情参数
export type type_system_role_detail = {
    id: string
}

// 角色添加参数
export type type_system_role_add = {
    name: string
    sort?: number
    is_disable?: number
    remark?: string
    menuIds?: string
}

// 角色编辑参数
export type type_system_role_edit = {
    id: string
    name: string
    sort?: number
    is_disable?: number
    remark?: string
    menuIds?: string
}

// 角色删除参数
export type type_system_role_del = {
    id: string
}

// 角色返回信息
export type type_system_role_resp = {
    id: string
    name: string
    remark: string
    menus: string[]
    member: number
    sort: number
    is_disable: number
    create_time: string
    update_time: string
}

// 角色简单返回信息
export type type_system_role_simple_resp = {
    id: string
    name: string
    create_time: string
    update_time: string
}

// 角色列表
export function roleLists(params?: any) {
    return request.get<Pages<type_system_role_resp>>({ url: '/system/role/list', params })
}

// 角色列表
export function roleAll(params?: any) {
    return request.get<type_system_role_simple_resp[]>({ url: '/system/role/all', params })
}

// 角色详情
export function roleDetail(params: type_system_role_detail) {
    return request.get<type_system_role_resp>({ url: '/system/role/detail', params })
}

// 添加角色
export function roleAdd(data: type_system_role_add) {
    return request.post({ url: '/system/role/add', data })
}

// 编辑角色
export function roleEdit(data: type_system_role_edit) {
    return request.post({ url: '/system/role/edit', data })
}

// 删除角色
export function roleDelete(data: type_system_role_del) {
    return request.post({ url: '/system/role/del', data })
}
