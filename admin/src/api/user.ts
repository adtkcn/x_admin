import config from '@/config'
import request from '@/utils/request'
import type { type_system_menu_resp } from '@/api/perms/menu'
// 登录参数
export type type_system_login = {
    email: string
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
    email?: string
    emailCode?: string
    password?: string
    currPassword?: string
}

// 用户信息响应
export type type_system_admin_self = {
    user: {
        id: string
        email: string
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
    return request.get<type_system_menu_resp[]>({ url: '/system/menu/route' })
}

// 编辑管理员信息
export function setUserInfo(data: type_system_admin_update) {
    return request.post({ url: '/system/admin/upInfo', data })
}

// 发送邮箱绑定验证码
export function sendEmailCode(data: { email: string }) {
    return request.post({ url: '/system/admin/sendEmailCode', data })
}

// 忘记密码-发送验证码
export function forgotPwdSendCode(data: { email: string }) {
    return request.post({ url: '/system/forgot-pwd/send-code', data })
}

// 忘记密码-重置密码
export function forgotPwdReset(data: { email: string; code: string; password: string }) {
    return request.post({ url: '/system/forgot-pwd/reset', data })
}
