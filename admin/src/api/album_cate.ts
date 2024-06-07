import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'

type album_cate = {
    id?: number
    pid?: number
    type?: number
    name?: string
    isDelete?: number
    createTime?: string
    updateTime?: string
    deleteTime?: string
}
// 查询
type album_cate_query = {
    pid?: number
    type?: number
    name?: string
    createTime?: string
    updateTime?: string
}
// 添加编辑
type album_cate_edit = {
    id?: number
    pid?: number
    type?: number
    name?: string
}

// 相册分类列表
export function album_cate_list(params?: album_cate_query) {
    return request.get<Pages<album_cate>>({ url: '/album_cate/list', params })
}
// 相册分类列表-所有
export function album_cate_list_all(params?: album_cate_query) {
    return request.get<Pages<album_cate>>({ url: '/album_cate/listAll', params })
}

// 相册分类详情
export function album_cate_detail(id: number | string) {
    return request.get<album_cate>({ url: '/album_cate/detail', params: { id } })
}

// 相册分类新增
export function album_cate_add(data: album_cate_edit) {
    return request.post<null>({ url: '/album_cate/add', data })
}

// 相册分类编辑
export function album_cate_edit(data: album_cate_edit) {
    return request.post<null>({ url: '/album_cate/edit', data })
}

// 相册分类删除
export function album_cate_delete(id: number | string) {
    return request.post<null>({ url: '/album_cate/del', data: { id } })
}

// 相册分类导入
export const album_cate_import_file = '/album_cate/ImportFile'

// 相册分类导出
export function album_cate_export_file(params: any) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/album_cate/ExportFile?token=${getToken()}&` +
        queryString.stringify(params))
}
