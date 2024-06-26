FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
# 移动到工作目录：/home/kuroko 这个目录 是你项目代码 放在linux上
WORKDIR /home/kuroko
# 将代码复制到容器中
COPY . .
# 将我们的代码编译成二进制可执行文件  可执行文件名为 app
RUN go build -o app ./cmd
# 移动到用于存放生成的二进制文件的 /dist 目录
# 创建 /dist 目录并将二进制文件和静态文件复制到该目录
RUN mkdir -p /dist/web/template \
    && cp -r ./web/template/* /dist/web/template/ \
    && cp ./app /dist/app
WORKDIR /dist
# 将二进制文件从 /home/kuroko 目录复制到这里
RUN cp /home/kuroko/app .
# 创建 /dist 目录并将二进制文件和静态文件复制到该目录
# 声明服务端口
EXPOSE 8080
# 启动容器时运行的命令
CMD ["./app"]