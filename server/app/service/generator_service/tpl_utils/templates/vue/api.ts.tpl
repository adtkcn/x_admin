import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'
import { clearEmpty } from '@/utils/util'

export type type_{{{.ModuleName}}} = {
{{{- range .Columns }}}
{{{- if or .IsList .IsPk }}}
    {{{.TsField }}}: {{{.TsType}}}{{{ if eq .IsRequired 0 }}}|null{{{ end }}}
{{{- end }}}
{{{- end }}}
}
// 查询
export type type_{{{.ModuleName}}}_query = {
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
    {{{.TsField }}}_start?: string
    {{{.TsField }}}_end?: string
{{{- else }}}
    {{{.TsField }}}?: {{{.TsType}}}
{{{- end }}}
{{{- end }}}
{{{- end }}}
}
// 添加编辑
export type type_{{{.ModuleName}}}_edit = {
{{{- range .Columns }}}
{{{- if or .IsEdit .IsInsert }}}
    {{{.TsField }}}?: {{{.TsType}}}{{{ if eq .IsRequired 0 }}}|null{{{ end }}}
{{{- end }}}
{{{- end }}}
}

// {{{.FunctionName}}}列表
export function {{{.ModuleName}}}_list(params?: type_{{{.ModuleName}}}_query) {
    return request.get<Pages<type_{{{.ModuleName}}}>>({ url: '/{{{.ModuleName}}}/list', params: clearEmpty(params) })
}
// {{{.FunctionName}}}列表-所有
export function {{{.ModuleName}}}_list_all(params?: type_{{{.ModuleName}}}_query) {
    return request.get<type_{{{.ModuleName}}}[]>({ url: '/{{{.ModuleName}}}/list_all', params: clearEmpty(params) })
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}_detail({{{ .PrimaryTsField }}}: {{{.PrimaryTsType}}}) {
    return request.get<type_{{{.ModuleName}}}>({ url: '/{{{.ModuleName}}}/detail', params: { {{{ .PrimaryTsField }}} } })
}

// {{{.FunctionName}}}新增
export function {{{.ModuleName}}}_add(data: type_{{{.ModuleName}}}_edit) {
    return request.post<string>({ url: '/{{{.ModuleName}}}/add', data })
}

// {{{.FunctionName}}}编辑
export function {{{.ModuleName}}}_edit(data: type_{{{.ModuleName}}}_edit) {
    return request.post<string>({ url: '/{{{.ModuleName}}}/edit', data })
}

// {{{.FunctionName}}}删除
export function {{{.ModuleName}}}_delete({{{ .PrimaryTsField }}}: {{{.PrimaryTsType}}}) {
    return request.post<null>({ url: '/{{{.ModuleName}}}/del', data: { {{{ .PrimaryTsField }}} } })
}
// {{{.FunctionName}}}删除-批量
export function {{{.ModuleName}}}_delete_batch(data: { ids: string }) {
    return request.post<null>({ url: '/{{{.ModuleName}}}/del_batch', data })
}

// {{{.FunctionName}}}导入
export const {{{.ModuleName}}}_import_file = '/{{{.ModuleName}}}/import_file'

// {{{.FunctionName}}}导出
export function {{{.ModuleName}}}_export_file(params: type_{{{.ModuleName}}}_query) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/{{{.ModuleName}}}/export_file?token=${getToken()}&` + queryString.stringify(clearEmpty(params)))
}
