import request from '@/utils/request'

// 部门列表参数
export type type_system_dept_list = {
    name?: string
    is_stop?: number
}

// 部门详情参数
export type type_system_dept_detail = {
    id: string
}

// 部门添加参数
export type type_system_dept_add = {
    pid?: string
    name?: string
    duty_id?: string
    duty?: string
    mobile?: string
    is_stop?: number
    sort?: number
}

// 部门编辑参数
export type type_system_dept_edit = {
    id?: string
    pid?: string
    name?: string
    duty_id?: string
    duty?: string
    mobile?: string
    is_stop?: number
    sort?: number
}

// 部门删除参数
export type type_system_dept_del = {
    id: string
}

// 部门返回信息
export type type_system_dept_resp = {
    id: string
    pid: string
    name: string
    duty_id: string
    duty: string
    mobile: string
    sort: number
    is_stop: number
    create_time: string
    update_time: string
    children?: type_system_dept_resp[]
}

// 部门列表
export function deptLists(params?: type_system_dept_list) {
    return request.get<type_system_dept_resp[]>({ url: '/system/dept/list', params })
}

// 部门列表-全部
export function deptAll() {
    return request.get<type_system_dept_resp[]>({ url: '/system/dept/all' })
}

// 部门详情
export function deptDetail(params: type_system_dept_detail) {
    return request.get<type_system_dept_resp>({ url: '/system/dept/detail', params })
}

// 添加部门
export function deptAdd(data: type_system_dept_add) {
    return request.post({ url: '/system/dept/add', data })
}

// 编辑部门
export function deptEdit(data: type_system_dept_edit) {
    return request.post({ url: '/system/dept/edit', data })
}

// 删除部门
export function deptDelete(data: type_system_dept_del) {
    return request.post({ url: '/system/dept/del', data })
}

// 部门拖拽排序参数
export type type_system_dept_sort = {
    ids: string[]
}

// 拖拽排序
export function deptSort(data: type_system_dept_sort) {
    return request.post({ url: '/system/dept/sort', data })
}
