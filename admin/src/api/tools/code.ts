import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 代码生成表列表参数
export type type_gen_table_list = {
    tableName?: string
    tableComment?: string
}

// 数据表列表参数
export type type_gen_db_list = {
    tableName?: string
}

// 选择表参数
export type type_gen_import_table = {
    tables: string[]
}

// 表详情参数
export type type_gen_table_detail = {
    id: string
}

// 同步表参数
export type type_gen_sync_table = {
    id: string
}

// 删除表参数
export type type_gen_del_table = {
    ids: string[]
}

// 编辑表字段参数
export type type_gen_edit_column = {
    id: string
    tableId: string
    columnName: string
    columnLength: number
    columnType: string
    goField: string
    goType: string
    columnComment: string
    isPk: number
    isIncrement: number
    isRequired: number
    isInsert: number
    isEdit: number
    isList: number
    isQuery: number
    queryType: string
    htmlType: string
    dictType: string
    listAllApi?: string
    createTime?: string
    updateTime?: string
}

// 编辑表参数
export type type_gen_edit_table = {
    id: string
    tableName: string
    tableComment: string
    packageName?: string
    moduleName?: string
    businessName?: string
    functionName?: string
    author?: string
    remark?: string
    columns: type_gen_edit_column[]
}

// 预览代码参数
export type type_gen_preview_code = {
    id: string
}

// 下载代码参数
export type type_gen_download_code = {
    tables: string
}

// 表信息返回
export type type_gen_table_resp = {
    id: string
    tableName: string
    tableComment: string
    packageName: string
    moduleName: string
    businessName: string
    functionName: string
    author: string
    remark: string
    createTime: string
    updateTime: string
}

// 代码生成已选数据表列表接口
export function generateTable(params: type_gen_table_list) {
    return request.get<Pages<type_gen_table_resp>>({ url: '/gen/list', params })
}

// 数据表列表接口
export function dataTable(params: type_gen_db_list) {
    return request.get({ url: '/gen/db', params })
}

//选择要生成代码的数据表
export function selectTable(params: type_gen_import_table) {
    return request.post({ url: '/gen/importTable', params })
}

// 已选择的数据表详情
export function tableDetail(params: type_gen_table_detail) {
    return request.get({ url: '/gen/detail', params })
}

//同步字段
export function syncColumn(params: type_gen_sync_table) {
    return request.post({ url: '/gen/syncTable', params })
}

//删除已选择的数据表
export function generateDelete(data: type_gen_del_table) {
    return request.post({ url: '/gen/delTable', data })
}

//编辑已选表字段
export function generateEdit(data: type_gen_edit_table) {
    return request.post({ url: '/gen/editTable', data })
}

//预览代码
export function generatePreview(params: type_gen_preview_code) {
    return request.get({ url: '/gen/previewCode', params })
}

//下载代码
export function downloadCode(params: type_gen_download_code) {
    return request.get(
        { responseType: 'blob', url: '/gen/downloadCode', params },
        {
            isTransformResponse: false
        }
    )
}
