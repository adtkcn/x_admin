import request from '@/utils/request'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import config from '@/config'
import type { Pages } from '@/utils/request'

export type type_system_admin_list = {
    username?: string
    nickname?: string
    roleId?: string
}

export type type_system_admin_detail = {
    id: string
}

export type type_system_admin_add = {
    deptId: string
    postId: string
    roleIds: string[]
    username: string
    nickname: string
    password: string
    avatar: string
    sort: number
    isDisable: number
}

export type type_system_admin_edit = {
    id: string
    deptId: string
    postId: string
    roleIds: string[]
    username: string
    nickname: string
    password: string
    avatar: string
    sort: number
    isDisable: number
}

export type type_system_admin_del = {
    id: string
}

export type type_system_admin_disable = {
    id: string
}

export type type_system_admin_resp = {
    id: string
    username: string
    nickname: string
    avatar: string
    role: string
    deptId: string
    dept: string
    postId: string
    post: string
    roleIds: string[]
    isDisable: number
    lastLoginIp: string
    lastLoginTime: string
    createTime: string
    updateTime: string
}

// 管理员列表
export function adminLists(params: type_system_admin_list) {
    return request.get<Pages<type_system_admin_resp>>({ url: '/system/admin/list', params })
}

// 管理员列表
export function adminListAll(params: type_system_admin_list) {
    return request.get<type_system_admin_resp[]>({ url: '/system/admin/listAll', params })
}

// 管理员详情
export function adminDetail(params: type_system_admin_detail) {
    return request.get<type_system_admin_resp>({ url: '/system/admin/detail', params })
}

// 管理员添加
export function adminAdd(data: type_system_admin_add) {
    return request.post({ url: '/system/admin/add', data })
}

// 管理员编辑
export function adminEdit(data: type_system_admin_edit) {
    return request.post({ url: '/system/admin/edit', data })
}

// 管理员删除
export function adminDelete(data: type_system_admin_del) {
    return request.post({ url: '/system/admin/del', data })
}

// 管理员状态切换
export function adminStatus(data: type_system_admin_disable) {
    return request.post({ url: '/system/admin/disable', data })
}

// 部门下的管理员
export function adminListByDeptId(params: { deptId: string }) {
    return request.get<type_system_admin_resp[]>({ url: '/system/admin/ListByDeptId', params })
}

// 导入
export const adminImportFile = '/system/admin/import_file'

// 导出
export function adminExportFile(params: type_system_admin_list) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/system/admin/export_file?token=${getToken()}&` +
        queryString.stringify(params))
}
