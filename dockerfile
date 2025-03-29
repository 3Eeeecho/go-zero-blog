# ========== 构建阶段 ==========
FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

# 安装必要的工具（git 用于 go mod，tzdata 用于时区）
RUN apk update --no-cache && \
    apk add --no-cache git tzdata

WORKDIR /app
COPY . .

# 下载依赖
RUN go mod download

# 安装 modd（热加载工具）
RUN go install github.com/cortesi/modd/cmd/modd@latest

# ========== 运行阶段 ==========
FROM golang:alpine

# 安装运行时依赖（git 用于 modd 获取代码，tzdata 用于时区）
RUN apk update --no-cache && \
    apk add --no-cache git tzdata

# 从构建阶段复制证书和时区数据
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/bin/modd /usr/local/bin/modd
COPY --from=builder /go/pkg/mod /go/pkg/mod  
# 复制应用程序代码（modd 需要监视的代码）
WORKDIR /app
COPY --from=builder /app .

# 设置时区
ENV TZ=Asia/Shanghai

# 暴露开发环境的端口
EXPOSE 8001 8002 8003 2001 2002 2003 2004

# 启动 modd 热加载
CMD ["modd"]