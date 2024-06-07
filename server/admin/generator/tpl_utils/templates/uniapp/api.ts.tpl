import { request } from '@/utils/request' 
import type { Pages } from '@/utils/request'
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
    return request<Pages<type_{{{.ModuleName}}}>>({
		url: '/{{{.ModuleName}}}/list',
		method: 'GET',
		data: params
	})
}
// {{{.FunctionName}}}列表-所有
export function {{{.ModuleName}}}_list_all(params?: type_{{{.ModuleName}}}_query) {
    return request<type_{{{.ModuleName}}}[]>({
		url: '/{{{.ModuleName}}}/listAll',
		method: 'GET',
		data: params
	})
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}_detail({{{ .PrimaryKey }}}: number | string) {
    return request<type_{{{.ModuleName}}}>({
		url: '/{{{.ModuleName}}}/detail',
		method: 'GET',
		data:  { {{{ .PrimaryKey }}} }
	})
}

// {{{.FunctionName}}}新增
export function {{{.ModuleName}}}_add(data: type_{{{.ModuleName}}}_edit) {
    return request<null>({
        url: '/{{{.ModuleName}}}/add',
        method: "POST",
        data,
    });
}

// {{{.FunctionName}}}编辑
export function {{{.ModuleName}}}_edit(data: type_{{{.ModuleName}}}_edit) {
    return request<null>({
        url: '/{{{.ModuleName}}}/edit',
        method: "POST",
        data,
    });
}

// {{{.FunctionName}}}删除
export function {{{.ModuleName}}}_delete({{{ .PrimaryKey }}}: number | string) {
    return request<null>({
        url: '/{{{.ModuleName}}}/del',
        method: "POST",
        data:{
             {{{ .PrimaryKey }}} 
        },
    });
}