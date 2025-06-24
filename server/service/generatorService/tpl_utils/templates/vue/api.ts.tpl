import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_{{{.ModuleName}}} = {
{{{- range .Columns }}}
    {{{toUpperCamelCase .GoField }}}?: {{{goToTsType .GoType}}}
{{{- end }}}
}
// 查询
export type type_{{{.ModuleName}}}_query = {
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
    {{{toUpperCamelCase .GoField }}}Start?: string
    {{{toUpperCamelCase .GoField }}}End?: string
{{{- else }}}
    {{{toUpperCamelCase .GoField }}}?: {{{goToTsType .GoType}}}
{{{- end }}}
{{{- end }}}
{{{- end }}}
}
// 添加编辑
export type type_{{{.ModuleName}}}_edit = {
{{{- range .Columns }}}
{{{- if or .IsEdit .IsInsert }}}
    {{{toUpperCamelCase .GoField }}}?: {{{goToTsType .GoType}}}
{{{- end }}}
{{{- end }}}
}

// {{{.FunctionName}}}列表
export function {{{.ModuleName}}}_list(params?: type_{{{.ModuleName}}}_query) {
    return request.get<Pages<type_{{{.ModuleName}}}>>({ url: '/{{{.ModuleName}}}/list', params: clearEmpty(params) })
}
// {{{.FunctionName}}}列表-所有
export function {{{.ModuleName}}}_list_all(params?: type_{{{.ModuleName}}}_query) {
    return request.get<type_{{{.ModuleName}}}[]>({ url: '/{{{.ModuleName}}}/listAll', params: clearEmpty(params) })
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}_detail({{{toUpperCamelCase .PrimaryKey }}}: number | string) {
    return request.get<type_{{{.ModuleName}}}>({ url: '/{{{.ModuleName}}}/detail', params: { {{{toUpperCamelCase .PrimaryKey }}} } })
}

// {{{.FunctionName}}}新增
export function {{{.ModuleName}}}_add(data: type_{{{.ModuleName}}}_edit) {
    return request.post<null>({ url: '/{{{.ModuleName}}}/add', data })
}

// {{{.FunctionName}}}编辑
export function {{{.ModuleName}}}_edit(data: type_{{{.ModuleName}}}_edit) {
    return request.post<null>({ url: '/{{{.ModuleName}}}/edit', data })
}

// {{{.FunctionName}}}删除
export function {{{.ModuleName}}}_delete({{{toUpperCamelCase .PrimaryKey }}}: number | string) {
    return request.post<null>({ url: '/{{{.ModuleName}}}/del', data: { {{{toUpperCamelCase .PrimaryKey }}} } })
}
// {{{.FunctionName}}}删除-批量
export function {{{.ModuleName}}}_delete_batch(data: { Ids: string }) {
    return request.post<null>({ url: '/{{{.ModuleName}}}/delBatch', data })
}

// {{{.FunctionName}}}导入
export const {{{.ModuleName}}}_import_file = '/{{{.ModuleName}}}/ImportFile'

// {{{.FunctionName}}}导出
export function {{{.ModuleName}}}_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/{{{.ModuleName}}}/ExportFile?token=${getToken()}&` + queryString.stringify(clearEmpty(params)))
}
