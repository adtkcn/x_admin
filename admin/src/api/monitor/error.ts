import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_monitor_error = {
    Id?: number
    ProjectKey?: string
    EventType?: string
    Path?: string
    Message?: string
    Stack?: string
    Md5?: string
    CreateTime?: string
}
// 查询
export type type_monitor_error_query = {
    ProjectKey?: string
    EventType?: string
    Path?: string
    Message?: string
    Stack?: string
    Md5?: string
    CreateTimeStart?: string
    CreateTimeEnd?: string
}
// 添加编辑
export type type_monitor_error_edit = {
    Id?: number
    ProjectKey?: string
    EventType?: string
    Path?: string
    Message?: string
    Stack?: string
    Md5?: string
}

// 监控-错误列列表
export function monitor_error_list(params?: type_monitor_error_query) {
    return request.get<Pages<type_monitor_error>>({
        url: '/monitor_error/list',
        params: clearEmpty(params)
    })
}
// 监控-错误列列表-所有
export function monitor_error_list_all(params?: type_monitor_error_query) {
    return request.get<type_monitor_error[]>({
        url: '/monitor_error/listAll',
        params: clearEmpty(params)
    })
}

// 监控-错误列详情
export function monitor_error_detail(Id: number | string) {
    return request.get<type_monitor_error>({ url: '/monitor_error/detail', params: { Id } })
}

// 监控-错误列新增
export function monitor_error_add(data: type_monitor_error_edit) {
    return request.post<null>({ url: '/monitor_error/add', data })
}

// 监控-错误列删除
export function monitor_error_delete(Id: number | string) {
    return request.post<null>({ url: '/monitor_error/del', data: { Id } })
}
// 监控-错误列删除-批量
export function monitor_error_delete_batch(data: { Ids: string }) {
    return request.post<null>({ url: '/monitor_error/delBatch', data })
}

// 监控-错误列导入
export const monitor_error_import_file = '/monitor_error/ImportFile'

// 监控-错误列导出
export function monitor_error_export_file(params: any) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/monitor_error/ExportFile?token=${getToken()}&` +
        queryString.stringify(clearEmpty(params)))
}
