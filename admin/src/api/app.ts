import request from '@/utils/request'
export type Config = {
    webName: string
    webLogo: string
    webFavicon: string
    webBackdrop: string
    ossDomain: string
    copyright: string
}
// 配置
export function getConfig() {
    return request.get<Config>({ url: '/common/index/config' })
}

export type Visitor = {
    date: string[]
    list: number[]
}
export type Console = {
    time: string
    flow_todo: number
    deviceTotal: number
    deviceOnline: number
    alertMinor: number
    alertSerious: number
    alertFault: number
    todayUsers: number
    totalUsers: number
}
// 工作台主页（仅统计，定时刷新）
export function getWorkbench() {
    return request.get<Console>({ url: '/common/index/console' })
}
// 在线人数趋势曲线初始数据（后续增量由 WS 推送）
export function getOnlineTrend() {
    return request.get<Visitor>({ url: '/common/index/online-trend' })
}
