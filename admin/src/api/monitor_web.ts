import request from '@/utils/request'
import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'

// 错误收集error列表
export function monitor_web_list(params?: Record<string, any>) {
    return request.get({ url: '/monitor_web/list', params })
}
// 错误收集error列表-所有
export function monitor_web_list_all(params?: Record<string, any>) {
    return request.get({ url: '/monitor_web/listAll', params })
}

// 错误收集error详情
export function monitor_web_detail(params: Record<string, any>) {
    return request.get({ url: '/monitor_web/detail', params })
}

// 错误收集error新增
export function monitor_web_add(params: Record<string, any>) {
    return request.post({ url: '/monitor_web/add', params })
}

// 错误收集error编辑
export function monitor_web_edit(params: Record<string, any>) {
    return request.post({ url: '/monitor_web/edit', params })
}

// 错误收集error删除
export function monitor_web_delete(params: Record<string, any>) {
    return request.post({ url: '/monitor_web/del', params })
}

// 错误收集error导入
export const monitor_web_import_file = '/monitor_web/ImportFile'

// 错误收集error导出
export function monitor_web_export_file(params: any) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/monitor_web/ExportFile?token=${getToken()}&` +
        queryString.stringify(params))
}
