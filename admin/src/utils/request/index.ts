import axios from 'axios'
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { merge } from 'lodash-es'
import configs from '@/config'
import { ContentTypeEnum, RequestCodeEnum } from '@/enums/requestEnums'
import { clearAuthInfo, getToken } from '../auth'
import feedback from '../feedback'
import NProgress from 'nprogress'
import router from '@/router'
import { PageEnum } from '@/enums/pageEnum'

// eslint-disable-next-line @typescript-eslint/no-explicit-any
interface ApiResponse<T = any> {
    code: number
    data: T
    message: string
}
export interface Pages<T> {
    count: number
    lists: T[]
    pageNo: number
    pageSize: number
}

export interface RequestOptions {
    /** 返回完整响应对象(含 headers 等) */
    isReturnDefaultResponse?: boolean
    /** 是否转换响应数据(提取 data) */
    isTransformResponse?: boolean
}

interface InternalConfig extends AxiosRequestConfig {
    requestOptions?: RequestOptions
}

// ========== 创建实例 ==========

const service: AxiosInstance = axios.create({
    timeout: configs.timeout,
    baseURL: configs.baseUrl + configs.urlPrefix,
    headers: {
        'Content-Type': ContentTypeEnum.JSON,
        version: configs.version
    }
})

// ========== 请求拦截 ==========

service.interceptors.request.use(
    (config) => {
        NProgress.start()
        const token = getToken()
        if (token) {
            config.headers = config.headers || {}
            config.headers.token = token
        }
        return config
    },
    (error) => {
        NProgress.done()
        return Promise.reject(error)
    }
)

// ========== 响应拦截 ==========

service.interceptors.response.use(
    (response: AxiosResponse<ApiResponse>) => {
        NProgress.done()
        const config = response.config as InternalConfig
        const options = config.requestOptions || {}

        // 返回原始响应
        if (options.isReturnDefaultResponse) {
            return response
        }

        // 不转换响应，直接返回 data
        if (!options.isTransformResponse) {
            return response.data
        }

        // 转换响应数据
        const { code, data, message } = response.data

        switch (code) {
            case RequestCodeEnum.SUCCESS:
                return data

            case RequestCodeEnum.TOKEN_INVALID:
            case RequestCodeEnum.TOKEN_EMPTY:
                clearAuthInfo()
                router.push(PageEnum.LOGIN)
                return Promise.reject(new Error(message || '登录已过期'))

            case RequestCodeEnum.PARAMS_VALID_ERROR:
                feedback.msgError(
                    (Array.isArray(data) ? data.join('、') : message) || '参数校验失败'
                )
                return Promise.reject(data)

            default:
                feedback.msgError(message || '请求失败')
                return Promise.reject(response.data)
        }
    },
    (error) => {
        NProgress.done()
        if (error.code === 'ERR_BAD_RESPONSE') {
            feedback.msgError('网络发生错误')
        } else if (error.code !== 'ERR_CANCELED') {
            feedback.msgError(error.message || '请求异常')
        }
        return Promise.reject(error)
    }
)

// ========== 请求方法 ==========

const defaultOptions: RequestOptions = {
    isReturnDefaultResponse: false,
    isTransformResponse: true
}

function request<T = unknown>(config: AxiosRequestConfig, options?: RequestOptions): Promise<T> {
    const opts = merge({}, defaultOptions, options)
    const requestConfig: InternalConfig = { ...config, requestOptions: opts }

    return service.request<unknown, T>(requestConfig)
}

function get<T = unknown>(config: AxiosRequestConfig, options?: RequestOptions): Promise<T> {
    return request<T>({ ...config, method: 'GET' }, options)
}

function post<T = unknown>(config: AxiosRequestConfig, options?: RequestOptions): Promise<T> {
    return request<T>({ ...config, method: 'POST' }, options)
}

export default { request, get, post, service }
