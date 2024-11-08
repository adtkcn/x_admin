
# nginx配置


## 前端页面使用history模式需要重定向
```nginx

location / {
    index /index.html;
    try_files $uri $uri/ /index.html;
}
# 找不到静态文件返回404
location ~* \.(?:html|css|js|png|jpg|jpeg|gif|webp|pdf|mp4|mp3|aac|ico|svg|woff|woff2|ttf|eot)$ {
    try_files $uri =404;
}
```