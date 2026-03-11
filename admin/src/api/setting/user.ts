import request from '@/utils/request'

// 用户设置返回信息
export type type_setting_user_resp = {
    defaultAvatar: string
}

// 用户设置保存参数
export type type_setting_user_save = {
    defaultAvatar: string
}

// 登录设置返回信息
export type type_setting_login_resp = {
    loginWay: number[]
    forceBindMobile: number
    openAgreement: number
    openOtherAuth: number
    autoLoginAuth: number[]
}

// 登录设置保存参数
export type type_setting_login_save = {
    loginWay: number[]
    forceBindMobile: number
    openAgreement: number
    openOtherAuth: number
    autoLoginAuth: number[]
}

/**
 * @return { Promise }
 * @description 获取用户设置
 */
export function getUserSetup() {
    return request.get<type_setting_user_resp>({ url: '/setting/user/detail' })
}

/**
 * @return { Promise }
 * @param { type_setting_user_save } data 默认用户头像
 * @description 设置用户设置
 */
export function setUserSetup(data: type_setting_user_save) {
    return request.post({ url: '/setting/user/save', data })
}

/**
 * @return { Promise }
 * @description 设置登录注册规则
 */
export function getLogin() {
    return request.get<type_setting_login_resp>({ url: '/setting/login/detail' })
}

/**
 * @return { Promise }
 * @param { type_setting_login_save } data 登录设置
 * @description 设置登录注册规则
 */
export function setLogin(data: type_setting_login_save) {
    return request.post({ url: '/setting/login/save', data })
}
