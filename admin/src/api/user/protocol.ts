import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_user_protocol = {
    Id?: number
    Title?: string
    Content?: string
    Sort?: number
    IsDelete?: number
    CreateTime?: string
    UpdateTime?: string
    DeleteTime?: string
}
// 查询
export type type_user_protocol_query = {
    Title?: string
    Content?: string
    Sort?: number
    CreateTimeStart?: string
    CreateTimeEnd?: string
    UpdateTimeStart?: string
    UpdateTimeEnd?: string
}
// 添加编辑
export type type_user_protocol_edit = {
    Id?: number
    Title?: string
    Content?: string
    Sort?: number
}

// 用户协议列表
export function user_protocol_list(params?: type_user_protocol_query) {
    return request.get<Pages<type_user_protocol>>({ url: '/user_protocol/list', params: clearEmpty(params) })
}
// 用户协议列表-所有
export function user_protocol_list_all(params?: type_user_protocol_query) {
    return request.get<type_user_protocol[]>({ url: '/user_protocol/listAll', params: clearEmpty(params) })
}

// 用户协议详情
export function user_protocol_detail(Id: number | string) {
    return request.get<type_user_protocol>({ url: '/user_protocol/detail', params: { Id } })
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
export function user_protocol_delete(Id: number | string) {
    return request.post<null>({ url: '/user_protocol/del', data: { Id } })
}

// 用户协议导入
export const user_protocol_import_file = '/user_protocol/ImportFile'

// 用户协议导出
export function user_protocol_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/user_protocol/ExportFile?token=${getToken()}&` + queryString.stringify(clearEmpty(params)))
}
