import { request } from '@/utils/request' 
import type { Pages } from '@/utils/request'
export type type_monitor_project = {
    id?: number;
    projectKey?: string;
    projectName?: string;
    projectType?: string;
    isDelete?: number;
    createTime?: string;
    updateTime?: string;
    deleteTime?: string;
}
// 查询
export type type_monitor_project_query = {
    projectKey?: string;
    projectName?: string;
    projectType?: string;
    createTimeStart?: string;
    createTimeEnd?: string;
    updateTimeStart?: string;
    updateTimeEnd?: string;
}
// 添加编辑
export type type_monitor_project_edit = {
    id?: number;
    projectKey?: string;
    projectName?: string;
    projectType?: string;
}


// 监控项目列表
export function monitor_project_list(params?: type_monitor_project_query) {
    return request<Pages<type_monitor_project>>({
		url: '/monitor_project/list',
		method: 'GET',
		data: params
	})
}
// 监控项目列表-所有
export function monitor_project_list_all(params?: type_monitor_project_query) {
    return request<type_monitor_project[]>({
		url: '/monitor_project/listAll',
		method: 'GET',
		data: params
	})
}

// 监控项目详情
export function monitor_project_detail(id: number | string) {
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
export function monitor_project_delete(id: number | string) {
    return request<null>({
        url: '/monitor_project/del',
        method: "POST",
        data:{
             id 
        },
    });
}