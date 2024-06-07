import { request } from '@/utils/request' 
import type { Pages } from '@/utils/request'
type monitor_project= {
    id?: number;
    projectName?: string;
    projectType?: string;
    projectKey?: string;
    updateTime?: string;
    createTime?: string;
}
// 监控项目列表
export function monitor_project_list(params?: monitor_project) {
    return request<Pages<monitor_project>>({
		url: '/monitor_project/list',
		method: 'GET',
		data: params
	})
}
// 监控项目列表-所有
export function monitor_project_list_all(params?: monitor_project) {
    return request<monitor_project[]>({
		url: '/monitor_project/listAll',
		method: 'GET',
		data: params
	})
}

// 监控项目详情
export function monitor_project_detail(id: number | string) {
    return request<monitor_project>({
		url: '/monitor_project/detail',
		method: 'GET',
		data:  { id }
	})
}

// 监控项目新增
export function monitor_project_add(data: monitor_project) {
    return request<null>({
        url: '/monitor_project/add',
        method: "POST",
        data,
    });
}

// 监控项目编辑
export function monitor_project_edit(data: monitor_project) {
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