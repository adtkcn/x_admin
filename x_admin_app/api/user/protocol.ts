import { request } from '@/utils/request' 
import type { Pages } from '@/utils/request'
import { clearObjEmpty } from "@/utils/utils";

export type type_user_protocol = {
    Id?: number;
    Title?: string;
    Content?: string;
    Sort?: number;
    IsDelete?: number;
    CreateTime?: string;
    UpdateTime?: string;
    DeleteTime?: string;
}
// 查询
export type type_user_protocol_query = {
    Title?: string;
    Content?: string;
    Sort?: number;
    CreateTimeStart?: string;
    CreateTimeEnd?: string;
    UpdateTimeStart?: string;
    UpdateTimeEnd?: string;
}
// 添加编辑
export type type_user_protocol_edit = {
    Id?: number;
    Title?: string;
    Content?: string;
    Sort?: number;
}


// 用户协议列表
export function user_protocol_list(params?: type_user_protocol_query) {
    return request<Pages<type_user_protocol>>({
		url: '/user_protocol/list',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}
// 用户协议列表-所有
export function user_protocol_list_all(params?: type_user_protocol_query) {
    return request<type_user_protocol[]>({
		url: '/user_protocol/listAll',
		method: 'GET',
		data: clearObjEmpty(params)
	})
}

// 用户协议详情
export function user_protocol_detail(Id: number | string) {
    return request<type_user_protocol>({
		url: '/user_protocol/detail',
		method: 'GET',
		data:  { Id }
	})
}

// 用户协议新增
export function user_protocol_add(data: type_user_protocol_edit) {
    return request<null>({
        url: '/user_protocol/add',
        method: "POST",
        data,
    });
}

// 用户协议编辑
export function user_protocol_edit(data: type_user_protocol_edit) {
    return request<null>({
        url: '/user_protocol/edit',
        method: "POST",
        data,
    });
}

// 用户协议删除
export function user_protocol_delete(Id: number | string) {
    return request<null>({
        url: '/user_protocol/del',
        method: "POST",
        data:{
             Id 
        },
    });
}