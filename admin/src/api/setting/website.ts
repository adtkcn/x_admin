import request from '@/utils/request'

// 版权信息参数
export type type_setting_copyright_item = {
    name: string
    link: string
}

// 版权信息返回信息
export type type_setting_copyright_resp = type_setting_copyright_item[]

// 网站设置参数
export type type_setting_website_save = {
    name?: string
    logo?: string
    favicon?: string
    backdrop?: string
    shopName?: string
    shopLogo?: string
}

// 网站设置返回信息
export type type_setting_website_resp = {
    name: string
    logo: string
    favicon: string
    backdrop: string
    shopName: string
    shopLogo: string
}

// 获取备案信息
export function getCopyright() {
    return request.get<type_setting_copyright_resp>({ url: '/setting/copyright/detail' })
}

// 设置备案信息
export function setCopyright(data: type_setting_copyright_item[]) {
    return request.post({ url: '/setting/copyright/save', data })
}

// 获取网站信息
export function getWebsite() {
    return request.get<type_setting_website_resp>({ url: '/setting/website/detail' })
}

// 设置网站信息
export function setWebsite(data: type_setting_website_save) {
    return request.post({ url: '/setting/website/save', data })
}

// 获取网站接口列表
export function getApiList() {
    return request.get({ url: '/apiList' })
}
