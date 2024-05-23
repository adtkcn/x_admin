import request from '@/utils/request'
import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'


// {{{.FunctionName}}}列表
export function {{{.ModuleName}}}_list(params?: Record<string, any>) {
    return request.get({ url: '/{{{.ModuleName}}}/list', params })
}
// {{{.FunctionName}}}列表-所有
export function {{{.ModuleName}}}_list_all(params?: Record<string, any>) {
    return request.get({ url: '/{{{.ModuleName}}}/listAll', params })
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}_detail({{{ .PrimaryKey }}}: number | string) {
    return request.get({ url: '/{{{.ModuleName}}}/detail', { {{{ .PrimaryKey }}} } })
}

// {{{.FunctionName}}}新增
export function {{{.ModuleName}}}_add(data: Record<string, any>) {
    return request.post({ url: '/{{{.ModuleName}}}/add', data })
}

// {{{.FunctionName}}}编辑
export function {{{.ModuleName}}}_edit(data: Record<string, any>) {
    return request.post({ url: '/{{{.ModuleName}}}/edit', data })
}

// {{{.FunctionName}}}删除
export function {{{.ModuleName}}}_delete({{{ .PrimaryKey }}}: number | string) {
    return request.post({ url: '/{{{.ModuleName}}}/del', { {{{ .PrimaryKey }}} } })
}

// {{{.FunctionName}}}导入
export const {{{.ModuleName}}}_import_file = '/{{{.ModuleName}}}/ImportFile'

// {{{.FunctionName}}}导出
export function {{{.ModuleName}}}_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/{{{.ModuleName}}}/ExportFile?token=${getToken()}&` + queryString.stringify(params))
}
