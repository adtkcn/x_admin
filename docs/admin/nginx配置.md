
# nginx配置


## 前端页面使用history模式需要重定向
```nginx
location / {
    index /index.html;
    try_files $uri $uri/ /index.html;
}
```