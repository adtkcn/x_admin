import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 字典类型列表参数
export type type_setting_dict_type_list = {
    dictName?: string
    dictType?: string
    dictStatus?: number
}

// 字典类型详情参数
export type type_setting_dict_type_detail = {
    id: string
}

// 字典类型添加参数
export type type_setting_dict_type_add = {
    dictName: string
    dictType: string
    dictRemark?: string
    dictStatus: number
}

// 字典类型编辑参数
export type type_setting_dict_type_edit = {
    id: string
    dictName: string
    dictType: string
    dictRemark: string
    dictStatus: number
}

// 字典类型删除参数
export type type_setting_dict_type_del = {
    ids: string[]
}

// 字典类型返回信息
export type type_setting_dict_type_resp = {
    id: string
    dictName: string
    dictType: string
    dictRemark: string
    dictStatus: number
    createTime: string
    updateTime: string
}

// 字典数据列表参数
export type type_setting_dict_data_list = {
    dictType?: string
    name?: string
    value?: string
    status?: number
}

// 字典数据添加参数
export type type_setting_dict_data_add = {
    typeId: string
    name: string
    value: string
    color?: string
    remark?: string
    sort?: number
    status?: number
}

// 字典数据编辑参数
export type type_setting_dict_data_edit = {
    id: string
    typeId: string
    name: string
    value: string
    color?: string
    remark?: string
    sort?: number
    status?: number
}

// 字典数据删除参数
export type type_setting_dict_data_del = {
    ids: string[]
}

// 字典数据返回信息
export type type_setting_dict_data_resp = {
    id: string
    typeId: string
    name: string
    value: string
    color: string
    remark: string
    sort: number
    status: number
    createTime: string
    updateTime: string
}

// 字典类型列表
export function dictTypeLists(params?: type_setting_dict_type_list) {
    return request.get<Pages<type_setting_dict_type_resp>>({
        url: '/setting/dict/type/list',
        params
    })
}

// 字典类型列表
export function dictTypeAll() {
    return request.get<type_setting_dict_type_resp[]>({ url: '/setting/dict/type/all' })
}

// 添加字典类型
export function dictTypeAdd(data: type_setting_dict_type_add) {
    return request.post({ url: '/setting/dict/type/add', data })
}

// 编辑字典类型
export function dictTypeEdit(data: type_setting_dict_type_edit) {
    return request.post({ url: '/setting/dict/type/edit', data })
}

// 删除字典类型
export function dictTypeDelete(data: type_setting_dict_type_del) {
    return request.post({ url: '/setting/dict/type/del', data })
}

// 字典数据列表
export function dictDataAll(params: type_setting_dict_data_list) {
    return request.get<type_setting_dict_data_resp[]>({ url: '/setting/dict/data/all', params })
}

// 添加字典数据
export function dictDataAdd(data: type_setting_dict_data_add) {
    return request.post({ url: '/setting/dict/data/add', data })
}

// 编辑字典数据
export function dictDataEdit(data: type_setting_dict_data_edit) {
    return request.post({ url: '/setting/dict/data/edit', data })
}

// 删除字典数据
export function dictDataDelete(data: type_setting_dict_data_del) {
    return request.post({ url: '/setting/dict/data/del', data })
}
