import request from '@/utils/request'
import type { Pages } from '@/utils/request'

// 系统日志列表参数
export type type_system_log_operate = {
    title?: string
    email?: string
    ip?: string
    type?: string
    status?: number
    url?: string
    startTime?: string
    endTime?: string
}

// 系统日志返回信息
export type type_system_log_resp = {
    id: string
    email: string
    nickname: string
    type: string
    title: string
    method: string
    ip: string
    url: string
    args: string
    error: string
    status: number
    taskTime: string
    startTime: string
    endTime: string
    createTime: string
}

// 获取系统环境
export function systemInfo() {
    return request.get({ url: '/monitor/server' })
}

// 获取系统日志列表
export function systemLogLists(params: type_system_log_operate) {
    return request.get<Pages<type_system_log_resp>>({ url: '/system/log/operate', params })
}

export type type_info = Record<string, string | number>
export type type_commandStats = {
    name: string
    value: string
}
export type type_dbSize = number

export type SystemCache = {
    info: type_info
    commandStats: type_commandStats[]
    dbSize: type_dbSize
}
// 系统缓存监控
export function systemCache() {
    return request.get<SystemCache>({ url: '/monitor/cache' })
}
