import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 文章分类列表参数
export type type_article_cate_list = {
    name?: string
    status?: number
}

// 文章分类详情参数
export type type_article_cate_detail = {
    id: string
}

// 文章分类添加参数
export type type_article_cate_add = {
    name: string
    pid?: string
    sort?: number
    status?: number
}

// 文章分类编辑参数
export type type_article_cate_edit = {
    id: string
    name: string
    pid?: string
    sort?: number
    status?: number
}

// 文章分类删除参数
export type type_article_cate_del = {
    id: string
}

// 文章分类状态参数
export type type_article_cate_status = {
    id: string
    status: number
}

// 文章分类返回信息
export type type_article_cate_resp = {
    id: string
    pid: string
    name: string
    sort: number
    status: number
    create_time: string
    update_time: string
}

// 文章列表参数
export type type_article_list = {
    title?: string
    cateId?: string
    status?: number
    create_time_start?: string
    create_time_end?: string
}

// 文章详情参数
export type type_article_detail = {
    id: string
}

// 文章添加参数
export type type_article_add = {
    title: string
    cateId: string
    content: string
    cover?: string
    summary?: string
    author?: string
    status?: number
    sort?: number
}

// 文章编辑参数
export type type_article_edit = {
    id: string
    title: string
    cateId: string
    content: string
    cover?: string
    summary?: string
    author?: string
    status?: number
    sort?: number
}

// 文章删除参数
export type type_article_del = {
    id: string
}

// 文章状态参数
export type type_article_status = {
    id: string
    status: number
}

// 文章返回信息
export type type_article_resp = {
    id: string
    title: string
    cateId: string
    cateName: string
    content: string
    cover: string
    summary: string
    author: string
    status: number
    sort: number
    create_time: string
    update_time: string
}

// 文章分类列表
export function articleCateLists(params?: type_article_cate_list) {
    return request.get<Pages<type_article_cate_resp>>({ url: '/article/cate/list', params })
}

// 文章分类列表
export function articleCateAll(params?: type_article_cate_list) {
    return request.get<type_article_cate_resp[]>({ url: '/article/cate/all', params })
}

// 添加文章分类
export function articleCateAdd(data: type_article_cate_add) {
    return request.post({ url: '/article/cate/add', data })
}

// 编辑文章分类
export function articleCateEdit(data: type_article_cate_edit) {
    return request.post({ url: '/article/cate/edit', data })
}

// 删除文章分类
export function articleCateDelete(data: type_article_cate_del) {
    return request.post({ url: '/article/cate/del', data })
}

// 文章分类详情
export function articleCateDetail(params: type_article_cate_detail) {
    return request.get<type_article_cate_resp>({ url: '/article/cate/detail', params })
}

// 文章分类状态
export function articleCateStatus(data: type_article_cate_status) {
    return request.post({ url: '/article/cate/change', data })
}

// 文章列表
export function articleLists(params?: type_article_list) {
    return request.get<Pages<type_article_resp>>({ url: '/article/list', params })
}

// 文章列表
export function articleAll(params?: type_article_list) {
    return request.get<type_article_resp[]>({ url: '/article/all', params })
}

// 添加文章
export function articleAdd(data: type_article_add) {
    return request.post({ url: '/article/add', data })
}

// 编辑文章
export function articleEdit(data: type_article_edit) {
    return request.post({ url: '/article/edit', data })
}

// 删除文章
export function articleDelete(data: type_article_del) {
    return request.post({ url: '/article/del', data })
}

// 文章详情
export function articleDetail(params: type_article_detail) {
    return request.get<type_article_resp>({ url: '/article/detail', params })
}

// 文章状态
export function articleStatus(data: type_article_status) {
    return request.post({ url: '/article/change', data })
}
