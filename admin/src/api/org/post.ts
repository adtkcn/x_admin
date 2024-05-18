import request from '@/utils/request'
// 岗位详情
export function postDetail(params: any) {
    return request.get({ url: '/system/post/detail', params })
}

// 岗位列表
export function postLists(params?: any) {
    return request.get({ url: '/system/post/list', params })
}
// 岗位全部列表
export function postAll(params?: any) {
    return request.get({ url: '/system/post/all', params })
}

// 添加岗位
export function postAdd(data: any) {
    return request.post({ url: '/system/post/add', data })
}

// 编辑岗位
export function postEdit(data: any) {
    return request.post({ url: '/system/post/edit', data })
}

// 删除岗位
export function postDelete(data: any) {
    return request.post({ url: '/system/post/del', data })
}
