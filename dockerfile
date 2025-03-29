# 构建阶段
FROM golang:alpine AS builder

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata git

WORKDIR /app
COPY . .

# 安装 modd 并编译所有服务到 /data 目录
RUN go install github.com/cortesi/modd/cmd/modd@latest

# 启动 modd
CMD ["modd"]