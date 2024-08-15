import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_system_log_sms = {
    Id?: number
    Scene?: number
    Mobile?: string
    Content?: string
    Status?: number
    Results?: string
    SendTime?: string
    CreateTime?: string
    UpdateTime?: string
}
// 查询
export type type_system_log_sms_query = {
    Scene?: number
    Mobile?: string
    Content?: string
    Status?: number
    Results?: string
    SendTime_start?: string
    SendTime_end?: string
    CreateTime_start?: string
    CreateTime_end?: string
    UpdateTime_start?: string
    UpdateTime_end?: string
}
// 添加编辑
export type type_system_log_sms_edit = {
    Id?: number
    Scene?: number
    Mobile?: string
    Content?: string
    Status?: number
    Results?: string
    SendTime?: string
}

// 系统短信日志列表
export function system_log_sms_list(params?: type_system_log_sms_query) {
    return request.get<Pages<type_system_log_sms>>({
        url: '/system_log_sms/list',
        params: clearEmpty(params)
    })
}
// 系统短信日志列表-所有
export function system_log_sms_list_all(params?: type_system_log_sms_query) {
    return request.get<type_system_log_sms[]>({
        url: '/system_log_sms/listAll',
        params: clearEmpty(params)
    })
}

// 系统短信日志详情
export function system_log_sms_detail(id: number | string) {
    return request.get<type_system_log_sms>({ url: '/system_log_sms/detail', params: { id } })
}

// 系统短信日志新增
export function system_log_sms_add(data: type_system_log_sms_edit) {
    return request.post<null>({ url: '/system_log_sms/add', data })
}

// 系统短信日志编辑
export function system_log_sms_edit(data: type_system_log_sms_edit) {
    return request.post<null>({
        url: '/system_log_sms/edit',
        data: data
        // data: clearEmpty(data)
    })
}

// 系统短信日志删除
export function system_log_sms_delete(id: number | string) {
    return request.post<null>({ url: '/system_log_sms/del', data: { id } })
}

// 系统短信日志导入
export const system_log_sms_import_file = '/system_log_sms/ImportFile'

// 系统短信日志导出
export function system_log_sms_export_file(params: any) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/system_log_sms/ExportFile?token=${getToken()}&` +
        queryString.stringify(clearEmpty(params)))
}
