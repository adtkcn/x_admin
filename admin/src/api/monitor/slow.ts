import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_monitor_slow = {
    Id?: number
    ProjectKey?: string
    ClientId?: string
    UserId?: string
    Path?: string
    Time?: number
    CreateTime?: string
}
// 查询
export type type_monitor_slow_query = {
    ProjectKey?: string
    ClientId?: string
    UserId?: string
    Path?: string
    Time?: number
    CreateTimeStart?: string
    CreateTimeEnd?: string
}
// 添加编辑
export type type_monitor_slow_edit = {
    Id?: number
    ProjectKey?: string
    ClientId?: string
    UserId?: string
    Path?: string
    Time?: number
}

// 监控-错误列列表
export function monitor_slow_list(params?: type_monitor_slow_query) {
    return request.get<Pages<type_monitor_slow>>({ url: '/monitor_slow/list', params: clearEmpty(params) })
}
// 监控-错误列列表-所有
export function monitor_slow_list_all(params?: type_monitor_slow_query) {
    return request.get<type_monitor_slow[]>({ url: '/monitor_slow/listAll', params: clearEmpty(params) })
}

// 监控-错误列详情
export function monitor_slow_detail(Id: number | string) {
    return request.get<type_monitor_slow>({ url: '/monitor_slow/detail', params: { Id } })
}

// 监控-错误列新增
export function monitor_slow_add(data: type_monitor_slow_edit) {
    return request.post<null>({ url: '/monitor_slow/add', data })
}

// 监控-错误列编辑
export function monitor_slow_edit(data: type_monitor_slow_edit) {
    return request.post<null>({ url: '/monitor_slow/edit', data })
}

// 监控-错误列删除
export function monitor_slow_delete(Id: number | string) {
    return request.post<null>({ url: '/monitor_slow/del', data: { Id } })
}
// 监控-错误列删除-批量
export function monitor_slow_delete_batch(data: { Ids: string }) {
    return request.post<null>({ url: '/monitor_slow/delBatch', data })
}

// 监控-错误列导入
export const monitor_slow_import_file = '/monitor_slow/ImportFile'

// 监控-错误列导出
export function monitor_slow_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/monitor_slow/ExportFile?token=${getToken()}&` + queryString.stringify(clearEmpty(params)))
}
