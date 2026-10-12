import request from '@/utils/request'

// ===================== 产品分类（树） =====================
export type type_iot_product_category = {
    id?: string
    pid?: string
    name?: string
    sort?: number
    remark?: string
    create_time?: string
    update_time?: string
    children?: type_iot_product_category[]
}
export function productCategoryAll() {
    return request.get<type_iot_product_category[]>({ url: '/iot/productCategory/all' })
}
export function productCategoryDetail(params: { id: string }) {
    return request.get<type_iot_product_category>({ url: '/iot/productCategory/detail', params })
}
export function productCategoryAdd(data: type_iot_product_category) {
    return request.post({ url: '/iot/productCategory/add', data })
}
export function productCategoryEdit(data: type_iot_product_category) {
    return request.post({ url: '/iot/productCategory/edit', data })
}
export function productCategoryDel(data: { id: string }) {
    return request.post({ url: '/iot/productCategory/del', data })
}
export function productCategorySort(data: { ids: string[] }) {
    return request.post({ url: '/iot/productCategory/sort', data })
}
