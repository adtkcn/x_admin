import type { AxiosRequestConfig, AxiosError } from 'axios'

import 'axios'
declare module 'axios' {
    // 扩展 RouteMeta
    interface AxiosRequestConfig {
        axiosHooks?: AxiosHooks
        requestOptions?: RequestOptions
    }
}

export interface RequestOptions {
    isParamsToData: boolean
    isReturnDefaultResponse: boolean
    isTransformResponse: boolean
    urlPrefix: string
    ignoreCancelToken: boolean
    withToken: boolean
    isOpenRetry: boolean
    retryCount: number
}

export interface AxiosHooks {
    requestInterceptorsHook?: (config: AxiosRequestConfig) => AxiosRequestConfig
    requestInterceptorsCatchHook?: (error: Error) => void
    responseInterceptorsHook?: (response: any) => any
    responseInterceptorsCatchHook?: (error: AxiosError) => void
}

export interface RequestData<T = any> {
    code: number
    data: T
    msg: string
    show: boolean
}
