# 打包
## 方式一：直接打包
```bash
# 在linux打包
go build -o x_admin .
# 在window打包
go build -o x_admin.exe .
```


## 方式二：运行pack.bat文件打包
```
在window运行pack.bat文件, 可以夸平台打包，压缩打包后产物；请按需取消注释
```
# 部署

## 运行方式
```bash
# 开发运行（需源码）
go run main.go

# 直接运行已编译二进制（默认读取同目录 .env.yaml）
./x_admin              # linux
x_admin.exe            # windows

# 指定配置文件
./x_admin -env=.env.online.yaml
```

## 上传所有需要的文件
- 打包后的二进制文件（`x_admin` / `x_admin.exe`）
- `public/*`（静态资源，必须随二进制一起上传，否则前端资源缺失）
- `.env.yaml`（或线上用 `.env.online.yaml`，通过 `-env` 指定）

> 配置文件名默认 `.env.yaml`，不是 `.env`（详见 [环境变量](./环境变量.md)）。




## 管理进程
1. 推荐使用pm2管理进程,因为我是前端

```bash
# https://pm2.io/

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
pm2 stop x_admin
# 重启
pm2 restart x_admin
# 查看日志
pm2 log x_admin 
```