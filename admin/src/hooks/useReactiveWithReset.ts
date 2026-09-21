import { reactive, toRaw } from 'vue'

export function useReactiveWithReset<T extends Record<string, any>>(init: T) {
    const state = reactive<T>({ ...init }) as T

    // 存一份纯净的初始值深拷贝（防止引用被污染）
    const pureInit = JSON.parse(JSON.stringify(init))

    const reset = () => {
        // 用深拷贝的纯净值覆盖
        Object.assign(state, JSON.parse(JSON.stringify(pureInit)))
    }
    //
    const setState = (info: T) => {
        reset()
        Object.assign(state, info)
    }

    return { state, reset, setState }
}
