import { request } from '@/utils/request' 
import { clearObjEmpty } from '@/utils/utils' 
import type { Pages } from '@/utils/request'

// 字段与后端 monitor_schema 的 json tag 保持一致（snake_case）
export type type_monitor_project = {
    id?: string;
    project_key?: string;
    project_name?: string;
    project_type?: string;
    status?: number;
    is_delete?: number;
    create_time?: string;
    update_time?: string;
}
// 查询（对应 MonitorProjectListReq 支持的过滤字段）
export type type_monitor_project_query = {
    project_key?: string;
    project_name?: string;
    project_type?: string;
    status?: number;
    create_time_start?: string;
    create_time_end?: string;
    update_time_start?: string;
    update_time_end?: string;
}
// 添加编辑（对应 MonitorProjectAddReq / MonitorProjectEditReq）
export type type_monitor_project_edit = {
    id?: string;
    project_key?: string;
    project_name?: string;
    project_type?: string;
    status?: number;
}


// 监控项目列表
export function monitor_project_list(params?: type_monitor_project_query) {
    return request<Pages<type_monitor_project>>({
		url: '/monitor_project/list',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}
// 监控项目列表-所有
export function monitor_project_list_all(params?: type_monitor_project_query) {
    return request<type_monitor_project[]>({
		url: '/monitor_project/list_all',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}

// 监控项目详情
export function monitor_project_detail(id: string) {
    return request<type_monitor_project>({
		url: '/monitor_project/detail',
		method: 'GET',
		data:  { id }
	})
}

// 监控项目新增
export function monitor_project_add(data: type_monitor_project_edit) {
    return request<null>({
        url: '/monitor_project/add',
        method: "POST",
        data,
    });
}

// 监控项目编辑
export function monitor_project_edit(data: type_monitor_project_edit) {
    return request<null>({
        url: '/monitor_project/edit',
        method: "POST",
        data,
    });
}

// 监控项目删除
export function monitor_project_delete(id: string) {
    return request<null>({
        url: '/monitor_project/del',
        method: "POST",
        data:{
             id 
        },
    });
}
