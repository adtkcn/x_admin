import { request } from '@/utils/request' 
import { clearObjEmpty } from '@/utils/utils' 

import type { Pages } from '@/utils/request'

// 字段与后端 monitor_schema 的 json tag 保持一致（snake_case）
export type type_monitor_client = {
    id?: string;
    project_key?: string;
    client_id?: string;
    user_id?: string;
    os?: string;
    browser?: string;
    country?: string;
    province?: string;
    city?: string;
    operator?: string;
    ip?: string;
    ua?: string;
    create_time?: string;
    update_time?: string;
    width?: number;
    height?: number;
    is_delete?: number;
}
// 查询（对应 MonitorClientListReq 支持的过滤字段）
export type type_monitor_client_query = {
    project_key?: string;
    client_id?: string;
    user_id?: string;
    os?: string;
    browser?: string;
    city?: string;
    ua?: string;
    create_time_start?: string;
    create_time_end?: string;
}
// 添加编辑（对应 MonitorClientAddReq / MonitorClientEditReq）
export type type_monitor_client_edit = {
    id?: string;
    project_key?: string;
    client_id?: string;
    user_id?: string;
    os?: string;
    browser?: string;
    ip?: string;
    ua?: string;
    width?: number;
    height?: number;
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
		url: '/monitor_client/list_all',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}

// 监控-客户端信息详情
export function monitor_client_detail(id: string) {
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
export function monitor_client_delete(id: string) {
    return request<null>({
        url: '/monitor_client/del',
        method: "POST",
        data:{
             id 
        },
    });
}
