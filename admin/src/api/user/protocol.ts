import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_user_protocol = {
    id?: string
    tag?: string
    version?: number
    title?: string
    content?: string

    create_time?: string
    update_time?: string
    created_by?: string
}
// 查询
export type type_user_protocol_query = {
    title?: string
    content?: string
    version?: number

    create_time_start?: string
    create_time_end?: string
    update_time_start?: string
    update_time_end?: string
}
// 添加编辑
export type type_user_protocol_edit = {
    id?: string
    tag?: string
    version?: number
    title?: string
    content?: string
}

// 用户协议列表
export function user_protocol_list(params?: type_user_protocol_query) {
    return request.get<Pages<type_user_protocol>>({
        url: '/user_protocol/list',
        params: clearEmpty(params)
    })
}
// 用户协议列表-所有
export function user_protocol_list_all(params?: type_user_protocol_query) {
    return request.get<type_user_protocol[]>({
        url: '/user_protocol/list_all',
        params: clearEmpty(params)
    })
}

// 用户协议详情
export function user_protocol_detail(id: string) {
    return request.get<type_user_protocol>({ url: '/user_protocol/detail', params: { id } })
}

// 用户协议新增
export function user_protocol_add(data: type_user_protocol_edit) {
    return request.post<null>({ url: '/user_protocol/add', data })
}

// 用户协议编辑
export function user_protocol_edit(data: type_user_protocol_edit) {
    return request.post<null>({ url: '/user_protocol/edit', data })
}

// 用户协议删除
export function user_protocol_delete(id: string) {
    return request.post<null>({ url: '/user_protocol/del', data: { id } })
}
// 用户协议删除-批量
export function user_protocol_delete_batch(data: { ids: string }) {
    return request.post<null>({ url: '/user_protocol/del_batch', data })
}

// 用户协议导入
export const user_protocol_import_file = '/user_protocol/import_file'

// 用户协议导出
export function user_protocol_export_file(params: type_user_protocol_query) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/user_protocol/export_file?token=${getToken()}&` +
        queryString.stringify(clearEmpty(params)))
}
