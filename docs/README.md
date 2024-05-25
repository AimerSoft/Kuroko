<div align="center">
 <a href="https://toaru.huijiwiki.com/wiki/%E7%99%BD%E4%BA%95%E9%BB%91%E5%AD%90">
  <img src="img/kuroko.png" alt="Kuroko" width = "256">
 </a>
 <h1>Kuroko | 短链接服务 </h1>
 “白井的能力是大能力者的<b>空间移动</b>（日文：空間移動；英文：Teleport），简单来说，就是可以把皮肤碰触到的东西包括自己在内，以无视三维空间规则的方式瞬间送至远方的能力。”
</div>

## 如何部署

### [体验demo](https://kuroko.umb.ink/)
0. 部署前请准备好你的Redis服务，并保证它可用
1. 拉取本项目到你要部署的服务器
2. 编辑docker-compose.yml文件，添加你的短连接域名和redis相关信息,举例：


```yml
version: '3'

services:
  kuroko:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - REDIS_HOST=127.0.0.1
      - REDIS_PORT=6379
      - REDIS_PWD=123456
      - DOMAIN=https://umb.ink/k/ # 你期望的短连接前缀，我们一般使用域名+Nginx代理下

```
3. 配置nginx。这里给出一个简单的nginx部署示例
- 注意：服务默认接口为/
```yaml
server {
  listen 443 ssl http2;
  # 改为你的网址
  server_name umb.ink;
  # 修改你的证书的公私钥
  ssl_certificate /home/ubuntu/xxx.crt; # 你的ssl证书路径
  ssl_certificate_key /home/ubuntu/xxx.key;
  location /k/ {
    # 改为容器的 PORT
    proxy_pass http://localhost:8080/;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Upgrade $http_upgrade;
  }
 location /k {
    # 改为容器的 PORT
    proxy_pass http://localhost:8080/;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Upgrade $http_upgrade;
  }
}
```
4. 运行
```shell
docker-compose up --build -d
```
5. 检查

   访问服务根路径,比如你配置的是https://umb.ink/k/，则访问https://umb.ink/k/。
   若返回下面的页面，则正常

   ![](https://umb.ink/static/img/02a2fbf79ec078e46daf75e4b59f6e9f.clipboard-2024-05-26.webp)
```
# TODO: 
- 高优：web与server整合同时部署
- 支持Docker-Compose部署，集成自动部署本地Redis
- Emoji短连接生成
- 颜文字短连接生成
- 梗短连接生成
- 短连接生成远程配置
- 八卦短连接生成
```