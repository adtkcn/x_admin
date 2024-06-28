# 打包
## 方式一：直接打包
```bash
# linux
go build -o x_admin .
# window
go build -o x_admin.exe .
```

## 方式二：goreleaser 同时打包多平台
```bash
# 安装第三方工具
go install github.com/goreleaser/goreleaser@latest
# 打包
goreleaser release --snapshot --clean
```


# 部署

## 上传所有需要的文件
- 打包后的二进制文件.exe
- resources/*
- .env



## 管理进程
我是前端er，所以我推荐使用pm2管理进程
https://pm2.io/
```bash
# 需要node环境
npm install pm2 -g
# 之前有一种不依赖node直接安装，好像不能用了,没找见
```
```bash
# 启动
pm2 start 打包后的二进制文件名 --name x_admin
# 开机启动
pm2 startup
# 保存
pm2 save
# 所有任务列表
pm2 list


# 停止
```bash
pm2 stop x_admin
# 重启
```bash
pm2 restart x_admin
# 查看日志
pm2 log x_admin 
```