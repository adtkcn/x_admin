import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 通知列表查询
export type type_notice_list = {
    type?: string
    is_read?: number // 0未读 1已读 -1全部
}

// 通知返回
export type type_notice_resp = {
    id: string
    type: string
    title: string
    content: string
    receiver_id: string
    sender_id: string
    url: string
    is_read: number
    read_time: string
    create_time: string
}

// 未读数量
export type type_notice_unread = {
    count: number
}

// 偏好设置：渠道清单（由后端定义，含当前开关状态）
export type type_notice_setting_channel = {
    key: string
    label: string
    enabled: number
}
export type type_notice_setting = {
    channels: type_notice_setting_channel[]
}
// 保存偏好请求：channel(渠道) -> is_enabled(开关)
export type type_notice_setting_save = {
    settings: Record<string, number>
}

// 通知列表
export function noticeList(params: type_notice_list & { pageNo: number; pageSize: number }) {
    return request.get<Pages<type_notice_resp>>({ url: '/system/notice/list', params })
}

// 未读数量
export function noticeUnreadCount() {
    return request.get<type_notice_unread>({ url: '/system/notice/unread_count' })
}

// 标记已读
export function noticeRead(data: { id: string }) {
    return request.post({ url: '/system/notice/read', data })
}

// 全部已读
export function noticeReadAll() {
    return request.post({ url: '/system/notice/read_all' })
}

// 删除通知
export function noticeDel(data: { id: string }) {
    return request.post({ url: '/system/notice/del', data })
}

// 获取偏好
export function noticeGetSetting() {
    return request.get<type_notice_setting>({ url: '/system/notice/setting' })
}

// 保存偏好
export function noticeSaveSetting(data: type_notice_setting_save) {
    return request.post({ url: '/system/notice/setting/save', data })
}
