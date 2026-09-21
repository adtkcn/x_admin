import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_monitor_client = {
    id?: string
    project_key?: string
    client_id?: string
    user_id?: string
    os?: string
    browser?: string
    country?: string
    province?: string
    city?: string
    operator?: string
    ip?: string

    ua?: string
    create_time?: string

    width?: string
    height?: string
}
// 查询
export type type_monitor_client_query = {
    project_key?: string
    os?: string
    browser?: string
    ua?: string
    create_time_start?: string
    create_time_end?: string
}
// 添加编辑
export type type_monitor_client_edit = {
    id?: string
    project_key?: string
    os?: string
    browser?: string
    ua?: string
}

// 监控-客户端信息列表
export function monitor_client_list(params?: type_monitor_client_query) {
    return request.get<Pages<type_monitor_client>>({
        url: '/monitor_client/list',
        params: clearEmpty(params)
    })
}
// 监控-客户端信息列表-所有
export function monitor_client_list_all(params?: type_monitor_client_query) {
    return request.get<type_monitor_client[]>({
        url: '/monitor_client/list_all',
        params: clearEmpty(params)
    })
}

// 监控-客户端信息详情
export function monitor_client_detail(id: string) {
    return request.get<type_monitor_client>({ url: '/monitor_client/detail', params: { id } })
}

export function monitor_client_errorUsers(id: string) {
    return request.get<type_monitor_client[]>({ url: '/monitor_client/errorUsers', params: { id } })
}

// 监控-客户端信息新增
export function monitor_client_add(data: type_monitor_client_edit) {
    return request.post<null>({ url: '/monitor_client/add', data })
}

// 监控-客户端信息编辑
export function monitor_client_edit(data: type_monitor_client_edit) {
    return request.post<null>({ url: '/monitor_client/edit', data })
}

// 监控-客户端信息删除
export function monitor_client_delete(id: string) {
    return request.post<null>({ url: '/monitor_client/del', data: { id } })
}
// 监控-客户端信息删除-批量
export function monitor_client_delete_batch(data: { ids: string }) {
    return request.post<null>({ url: '/monitor_client/del_batch', data })
}

// 监控-客户端信息导入
export const monitor_client_import_file = '/monitor_client/import_file'

// 监控-客户端信息导出
export function monitor_client_export_file(params: type_monitor_client_query) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/monitor_client/export_file?token=${getToken()}&` +
        queryString.stringify(clearEmpty(params)))
}
