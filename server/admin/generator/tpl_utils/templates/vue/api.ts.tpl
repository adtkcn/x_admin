import request from '@/utils/request'
import type { Pages } from '@/utils/request'

import config from '@/config'
import queryString from 'query-string'
import { getToken } from '@/utils/auth'

export type type_{{{.ModuleName}}} = {
{{{- range .Columns }}}
    {{{ toCamelCase .GoField }}}?: {{{goToTsType .GoType}}};
{{{- end }}}
}
// 查询
export type type_{{{.ModuleName}}}_query = {
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
    {{{ toCamelCase .GoField }}}Start?: string;
    {{{ toCamelCase .GoField }}}End?: string;
{{{- else }}}
    {{{ toCamelCase .GoField }}}?: {{{goToTsType .GoType}}};
{{{- end }}}
{{{- end }}}
{{{- end }}}
}
// 添加编辑
export type type_{{{.ModuleName}}}_edit = {
{{{- range .Columns }}}
{{{- if or .IsEdit .IsInsert }}}
    {{{ toCamelCase .GoField }}}?: {{{goToTsType .GoType}}};
{{{- end }}}
{{{- end }}}
}

// {{{.FunctionName}}}列表
export function {{{.ModuleName}}}_list(params?: type_{{{.ModuleName}}}_query) {
    return request.get<Pages<type_{{{.ModuleName}}}>>({ url: '/{{{.ModuleName}}}/list', params })
}
// {{{.FunctionName}}}列表-所有
export function {{{.ModuleName}}}_list_all(params?: type_{{{.ModuleName}}}_query) {
    return request.get<type_{{{.ModuleName}}}[]>({ url: '/{{{.ModuleName}}}/listAll', params })
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}_detail({{{ .PrimaryKey }}}: number | string) {
    return request.get<type_{{{.ModuleName}}}>({ url: '/{{{.ModuleName}}}/detail', params: { {{{ .PrimaryKey }}} } })
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
export function {{{.ModuleName}}}_delete({{{ .PrimaryKey }}}: number | string) {
    return request.post<null>({ url: '/{{{.ModuleName}}}/del', data: { {{{ .PrimaryKey }}} } })
}

// {{{.FunctionName}}}导入
export const {{{.ModuleName}}}_import_file = '/{{{.ModuleName}}}/ImportFile'

// {{{.FunctionName}}}导出
export function {{{.ModuleName}}}_export_file(params: any) {
    return (window.location.href =`${config.baseUrl}${config.urlPrefix}/{{{.ModuleName}}}/ExportFile?token=${getToken()}&` + queryString.stringify(params))
}
