import { request } from '@/utils/request' 
import type { Pages } from '@/utils/request'
import { clearObjEmpty } from "@/utils/utils";

export type type_system_log_sms = {
    Id?: number;
    Scene?: number;
    Mobile?: string;
    Content?: string;
    Status?: number;
    Results?: string;
    SendTime?: string;
    CreateTime?: string;
    UpdateTime?: string;
}
// 查询
export type type_system_log_sms_query = {
    Scene?: number;
    Mobile?: string;
    Content?: string;
    Status?: number;
    Results?: string;
    SendTimeStart?: string;
    SendTimeEnd?: string;
    CreateTimeStart?: string;
    CreateTimeEnd?: string;
    UpdateTimeStart?: string;
    UpdateTimeEnd?: string;
}
// 添加编辑
export type type_system_log_sms_edit = {
    Id?: number;
    Scene?: number;
    Mobile?: string;
    Content?: string;
    Status?: number;
    Results?: string;
    SendTime?: string;
}


// 系统短信日志列表
export function system_log_sms_list(params?: type_system_log_sms_query) {
    return request<Pages<type_system_log_sms>>({
		url: '/system_log_sms/list',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}
// 系统短信日志列表-所有
export function system_log_sms_list_all(params?: type_system_log_sms_query) {
    return request<type_system_log_sms[]>({
		url: '/system_log_sms/listAll',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}

// 系统短信日志详情
export function system_log_sms_detail(Id: number | string) {
    return request<type_system_log_sms>({
		url: '/system_log_sms/detail',
		method: 'GET',
		data:  { Id }
	})
}

// 系统短信日志新增
export function system_log_sms_add(data: type_system_log_sms_edit) {
    return request<null>({
        url: '/system_log_sms/add',
        method: "POST",
        data,
    });
}

// 系统短信日志编辑
export function system_log_sms_edit(data: type_system_log_sms_edit) {
    return request<null>({
        url: '/system_log_sms/edit',
        method: "POST",
        data,
    });
}

// 系统短信日志删除
export function system_log_sms_delete(Id: number | string) {
    return request<null>({
        url: '/system_log_sms/del',
        method: "POST",
        data:{
             Id 
        },
    });
}