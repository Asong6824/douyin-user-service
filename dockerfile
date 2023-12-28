# 使用官方Go镜像作为构建环境
FROM golang:latest as builder

# 设置工作目录
WORKDIR /app

# 使用Alpine镜像作为运行环境，减小镜像体积
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 设置容器启动时执行的命令
CMD ["/bin/sh"]
