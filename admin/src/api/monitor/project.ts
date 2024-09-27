import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_monitor_project = {
    Id?: number
    ProjectKey?: string
    ProjectName?: string
    ProjectType?: string
    Status?: number
    IsDelete?: number
    CreateTime?: string
    UpdateTime?: string
    DeleteTime?: string
}
// 查询
export type type_monitor_project_query = {
    ProjectKey?: string
    ProjectName?: string
    ProjectType?: string
    Status?: number
    CreateTimeStart?: string
    CreateTimeEnd?: string
    UpdateTimeStart?: string
    UpdateTimeEnd?: string
}
// 添加编辑
export type type_monitor_project_edit = {
    Id?: number
    ProjectKey?: string
    ProjectName?: string
    ProjectType?: string
    Status?: number
}

// 监控项目列表
export function monitor_project_list(params?: type_monitor_project_query) {
    return request.get<Pages<type_monitor_project>>({ url: '/monitor_project/list', params: clearEmpty(params) })
}
// 监控项目列表-所有
export function monitor_project_list_all(params?: type_monitor_project_query) {
    return request.get<type_monitor_project[]>({ url: '/monitor_project/listAll', params: clearEmpty(params) })
}

// 监控项目详情
export function monitor_project_detail(Id: number | string) {
    return request.get<type_monitor_project>({ url: '/monitor_project/detail', params: { Id } })
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
export function monitor_project_delete(Id: number | string) {
    return request.post<null>({ url: '/monitor_project/del', data: { Id } })
}
// 监控项目删除-批量
export function monitor_project_delete_batch(data: { Ids: string }) {
    return request.post<null>({ url: '/monitor_project/delBatch', data })
}

// 监控项目导入
export const monitor_project_import_file = '/monitor_project/ImportFile'

// 监控项目导出
export function monitor_project_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/monitor_project/ExportFile?token=${getToken()}&` + queryString.stringify(clearEmpty(params)))
}
