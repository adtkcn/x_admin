import request from '@/utils/request'

export function fileCateAdd(data: Record<string, any>) {
    return request.post({ url: '/common/album/cateAdd', data })
}

export function fileCateEdit(data: Record<string, any>) {
    return request.post({ url: '/common/album/cateRename', data })
}

// 文件分类删除
export function fileCateDelete(data: Record<string, any>) {
    return request.post({ url: '/common/album/cateDel', data })
}

// 文件分类列表
export function fileCateLists(params: Record<string, any>) {
    return request.get({ url: '/common/album/cateList', params })
}

// 文件列表
export function fileList(params: Record<string, any>) {
    return request.get({ url: '/common/album/albumList', params })
}

// 文件删除
export function fileDelete(data: Record<string, any>) {
    return request.post({ url: '/common/album/albumDel', data })
}

// 文件移动
export function fileMove(data: Record<string, any>) {
    return request.post({ url: '/common/album/albumMove', data })
}

// 文件重命名
export function fileRename(data: { id: number; name: string }) {
    return request.post({ url: '/common/album/albumRename', data })
}
