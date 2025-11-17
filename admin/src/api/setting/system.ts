import request from '@/utils/request'
export interface SystemLogResp {
    title: string
    username: string
    url: string
    ip: string
    error: string
    taskTime: string
    createTime: string
}
// 获取系统环境
export function systemInfo() {
    return request.get({ url: '/monitor/server' })
}

// 获取系统日志列表
export function systemLogLists(params: any) {
    return request.get({ url: '/system/log/operate', params })
}

// 系统缓存监控
export function systemCache() {
    return request.get({ url: '/monitor/cache' })
}
