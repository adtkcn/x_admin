import { reactive, toRaw } from 'vue'

// 分页钩子函数
interface Options {
    page?: number
    size?: number
    fetchFun: (_arg: any) => Promise<any>
    params?: Record<any, any>
    firstLoading?: boolean
}

/**
 *
 * @param {object} options
 * @param {number} options.page 页码
 * @param {number} options.size 每页条数
 * @param {function} options.fetchFun 分页接口函数
 * @param {object} options.params 分页参数
 * @param {boolean} options.firstLoading 是否首次加载
 *
 *
 */
export function usePaging<T>(options: Options) {
    const { page = 1, size = 15, fetchFun, params = {}, firstLoading = false } = options
    // 记录分页初始参数
    const paramsInit: Record<any, any> = Object.assign({}, toRaw(params))
    // 分页数据
    const pager = reactive({
        page,
        size,
        loading: firstLoading,
        count: 0,
        lists: [] as T[]
    })
    // 请求分页接口
    const getLists = () => {
        pager.loading = true
        return fetchFun({
            pageNo: pager.page,
            pageSize: pager.size,
            ...params
        })
            .then((res: any) => {
                pager.count = res?.count
                pager.lists = res?.lists
                return Promise.resolve(res)
            })
            .catch((err: any) => {
                return Promise.reject(err)
            })
            .finally(() => {
                pager.loading = false
            })
    }
    // 重置为第一页
    const resetPage = () => {
        pager.page = 1
        getLists()
    }
    // 重置参数
    const resetParams = () => {
        Object.keys(paramsInit).forEach((item) => {
            params[item] = paramsInit[item]
        })

        resetPage()
    }
    return {
        pager,
        getLists,
        resetParams,
        resetPage
    }
}
