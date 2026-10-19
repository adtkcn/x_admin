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

type Visitor = {
    date: string[]
    list: number[]
}
export type Console = {
    today: {
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
    visitor: Visitor
}
// 工作台主页
export function getWorkbench() {
    return request.get<Console>({ url: '/common/index/console' })
}
