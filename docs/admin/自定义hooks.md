# 自定义hooks

1. 获取分页数据 usePaging

```js
import { usePaging } from '@/hooks/usePaging'
/**
 *
 * @param {object} options
 * @param {number} options.page 初始页码
 * @param {number} options.size 每页条数
 * @param {function} options.fetchFun 分页接口函数
 * @param {object} options.params 分页参数
 * @param {boolean} options.firstLoading 是否首次加载
 * 
 */
const { pager, getLists, resetPage, resetParams } = usePaging<any>(options)
// any为接口返回的列表数据类型

// pager.page 页码
// pager.size 每页条数
// pager.loading 是否加载中
// pager.count 总条数
// pager.lists 数据列表

// getLists() 获取数据列表
// resetPage() 重置页码并获取数据列表
// resetParams() 重置参数、页码并获取数据列表


```
2. 获取字典 useDictData
```js
import { useDictData } from '@/hooks/useDictOptions'
const { dictData } = useDictData<{
    flow_apply_status: any[]
    flow_group: any[]
}>(['flow_apply_status', 'flow_group'])

// 返回类型
// dictData.flow_apply_status: any[];
// dictData.flow_group: any[];
 
```

3. 获取不分页列表 useDictOptions
```ts
import { useDictOptions } from '@/hooks/useDictOptions'
const { optionsData,refresh } = useDictOptions<{
    articleCate: any[]
}>({
    articleCate: {
        api: articleCateAll
    }
})

// optionsData.articleCate: any[] 数据列表
// refresh 刷新数据
```

4. 函数防重锁 useLockFn
```ts
import { useLockFn } from '@/hooks/useLockFn'
const { lockFn, isLock } = useLockFn(asyncFn)

// asyncFn: 需要防止并发执行的异步函数
// lockFn: 包装后的函数，执行中再次调用会被忽略（避免重复提交）
// isLock: 当前是否处于锁定（执行中）状态，可用于禁用按钮
```

5. 带重置的响应式对象 useReactiveWithReset
```ts
import { useReactiveWithReset } from '@/hooks/useReactiveWithReset'
const { state, reset, setState } = useReactiveWithReset<{ name: string; age: number }>({
    name: '',
    age: 0
})

// state: reactive 响应式对象，直接读写
// reset(): 重置回初始值（深拷贝，避免引用污染）
// setState(info): 先 reset 再合并新数据
```

6. 监听路由变化 useWatchRoute
```ts
import { useWatchRoute } from '@/hooks/useWatchRoute'
const { route } = useWatchRoute((route) => {
    // route 变化（含首次 immediate）时触发
})

// route: 当前路由对象（RouteLocationNormalizedLoaded）
// callback 在路由变化时立即调用（immediate: true）
```

7. 全局 WebSocket useGlobalWs
```ts
import {
    useGlobalWs,
    onWsMessage,
    initGlobalWs,
    destroyGlobalWs
} from '@/hooks/useGlobalWs'

// 在 layout 中初始化唯一连接（单例）
initGlobalWs()

// 组件内订阅指定类型消息，卸载时自动取消订阅
onWsMessage('notice', (msg) => {
    // msg.type / msg.data，与后端 WsResponse { type, data } 对应
})
// 传 '*' 可订阅所有消息

// 获取连接状态与发送能力
const { status, send, close } = useGlobalWs()
// status: 'connecting' | 'open' | 'closed'（readonly）
// send(data): 发送文本/二进制消息
// close(): 关闭连接

// 销毁全局连接
destroyGlobalWs()
```
