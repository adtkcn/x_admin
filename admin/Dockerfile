FROM node:18-alpine as build

# 设置工作目录为/app
WORKDIR /app

# 提前提高编译速度
COPY package.json .npmrc ./
# 安装依赖
RUN npm i

# 清理不需要的依赖项和文件
# RUN npm prune --production
COPY . /app

RUN npm run prod


FROM nginx:alpine
# 从.docker-compose/nginx/conf.d/目录拷贝my.conf到容器内的/etc/nginx/conf.d/my.conf
COPY .docker-compose/my.conf /etc/nginx/conf.d/my.conf

# 从第一阶段进行拷贝文件
COPY --from=0 /app/dist /usr/share/nginx/html

# 查看/etc/nginx/nginx.conf文件
RUN cat /etc/nginx/nginx.conf

# 查看 /etc/nginx/conf.d/my.conf
RUN cat /etc/nginx/conf.d/my.conf

# 查看 文件是否拷贝成功
RUN ls -al /usr/share/nginx/html

EXPOSE 8080
