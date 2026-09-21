import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_monitor_project = {
    id: string
    project_key: string
    project_name: string
    project_type: string
    status: number
    is_delete: number
    create_time: string
    update_time: string
    delete_time: string
}
// 查询
export type type_monitor_project_query = {
    project_key?: string
    project_name?: string
    project_type?: string
    status?: number
    create_time_start?: string
    create_time_end?: string
    update_time_start?: string
    update_time_end?: string
}
// 添加编辑
export type type_monitor_project_edit = {
    id?: string
    project_key?: string
    project_name?: string
    project_type?: string
    status?: number
}

// 监控项目列表
export function monitor_project_list(params?: type_monitor_project_query) {
    return request.get<Pages<type_monitor_project>>({
        url: '/monitor_project/list',
        params: clearEmpty(params)
    })
}
// 监控项目列表-所有
export function monitor_project_list_all(params?: type_monitor_project_query) {
    return request.get<type_monitor_project[]>({
        url: '/monitor_project/list_all',
        params: clearEmpty(params)
    })
}

// 监控项目详情
export function monitor_project_detail(id: string) {
    return request.get<type_monitor_project>({ url: '/monitor_project/detail', params: { id } })
}

// 监控项目新增
export function monitor_project_add(data: type_monitor_project_edit) {
    return request.post<null>({ url: '/monitor_project/add', data })
}

// 监控项目编辑
export function monitor_project_edit(data: type_monitor_project_edit) {
    return request.post<null>({ url: '/monitor_project/edit', data })
}

// 监控项目删除
export function monitor_project_delete(id: string) {
    return request.post<null>({ url: '/monitor_project/del', data: { id } })
}
// 监控项目删除-批量
export function monitor_project_delete_batch(data: { ids: string }) {
    return request.post<null>({ url: '/monitor_project/del_batch', data })
}

// 监控项目导入
export const monitor_project_import_file = '/monitor_project/import_file'

// 监控项目导出
export function monitor_project_export_file(params: type_monitor_project_query) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/monitor_project/export_file?token=${getToken()}&` +
        queryString.stringify(clearEmpty(params)))
}
