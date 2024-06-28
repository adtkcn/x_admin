import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'

export type type_monitor_client = {
    id?: number;
    projectKey?: string;
    clientId?: string;
    userId?: string;
    os?: string;
    browser?: string;
    city?: string;
    width?: number;
    height?: number;
    ua?: string;
    createTime?: string;
    clientTime?: string;
}
// 查询
export type type_monitor_client_query = {
    projectKey?: string;
    clientId?: string;
    userId?: string;
    os?: string;
    browser?: string;
    city?: string;
    width?: number;
    height?: number;
    ua?: string;
    createTimeStart?: string;
    createTimeEnd?: string;
    clientTimeStart?: string;
    clientTimeEnd?: string;
}
// 添加编辑
export type type_monitor_client_edit = {
    clientId?: string;
    userId?: string;
    os?: string;
    browser?: string;
    city?: string;
    width?: number;
    height?: number;
    ua?: string;
    clientTime?: string;
}

// 监控-客户端信息列表
export function monitor_client_list(params?: type_monitor_client_query) {
    return request.get<Pages<type_monitor_client>>({ url: '/monitor_client/list', params })
}
// 监控-客户端信息列表-所有
export function monitor_client_list_all(params?: type_monitor_client_query) {
    return request.get<Pages<type_monitor_client>>({ url: '/monitor_client/listAll', params })
}

// 监控-客户端信息详情
export function monitor_client_detail(id: number | string) {
    return request.get<type_monitor_client>({ url: '/monitor_client/detail', params: { id } })
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
export function monitor_client_delete(id: number | string) {
    return request.post<null>({ url: '/monitor_client/del', data: { id } })
}

// 监控-客户端信息导入
export const monitor_client_import_file = '/monitor_client/ImportFile'

// 监控-客户端信息导出
export function monitor_client_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/monitor_client/ExportFile?token=${getToken()}&` + queryString.stringify(params))
}
