import { RequestMethodsEnum } from '@/enums/requestEnums'
import axios, { AxiosError } from 'axios'
import type {
    AxiosInstance,
    AxiosRequestConfig,
    AxiosResponse,
    InternalAxiosRequestConfig
} from 'axios'
import { isFunction, merge, cloneDeep } from 'lodash-es'

import type { RequestData, RequestOptions } from './type'

export class Axios {
    private axiosInstance: AxiosInstance
    private readonly config: AxiosRequestConfig
    private readonly options: RequestOptions
    constructor(config: AxiosRequestConfig) {
        this.config = config
        this.options = config.requestOptions
        this.axiosInstance = axios.create(config)
        this.setupInterceptors()
    }

    /**
     * @description 获取axios实例
     */
    getAxiosInstance() {
        return this.axiosInstance
    }

    /**
     * @description 设置拦截器
     */
    setupInterceptors() {
        if (!this.config.axiosHooks) {
            return
        }
        const {
            requestInterceptorsHook,
            requestInterceptorsCatchHook,
            responseInterceptorsHook,
            responseInterceptorsCatchHook
        } = this.config.axiosHooks
        this.axiosInstance.interceptors.request.use(
            (config) => {
                if (isFunction(requestInterceptorsHook)) {
                    config = requestInterceptorsHook(config) as InternalAxiosRequestConfig
                }
                return config
            },
            (err: Error) => {
                if (isFunction(requestInterceptorsCatchHook)) {
                    requestInterceptorsCatchHook(err)
                }
                return err
            }
        )
        this.axiosInstance.interceptors.response.use(
            (response: any) => {
                if (isFunction(responseInterceptorsHook)) {
                    response = responseInterceptorsHook(response)
                }
                return response
            },
            (err: AxiosError) => {
                if (isFunction(responseInterceptorsCatchHook)) {
                    responseInterceptorsCatchHook(err)
                }

                return Promise.reject(err)
            }
        )
    }

    /**
     * @description get请求
     */
    get<T = any>(
        config: Partial<AxiosRequestConfig>,
        options?: Partial<RequestOptions>
    ): Promise<T> {
        return this.request({ ...config, method: RequestMethodsEnum.GET }, options)
    }

    /**
     * @description post请求
     */
    post<T = any>(
        config: Partial<AxiosRequestConfig>,
        options?: Partial<RequestOptions>
    ): Promise<T> {
        return this.request({ ...config, method: RequestMethodsEnum.POST }, options)
    }

    /**
     * @description 请求函数
     */
    request<T = any>(
        config: Partial<AxiosRequestConfig>,
        options?: Partial<RequestOptions>
    ): Promise<any> {
        const opt: RequestOptions = merge({}, this.options, options)
        const axioxConfig: AxiosRequestConfig = {
            ...cloneDeep(config),
            requestOptions: opt
        }
        const { urlPrefix } = opt
        // 拼接请求前缀如api
        if (urlPrefix) {
            axioxConfig.url = `${urlPrefix}${config.url}`
        }
        return new Promise((resolve, reject) => {
            this.axiosInstance
                .request<any, AxiosResponse<RequestData<T>>>(axioxConfig)
                .then((res) => {
                    resolve(res)
                })
                .catch((err) => {
                    reject(err)
                })
        })
    }
}
