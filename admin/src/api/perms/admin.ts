import request from '@/utils/request'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import config from '@/config'
// 管理员列表
export function adminLists(params: any) {
    return request.get({ url: '/system/admin/list', params })
}
// 管理员列表
export function adminListAll(params: any) {
    return request.get({ url: '/system/admin/listAll', params })
}
// 管理员详情
export function adminDetail(params: any) {
    return request.get({ url: '/system/admin/detail', params })
}

// 管理员添加
export function adminAdd(data: any) {
    return request.post({ url: '/system/admin/add', data })
}

// 管理员编辑
export function adminEdit(data: any) {
    return request.post({ url: '/system/admin/edit', data })
}

// 管理员删除
export function adminDelete(data: any) {
    return request.post({ url: '/system/admin/del', data })
}

// 管理员删除
export function adminStatus(data: any) {
    return request.post({ url: '/system/admin/disable', data })
}

// 部门下的管理员
export function adminListByDeptId(params: any) {
    return request.get({ url: '/system/admin/ListByDeptId', params })
}

// 导入
export const adminImportFile = '/system/admin/ImportFile'

// 导出
export function adminExportFile(params: any) {
    // return request.get({ url: '/system/admin/ExportFile', params })
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/system/admin/ExportFile?token=${getToken()}&` +
        queryString.stringify(params))
}
