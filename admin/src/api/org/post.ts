import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 岗位列表参数
export type type_system_post_list = {
    code?: string
    name?: string
    is_stop?: number
}

// 岗位详情参数
export type type_system_post_detail = {
    id: string
}

// 岗位编辑参数
export type type_system_post_edit = {
    id?: string
    code?: string
    name?: string
    remarks?: string
    is_stop?: number
    sort?: number
}

// 岗位删除参数
export type type_system_post_del = {
    id: string
}

// 岗位返回信息
export type type_system_post_resp = {
    id: string
    code: string
    name: string
    remarks: string
    sort: number
    is_stop: number
    create_time: string
    update_time: string
}

// 岗位详情
export function postDetail(params: type_system_post_detail) {
    return request.get<type_system_post_resp>({ url: '/system/post/detail', params })
}

// 岗位列表
export function postLists(params?: type_system_post_list) {
    return request.get<Pages<type_system_post_resp>>({ url: '/system/post/list', params })
}

// 岗位全部列表
export function postAll() {
    return request.get<type_system_post_resp[]>({ url: '/system/post/all' })
}

// 添加岗位
export function postAdd(data: type_system_post_edit) {
    return request.post({ url: '/system/post/add', data })
}

// 编辑岗位
export function postEdit(data: type_system_post_edit) {
    return request.post({ url: '/system/post/edit', data })
}

// 删除岗位
export function postDelete(data: type_system_post_del) {
    return request.post({ url: '/system/post/del', data })
}
