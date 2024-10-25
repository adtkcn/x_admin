import request from '@/utils/request'
import { dictDataAll } from '@/api/setting/dict'
import { reactive, toRaw } from 'vue'

interface Options {
    [propName: string]: {
        api: PromiseFun
        params?: Record<string, any>
        transformData?(data: any): any
    }
}

export function useDictOptions<T = any>(options: Options) {
    const optionsData: any = reactive({})
    const optionsKey = Object.keys(options)
    const apiLists = optionsKey.map((key) => {
        const value = options[key]
        optionsData[key] = []
        return () => value.api(toRaw(value.params) || {})
    })

    const refresh = async () => {
        const res = await Promise.allSettled<Promise<any>>(apiLists.map((api) => api()))
        // console.log(res)

        res.forEach((item, index) => {
            const key = optionsKey[index]
            if (item.status == 'fulfilled') {
                const { transformData } = options[key]
                const data = transformData ? transformData(item.value) : item.value
                optionsData[key] = data
            }
        })
    }
    refresh()
    return {
        optionsData: optionsData as T,
        refresh
    }
}

export type type_dict = {
    color?: string
    createTime?: string
    id?: number
    name?: string
    remark?: string
    sort?: number
    status?: number
    typeId?: number
    updateTime?: string
    value?: string
}
export function useDictData<T = any>(dict: string[]) {
    const options: Options = {}
    for (const type of dict) {
        options[type] = {
            api: dictDataAll,
            params: {
                dictType: type
            }
        }
    }
    const { optionsData } = useDictOptions<T>(options)
    return {
        dictData: optionsData
    }
}

export function useListAllData<T = any>(paths: Record<string, string>) {
    const options: Options = {}
    for (const key in paths) {
        options[key] = {
            api: () => request.get({ url: paths[key], params: {} }),
            transformData: (data: any) => {
                console.log('data', data)

                return data
            }

            // params: {
            //     dictType: path
            // }
        }
    }
    const { optionsData } = useDictOptions<T>(options)
    return {
        listAllData: optionsData
    }
}
