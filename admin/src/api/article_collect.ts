import request from '@/utils/request'

// 文章收藏列表
export function article_collect_lists(params?: Record<string, any>) {
    return request.get({ url: '/article_collect/list', params })
}

// 文章收藏详情
export function article_collect_detail(params: Record<string, any>) {
    return request.get({ url: '/article_collect/detail', params })
}

// 文章收藏新增
export function article_collect_add(params: Record<string, any>) {
    return request.post({ url: '/article_collect/add', params })
}

// 文章收藏编辑
export function article_collect_edit(params: Record<string, any>) {
    return request.post({ url: '/article_collect/edit', params })
}

// 文章收藏删除
export function article_collect_delete(params: Record<string, any>) {
    return request.post({ url: '/article_collect/del', params })
}
