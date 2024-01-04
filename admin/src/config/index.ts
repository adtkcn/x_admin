const config = {
    terminal: 1, //终端
    title: '后台管理系统', //网站默认标题
    version: '1.3.3', //版本号
    // import.meta.env.VITE_APP_BASE_URL ||
    baseUrl: '', //请求接口域名
    urlPrefix: '/api/admin', //请求默认前缀
    timeout: 60 * 3000 //请求超时时长
}

export default config
