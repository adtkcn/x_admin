# @adtkcn/x_err_sdk

错误监控 SDK，按平台输出独立产物：`web`、`uni`（uniapp）、`react-native`。
公共的批量/定时上报逻辑在 `Base`，各平台只需实现 `IErrorEvent` 采集接口。

## Web

```html
<script src="/XErr.umd.js"></script>
<script>
    const xErr = new XErr.Base(
        {
            Dns: `${location.origin}/api`,
            Pid: 'e19e3be20de94f49b68fafb4c30668bc',
            Uid: ''
        },
        new XErr.Web({})
    )
    xErr.SetUid(2) //设置用户ID
</script>
```

ESM 引入：

```js
import { Base, Web } from '@adtkcn/x_err_sdk/web'
new Base({ Dns, Pid }, new Web({}))
```

## uniapp

需在 `main.js` 中、`App()` 注册之前实例化 SDK（SDK 通过包装全局 `App` 注入 `onError` / `onUnhandledRejection`）。

```js
import { Base, Uni } from '@adtkcn/x_err_sdk/uni'

const xErr = new Base(
    { Dns: 'https://xxx/api', Pid: 'xxxx', Uid: '' },
    new Uni({})
)
xErr.SetUid(2)
```

## react-native

错误监听基于全局 `ErrorUtils.setGlobalHandler`；RN 无原生同步存储，默认使用内存缓存，
可传入同步存储（如 mmkv）实现持久化：

```js
import { Base, ReactNative } from '@adtkcn/x_err_sdk/react-native'

const xErr = new Base(
    { Dns: 'https://xxx/api', Pid: 'xxxx', Uid: '' },
    new ReactNative({ storage: mySyncStorage }) // storage 可选
)
xErr.SetUid(2)
```

## 构建

```bash
npm run build   # 分别打包 web / uni / react-native 三套产物到 dist/
```
