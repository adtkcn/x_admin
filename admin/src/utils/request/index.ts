import { merge } from 'lodash-es'
import configs from '@/config'
import { Axios } from './axios'
import { ContentTypeEnum, RequestCodeEnum } from '@/enums/requestEnums'
import type { AxiosHooks, NewAxiosRequestConfig, NewInternalAxiosRequestConfig } from './type'
import { clearAuthInfo, getToken } from '../auth'
import feedback from '../feedback'
import NProgress from 'nprogress'
import { AxiosError } from 'axios'
// AxiosRequestConfig
import type { AxiosResponse } from 'axios'

import router from '@/router'
import { PageEnum } from '@/enums/pageEnum'

export interface Pages<T> {
    count: number
    lists: T[]
    pageNo: number
    pageSize: number
}
export interface Response<T> {
    code: number
    message: string
    data: T
}

// 处理axios的钩子函数
const axiosHooks: AxiosHooks = {
    requestInterceptorsHook(config) {
        NProgress.start()

        // const params = config.params || {}
        const headers = config.headers || {}

        // 添加token
        const token = getToken()
        if (token) {
            headers.token = token
        }

        config.headers = headers
        return config
    },
    requestInterceptorsCatchHook(err) {
        NProgress.done()
        return err
    },
    async responseInterceptorsHook(response: AxiosResponse) {
        console.log('返回Hook', response)

        NProgress.done()
        const config: NewInternalAxiosRequestConfig = response.config
        const { isTransformResponse, isReturnDefaultResponse } = config.requestOptions

        //返回默认响应，当需要获取响应头及其他数据时可使用
        if (isReturnDefaultResponse) {
            return response
        }
        // 是否需要对数据进行处理
        if (!isTransformResponse) {
            return response.data
        }
        const { code, data, message } = response.data
        switch (code) {
            case RequestCodeEnum.SUCCESS:
                return data

            case RequestCodeEnum.PARAMS_TYPE_ERROR:
            case RequestCodeEnum.PARAMS_VALID_ERROR:
            case RequestCodeEnum.REQUEST_METHOD_ERROR:
            case RequestCodeEnum.ASSERT_ARGUMENT_ERROR:
            case RequestCodeEnum.ASSERT_MYBATIS_ERROR:
            case RequestCodeEnum.LOGIN_ACCOUNT_ERROR:
            case RequestCodeEnum.LOGIN_DISABLE_ERROR:
            case RequestCodeEnum.NO_PERMISSTION:
            case RequestCodeEnum.FAILED:
            case RequestCodeEnum.SYSTEM_ERROR:
                message && feedback.msgError(message)
                return Promise.reject(data)

            case RequestCodeEnum.TOKEN_INVALID:
            case RequestCodeEnum.TOKEN_EMPTY:
                clearAuthInfo()
                router.push(PageEnum.LOGIN)
                return Promise.reject()

            default:
                return data
        }
    },

    /**
     * 错误处理
     * @param error
     * @returns
     */
    responseInterceptorsCatchHook(error) {
        console.log('返回异常Hook', error)

        NProgress.done()
        if (error.code !== AxiosError.ERR_CANCELED) {
            error.message && feedback.msgError(error.message)
        }
        return Promise.reject(error)
    }
}

const defaultOptions: NewAxiosRequestConfig = {
    timeout: configs.timeout,
    // 基础接口地址
    baseURL: configs.baseUrl,
    headers: { 'Content-Type': ContentTypeEnum.JSON, version: configs.version }

    // // 处理 axios的钩子函数
    // axiosHooks: axiosHooks,
    // 每个接口可以单独配置
}
const requestOptions = {
    //是否返回默认的响应
    isReturnDefaultResponse: false,
    // 需要对返回数据进行处理
    isTransformResponse: true,
    // 接口拼接地址
    urlPrefix: configs.urlPrefix
}

function createAxios(opt?: Partial<NewAxiosRequestConfig>) {
    return new Axios(
        // 深度合并
        merge(defaultOptions, opt || {}),
        requestOptions,
        axiosHooks
    )
}
const request = createAxios()
export default request
