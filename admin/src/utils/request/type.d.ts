import type { AxiosRequestConfig, AxiosError } from 'axios'

import 'axios'
declare module 'axios' {
    // 扩展 RouteMeta
    interface AxiosRequestConfig {
        requestOptions?: RequestOptions
    }
}

export interface RequestOptions {
    isReturnDefaultResponse: boolean
    isTransformResponse: boolean
    urlPrefix: string
}

export interface AxiosHooks {
    requestInterceptorsHook?: (config: AxiosRequestConfig) => AxiosRequestConfig
    requestInterceptorsCatchHook?: (error: Error) => void
    responseInterceptorsHook?: (response: any) => any
    responseInterceptorsCatchHook?: (error: AxiosError) => void
}

export interface ResponseData<T = any> {
    code: number
    data: T
    msg: string
}
