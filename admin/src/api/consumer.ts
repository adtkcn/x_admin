import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 用户列表参数
export type type_user_list = {
    username?: string
    nickname?: string
    mobile?: string
    status?: number
    createTimeStart?: string
    createTimeEnd?: string
}

// 用户详情参数
export type type_user_detail = {
    id: string
}

// 用户编辑参数
export type type_user_edit = {
    id: string
    nickname?: string
    mobile?: string
    email?: string
    status?: number
}

// 用户返回信息
export type type_user_resp = {
    id: string
    username: string
    nickname: string
    mobile: string
    email: string
    avatar: string
    status: number
    createTime: string
    updateTime: string
    lastLoginTime: string
    lastLoginIp: string
}

// 用户列表
export function getUserList(params: type_user_list) {
    return request.get<Pages<type_user_resp>>({ url: '/user/list', params })
}

// 用户详情
export function getUserDetail(params: type_user_detail) {
    return request.get<type_user_resp>({ url: '/user/detail', params })
}

// 用户编辑
export function userEdit(data: type_user_edit) {
    return request.post({ url: '/user/edit', data })
}
