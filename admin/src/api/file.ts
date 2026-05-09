import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 文件分类添加参数
export type type_file_cate_add = {
    name: string
    pid?: string
}

// 文件分类编辑参数
export type type_file_cate_edit = {
    id: string
    name: string
}

// 文件分类删除参数
export type type_file_cate_del = {
    id: string
}

// 文件分类列表参数
export type type_file_cate_list = {
    pid?: string
}

// 文件分类返回信息
export type type_file_cate_resp = {
    id: string
    pid: string
    name: string
    createTime: string
}

// 文件列表参数
export type type_file_list = {
    cateId?: string
    name?: string
    page?: number
    limit?: number
}

// 文件删除参数
export type type_file_del = {
    ids: string[]
}

// 文件移动参数
export type type_file_move = {
    ids: string[]
    cid: string
}

// 文件重命名参数
export type type_file_rename = {
    id: string
    name: string
}

// 文件返回信息
export type type_file_resp = {
    id: string
    cateId: string
    name: string
    path: string
    url: string
    size: number
    type: string
    createTime: string
}

// 添加文件分类
export function fileCateAdd(data: type_file_cate_add) {
    return request.post({ url: '/common/album/cateAdd', data })
}

// 编辑文件分类
export function fileCateEdit(data: type_file_cate_edit) {
    return request.post({ url: '/common/album/cateRename', data })
}

// 文件分类删除
export function fileCateDelete(data: type_file_cate_del) {
    return request.post({ url: '/common/album/cateDel', data })
}

// 文件分类列表
export function fileCateLists(params: type_file_cate_list) {
    return request.get<type_file_cate_resp[]>({ url: '/common/album/cateList', params })
}

// 文件列表
export function fileList(params: type_file_list) {
    return request.get<Pages<type_file_resp>>({ url: '/common/album/albumList', params })
}

// 文件删除
export function fileDelete(data: type_file_del) {
    return request.post({ url: '/common/album/albumDel', data })
}

// 文件移动
export function fileMove(data: type_file_move) {
    return request.post({ url: '/common/album/albumMove', data })
}

// 文件重命名
export function fileRename(data: type_file_rename) {
    return request.post({ url: '/common/album/albumRename', data })
}
