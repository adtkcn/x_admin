import { request } from '@/utils/request' 
import type { Pages } from '@/utils/request'
import { clearObjEmpty } from "@/utils/utils";

export type type_{{{.ModuleName}}} = {
{{{- range .Columns }}}
    {{{.TsField }}}: {{{.TsType}}}{{{ if eq .IsRequired 0 }}}|null{{{ end }}};
{{{- end }}}
}
// 查询
export type type_{{{.ModuleName}}}_query = {
{{{- range .Columns }}}
{{{- if .IsQuery }}}
{{{- if eq .HtmlType "datetime" }}}
    {{{.TsField }}}_start?: string;
    {{{.TsField }}}_end?: string;
{{{- else }}}
    {{{.TsField }}}?: {{{.TsType}}};
{{{- end }}}
{{{- end }}}
{{{- end }}}
}
// 添加编辑
export type type_{{{.ModuleName}}}_edit = {
{{{- range .Columns }}}
{{{- if or .IsEdit .IsInsert }}}
    {{{.TsField }}}?: {{{.TsType}}}{{{ if eq .IsRequired 0 }}}|null{{{ end }}};
{{{- end }}}
{{{- end }}}
}


// {{{.FunctionName}}}列表
export function {{{.ModuleName}}}_list(params?: type_{{{.ModuleName}}}_query) {
    return request<Pages<type_{{{.ModuleName}}}>>({
		url: '/{{{.ModuleName}}}/list',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}
// {{{.FunctionName}}}列表-所有
export function {{{.ModuleName}}}_list_all(params?: type_{{{.ModuleName}}}_query) {
    return request<type_{{{.ModuleName}}}[]>({
		url: '/{{{.ModuleName}}}/list_all',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}_detail({{{.PrimaryTsField }}}: {{{.PrimaryTsType}}}) {
    return request<type_{{{.ModuleName}}}>({
		url: '/{{{.ModuleName}}}/detail',
		method: 'GET',
		data:  { {{{.PrimaryTsField }}} }
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
export function {{{.ModuleName}}}_delete({{{.PrimaryTsField }}}: {{{.PrimaryTsType}}}) {
    return request<null>({
        url: '/{{{.ModuleName}}}/del',
        method: "POST",
        data:{
             {{{.PrimaryTsField }}} 
        },
    });
}