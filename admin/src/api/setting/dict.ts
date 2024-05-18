import request from '@/utils/request'

// 字典类型列表
export function dictTypeLists(params?: any) {
    return request.get({ url: '/setting/dict/type/list', params })
}

// 字典类型列表
export function dictTypeAll(params?: any) {
    return request.get({ url: '/setting/dict/type/all', params })
}

// 添加字典类型
export function dictTypeAdd(data: any) {
    return request.post({ url: '/setting/dict/type/add', data })
}

// 编辑字典类型
export function dictTypeEdit(data: any) {
    return request.post({ url: '/setting/dict/type/edit', data })
}

// 删除字典类型
export function dictTypeDelete(data: any) {
    return request.post({ url: '/setting/dict/type/del', data })
}

// 字典数据列表
// export function dictDataLists(params: any) {
//     return request.get(
//         { url: '/setting/dict/data/list', params },
//     )
// }

// 字典数据列表
export function dictDataAll(params: any) {
    return request.get({ url: '/setting/dict/data/all', params })
}

// 添加字典数据
export function dictDataAdd(data: any) {
    return request.post({ url: '/setting/dict/data/add', data })
}

// 编辑字典数据
export function dictDataEdit(data: any) {
    return request.post({ url: '/setting/dict/data/edit', data })
}

// 删除字典数据
export function dictDataDelete(data: any) {
    return request.post({ url: '/setting/dict/data/del', data })
}
