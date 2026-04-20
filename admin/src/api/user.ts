import config from '@/config'
import request from '@/utils/request'

// 登录参数
export type type_system_login = {
    username: string
    password: string
}

// 登录响应
export type type_system_login_resp = {
    token: string
}

// 用户信息更新参数
export type type_system_admin_update = {
    nickname: string
    avatar?: string
    password?: string
    currPassword?: string
}

// 用户信息响应
export type type_system_admin_self = {
    user: {
        id: string
        username: string
        nickname: string
        avatar: string
        role: string
        dept: string
        isDisable: number
        lastLoginIp: string
        lastLoginTime: string
        createTime: string
        updateTime: string
    }
    permissions: string[]
}

// 登录
export function login(data: type_system_login) {
    return request.post<type_system_login_resp>({
        url: '/system/login',
        data: { ...data, terminal: config.terminal }
    })
}

// 退出登录
export function logout() {
    return request.post({ url: '/system/logout' })
}

// 用户信息
export function getUserInfo() {
    return request.get<type_system_admin_self>({ url: '/system/admin/self' })
}

// 菜单路由
export function getMenu() {
    return request.get({ url: '/system/menu/route' })
}

// 编辑管理员信息
export function setUserInfo(data: type_system_admin_update) {
    return request.post({ url: '/system/admin/upInfo', data })
}
