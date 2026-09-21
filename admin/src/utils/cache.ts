interface CacheData<T> {
    expire: number
    value: T
}

const cache = {
    key: 'x_admin_',
    //设置缓存(expire为缓存时效,单位秒)
    set<T>(key: string, value: T, expire?: number): void {
        key = this.getKey(key)
        const data: CacheData<T> = {
            expire: expire ? this.time() + expire : 0,
            value
        }

        try {
            window.localStorage.setItem(key, JSON.stringify(data))
        } catch (e) {
            console.error('缓存设置失败:', e)
        }
    },
    get<T>(key: string): T | null {
        key = this.getKey(key)
        try {
            const dataStr = window.localStorage.getItem(key)
            if (!dataStr) {
                return null
            }
            const data: CacheData<T> = JSON.parse(dataStr)
            const { value, expire } = data
            if (expire && expire < this.time()) {
                window.localStorage.removeItem(key)
                return null
            }
            return value
        } catch (e) {
            console.error('缓存读取失败:', e)
            return null
        }
    },
    //获取当前时间
    time() {
        return Math.round(new Date().getTime() / 1000)
    },
    remove(key: string) {
        key = this.getKey(key)
        window.localStorage.removeItem(key)
    },
    getKey(key: string) {
        return this.key + key
    }
}

export default cache
