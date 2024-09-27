import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_monitor_client = {
    Id?: number
    ProjectKey?: string
    ClientId?: string
    UserId?: string
    Os?: string
    Browser?: string
    City?: string
    Width?: number
    Height?: number
    Ua?: string
    CreateTime?: string
    ClientTime?: string
}
// 查询
export type type_monitor_client_query = {
    ProjectKey?: string
    ClientId?: string
    UserId?: string
    Os?: string
    Browser?: string
    City?: string
    Width?: number
    Height?: number
    Ua?: string
    CreateTimeStart?: string
    CreateTimeEnd?: string
    ClientTimeStart?: string
    ClientTimeEnd?: string
}
// 添加编辑
export type type_monitor_client_edit = {
    Id?: number
    ProjectKey?: string
    ClientId?: string
    UserId?: string
    Os?: string
    Browser?: string
    City?: string
    Width?: number
    Height?: number
    Ua?: string
    ClientTime?: string
}

// 监控-客户端信息列表
export function monitor_client_list(params?: type_monitor_client_query) {
    return request.get<Pages<type_monitor_client>>({ url: '/monitor_client/list', params: clearEmpty(params) })
}
// 监控-客户端信息列表-所有
export function monitor_client_list_all(params?: type_monitor_client_query) {
    return request.get<type_monitor_client[]>({ url: '/monitor_client/listAll', params: clearEmpty(params) })
}

// 监控-客户端信息详情
export function monitor_client_detail(Id: number | string) {
    return request.get<type_monitor_client>({ url: '/monitor_client/detail', params: { Id } })
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
export function monitor_client_delete(Id: number | string) {
    return request.post<null>({ url: '/monitor_client/del', data: { Id } })
}
// 监控-客户端信息删除-批量
export function monitor_client_delete_batch(data: { Ids: string }) {
    return request.post<null>({ url: '/monitor_client/delBatch', data })
}

// 监控-客户端信息导入
export const monitor_client_import_file = '/monitor_client/ImportFile'

// 监控-客户端信息导出
export function monitor_client_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/monitor_client/ExportFile?token=${getToken()}&` + queryString.stringify(clearEmpty(params)))
}
