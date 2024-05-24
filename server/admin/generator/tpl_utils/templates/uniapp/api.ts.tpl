import { request } from '@/utils/request' 


// {{{.FunctionName}}}列表
export function {{{.ModuleName}}}_list(params?: Record<string, any>) {
    return request({
		url: '/{{{.ModuleName}}}/list',
		method: 'GET',
		data: params
	})
}
// {{{.FunctionName}}}列表-所有
export function {{{.ModuleName}}}_list_all(params?: Record<string, any>) {
    return request({
		url: '/{{{.ModuleName}}}/listAll',
		method: 'GET',
		data: params
	})
}

// {{{.FunctionName}}}详情
export function {{{.ModuleName}}}_detail({{{ .PrimaryKey }}}: number | string) {
    return request({
		url: '/{{{.ModuleName}}}/detail',
		method: 'GET',
		data:  { {{{ .PrimaryKey }}} }
	})
}

// {{{.FunctionName}}}新增
export function {{{.ModuleName}}}_add(data: Record<string, any>) {
    return request({
        url: '/{{{.ModuleName}}}/add',
        method: "POST",
        data,
    });
}

// {{{.FunctionName}}}编辑
export function {{{.ModuleName}}}_edit(data: Record<string, any>) {
    return request({
        url: '/{{{.ModuleName}}}/edit',
        method: "POST",
        data,
    });
}

// {{{.FunctionName}}}删除
export function {{{.ModuleName}}}_delete({{{ .PrimaryKey }}}: number | string) {
    return request.post({ url: '/{{{.ModuleName}}}/del', { {{{ .PrimaryKey }}} } })
    return request({
        url: '/{{{.ModuleName}}}/del',
        method: "POST",
        data:{
             {{{ .PrimaryKey }}} 
        },
    });
}