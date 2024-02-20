import request from '@/utils/request'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
// 管理员列表
export function adminLists(params: any) {
    return request.get({ url: '/system/admin/list', params })
}

// 管理员添加
export function adminAdd(params: any) {
    return request.post({ url: '/system/admin/add', params })
}

// 管理员编辑
export function adminDetail(params: any) {
    return request.get({ url: '/system/admin/detail', params })
}

// 管理员编辑
export function adminEdit(params: any) {
    return request.post({ url: '/system/admin/edit', params })
}

// 管理员删除
export function adminDelete(params: any) {
    return request.post({ url: '/system/admin/del', params })
}

// 管理员删除
export function adminStatus(params: any) {
    return request.post({ url: '/system/admin/disable', params })
}

// 部门下的管理员
export function adminListByDeptId(params: any) {
    return request.get({ url: '/system/admin/ListByDeptId', params })
}

// 导出
export function adminExportFile(params: any) {
    // return request.get({ url: '/system/admin/ExportFile', params })
    return (window.location.href =
        `/api/admin/system/admin/ExportFile?token=${getToken()}&` + queryString.stringify(params))
}
