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