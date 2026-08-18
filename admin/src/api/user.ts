import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// ===================== 用户 Users =====================
export type type_user_list = {
    keyword?: string
    status?: string
    create_time_start?: string
    create_time_end?: string
}
export type type_user_edit = {
    id?: string
    nickname?: string
    avatar?: string
    phone?: string
    phone_code?: string
    status?: number
}
export type type_user_resp = {
    id: string
    email: string
    nickname: string
    avatar: string
    phone: string
    phone_code: string
    status: number
    last_login_ip: string
    last_login_time: string
    create_time: string
}
export function userLists(params: type_user_list) {
    return request.get<Pages<type_user_resp>>({ url: '/user/list', params })
}
export function userDetail(params: { id: string }) {
    return request.get<type_user_resp>({ url: '/user/detail', params })
}
export function userEdit(data: type_user_edit) {
    return request.post({ url: '/user/edit', data })
}
export function userDisable(data: { id: string; status: number }) {
    return request.post({ url: '/user/disable', data })
}
export function userKick(data: { id: string }) {
    return request.post({ url: '/user/kick', data })
}
