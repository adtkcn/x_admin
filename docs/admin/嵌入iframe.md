## 嵌入iframe
1. 添加菜单
2. 选择组件路径：IframeComponent
3. 路由参数：
```
{"url": "/api/static/api/index.html"}或者 url=/api/static/api/index.html
```

## 路由同步说明

父应用（`IframeComponent.vue`）与子应用的 URL 会双向同步，四种场景表现如下：

| 场景 | 触发路径 | 行为 |
| --- | --- | --- |
| 主应用跳转 | 修改 `query.url` | iframe `:src` 变化，重新加载 |
| 子应用内部跳转 | 子应用 `postMessage` 上报 | 父应用 `router.replace` 更新 `query.url`，**不会**重新加载 iframe |
| 用户刷新浏览器 | `query.url` 仍在地址栏 | iframe 加载同一个 URL |
| 用户复制链接 | 地址栏带 `?url=xxx` | 分享后可还原到子应用当前页面 |

## 子应用接入

子应用无需引入任何 SDK，只要在**内部路由变化**时向父应用上报当前完整 URL 即可：

```js
// 运行在 iframe 内的子应用，路由变化时执行
window.parent.postMessage(
    { type: 'iframe-url-change', url: window.location.href },
    '*'
)
```

约定：
- 消息体的 `type` 必须为 `iframe-url-change`，`url` 为字符串（建议上报完整 URL）。
- 建议监听子应用自身的路由变化（如 vue-router 的 `afterEach` 钩子）后调用上述上报。
- 父应用内部用 `innerUrl` 标记区分「父应用主动跳转」与「子应用回写」，回写时只更新地址栏、不重载 iframe。

```js
// vue-router 子应用示例
router.afterEach(() => {
    window.parent.postMessage(
        { type: 'iframe-url-change', url: window.location.href },
        '*'
    )
})
```

> 提示：父应用当前未对 `event.origin` 做白名单校验，若子应用为跨域嵌入，请在 `handleMessage` 中增加来源校验。