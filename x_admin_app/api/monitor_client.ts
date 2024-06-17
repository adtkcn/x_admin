import { request } from '@/utils/request' 
import { clearObjEmpty } from '@/utils/utils' 

import type { Pages } from '@/utils/request'
export type type_monitor_client = {
    id?: number;
    projectKey?: string;
    clientId?: string;
    userId?: string;
    os?: string;
    browser?: string;
    city?: string;
    width?: number;
    height?: number;
    ua?: string;
    createTime?: string;
    clientTime?: string;
}
// 查询
export type type_monitor_client_query = {
    projectKey?: string;
    clientId?: string;
    userId?: string;
    os?: string;
    browser?: string;
    city?: string;
    width?: number;
    height?: number;
    ua?: string;
    createTimeStart?: string;
    createTimeEnd?: string;
    clientTimeStart?: string;
    clientTimeEnd?: string;
}
// 添加编辑
export type type_monitor_client_edit = {
    id?: number;
    clientId?: string;
    userId?: string;
    os?: string;
    browser?: string;
    city?: string;
    width?: number;
    height?: number;
    ua?: string;
    clientTime?: string;
}


// 监控-客户端信息列表
export function monitor_client_list(params?: type_monitor_client_query) {
    return request<Pages<type_monitor_client>>({
		url: '/monitor_client/list',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}
// 监控-客户端信息列表-所有
export function monitor_client_list_all(params?: type_monitor_client_query) {
    return request<type_monitor_client[]>({
		url: '/monitor_client/listAll',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}

// 监控-客户端信息详情
export function monitor_client_detail(id: number | string) {
    return request<type_monitor_client>({
		url: '/monitor_client/detail',
		method: 'GET',
		data:  { id }
	})
}

// 监控-客户端信息新增
export function monitor_client_add(data: type_monitor_client_edit) {
    return request<null>({
        url: '/monitor_client/add',
        method: "POST",
        data,
    });
}

// 监控-客户端信息编辑
export function monitor_client_edit(data: type_monitor_client_edit) {
    return request<null>({
        url: '/monitor_client/edit',
        method: "POST",
        data,
    });
}

// 监控-客户端信息删除
export function monitor_client_delete(id: number | string) {
    return request<null>({
        url: '/monitor_client/del',
        method: "POST",
        data:{
             id 
        },
    });
}