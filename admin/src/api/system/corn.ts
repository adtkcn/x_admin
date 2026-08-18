import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_system_corn = {
    id: string
    task_name: string
    task_code: string
    corn_expr: string
    status: number
    created_by: string
    create_time: string
    update_time: string
}
// 查询
export type type_system_corn_query = {
    task_name?: string
    task_code?: string
    corn_expr?: string
    status?: number
    created_by?: string
    create_time_start?: string
    create_time_end?: string
    update_time_start?: string
    update_time_end?: string
    nickname?: string
}
// 添加编辑
export type type_system_corn_edit = {
    id?: string
    task_name?: string
    task_code?: string
    corn_expr?: string
    status?: number
}
export type type_task = {
    lock_ttl: number
    task_code: string
    task_desc: string
}
// 定时任务列表
export function system_corn_list(params?: type_system_corn_query) {
    return request.get<Pages<type_system_corn>>({
        url: '/system_corn/list',
        params: clearEmpty(params)
    })
}
// 定时任务列表-所有
export function system_corn_list_all(params?: type_system_corn_query) {
    return request.get<type_system_corn[]>({
        url: '/system_corn/list_all',
        params: clearEmpty(params)
    })
}

// 定时任务详情
export function system_corn_detail(id: string) {
    return request.get<type_system_corn>({ url: '/system_corn/detail', params: { id } })
}

// 定时任务新增
export function system_corn_add(data: type_system_corn_edit) {
    return request.post<null>({ url: '/system_corn/add', data })
}

// 定时任务编辑
export function system_corn_edit(data: type_system_corn_edit) {
    return request.post<null>({ url: '/system_corn/edit', data })
}

// 定时任务删除
export function system_corn_delete(id: string) {
    return request.post<null>({ url: '/system_corn/del', data: { id } })
}
// 定时任务删除-批量
export function system_corn_delete_batch(data: { ids: string }) {
    return request.post<null>({ url: '/system_corn/del_batch', data })
}

// 定时任务导入
export const system_corn_import_file = '/system_corn/import_file'

// 定时任务导出
export function system_corn_export_file(params: type_system_corn_query) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/system_corn/export_file?token=${getToken()}&` +
        queryString.stringify(clearEmpty(params)))
}

// 自定义任务列表-所有
export function system_corn_getTaskList() {
    return request.get<type_task[]>({
        url: '/system_corn/getTaskList'
    })
}
