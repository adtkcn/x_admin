import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'

// 相册分类返回信息
export type type_album_cate = {
    id?: string
    pid?: string
    name?: string
    isDelete?: number
    createTime?: string
    updateTime?: string
    deleteTime?: string
}

// 相册分类查询参数
export type type_album_cate_query = {
    pid?: string
    name?: string
    createTimeStart?: string
    createTimeEnd?: string
    updateTimeStart?: string
    updateTimeEnd?: string
}

// 相册分类添加编辑参数
export type type_album_cate_edit = {
    id?: string
    pid?: string
    name?: string
}

// 相册分类列表
export function album_cate_list(params?: type_album_cate_query) {
    return request.get<Pages<type_album_cate>>({ url: '/album_cate/list', params })
}

// 相册分类列表-所有
export function album_cate_list_all(params?: type_album_cate_query) {
    return request.get<Pages<type_album_cate>>({ url: '/album_cate/list_all', params })
}

// 相册分类详情
export function album_cate_detail(id: number | string) {
    return request.get<type_album_cate>({ url: '/album_cate/detail', params: { id } })
}

// 相册分类新增
export function album_cate_add(data: type_album_cate_edit) {
    return request.post<null>({ url: '/album_cate/add', data })
}

// 相册分类编辑
export function album_cate_edit(data: type_album_cate_edit) {
    return request.post<null>({ url: '/album_cate/edit', data })
}

// 相册分类删除
export function album_cate_delete(id: number | string) {
    return request.post<null>({ url: '/album_cate/del', data: { id } })
}

// 相册分类导入
export const album_cate_import_file = '/album_cate/import_file'

// 相册分类导出
export function album_cate_export_file(params: type_album_cate_query) {
    return (window.location.href =
        `${config.baseUrl}${config.urlPrefix}/album_cate/export_file?token=${getToken()}&` +
        queryString.stringify(params))
}
