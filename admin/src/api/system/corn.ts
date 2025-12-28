import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_system_corn = {
    Id?: string
    TaskName?: string
    TaskCode?: string
    CornExpr?: string
    Status?: number
    CreatedBy?: string
    IsDelete?: number
    CreateTime?: string
    UpdateTime?: string
    DeleteTime?: string
}
// 查询
export type type_system_corn_query = {
    TaskName?: string
    TaskCode?: string
    CornExpr?: string
    Status?: any
    CreatedBy?: string
    CreateTimeStart?: string
    CreateTimeEnd?: string
    UpdateTimeStart?: string
    UpdateTimeEnd?: string
    CreatedByNickname?: string
    CreatedByUsername?: string
}
// 添加编辑
export type type_system_corn_edit = {
    Id?: string
    TaskName?: string
    TaskCode?: string
    CornExpr?: string
    Status?: number
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
        url: '/system_corn/listAll',
        params: clearEmpty(params)
    })
}

// 定时任务详情
export function system_corn_detail(Id: number | string) {
    return request.get<type_system_corn>({ url: '/system_corn/detail', params: { Id } })
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
export function system_corn_delete(Id: number | string) {
    return request.post<null>({ url: '/system_corn/del', data: { Id } })
}
// 定时任务删除-批量
export function system_corn_delete_batch(data: { Ids: string }) {
    return request.post<null>({ url: '/system_corn/delBatch', data })
}

// 定时任务导入
export const system_corn_import_file = '/system_corn/ImportFile'

// 定时任务导出
export function system_corn_export_file(params: any) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/system_corn/ExportFile?token=${getToken()}&` +
        queryString.stringify(clearEmpty(params)))
}
