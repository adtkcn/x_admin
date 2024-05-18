import request from '@/utils/request'

// 文章分类列表
export function articleCateLists(params?: any) {
    return request.get({ url: '/article/cate/list', params })
}
// 文章分类列表
export function articleCateAll(params?: any) {
    return request.get({ url: '/article/cate/all', params })
}

// 添加文章分类
export function articleCateAdd(data: any) {
    return request.post({ url: '/article/cate/add', data })
}

// 编辑文章分类
export function articleCateEdit(data: any) {
    return request.post({ url: '/article/cate/edit', data })
}

// 删除文章分类
export function articleCateDelete(data: any) {
    return request.post({ url: '/article/cate/del', data })
}

// 文章分类详情
export function articleCateDetail(params: any) {
    return request.get({ url: '/article/cate/detail', params })
}

// 文章分类状态
export function articleCateStatus(data: any) {
    return request.post({ url: '/article/cate/change', data })
}

// 文章列表
export function articleLists(params?: any) {
    return request.get({ url: '/article/list', params })
}
// 文章列表
export function articleAll(params?: any) {
    return request.get({ url: '/article/all', params })
}

// 添加文章
export function articleAdd(data: any) {
    return request.post({ url: '/article/add', data })
}

// 编辑文章
export function articleEdit(data: any) {
    return request.post({ url: '/article/edit', data })
}

// 删除文章
export function articleDelete(data: any) {
    return request.post({ url: '/article/del', data })
}

// 文章详情
export function articleDetail(params: any) {
    return request.get({ url: '/article/detail', params })
}

// 文章分类状态
export function articleStatus(data: any) {
    return request.post({ url: '/article/change', data })
}
