#!/bin/sh
# 后台启动所有服务
./data/server/usercenter-rpc -f app/usercenter/cmd/rpc/etc/usercenter.yaml &
./data/server/usercenter-api -f app/usercenter/cmd/api/etc/user-api.yaml &
./data/server/tag-rpc -f app/tag/cmd/rpc/etc/tag.yaml &
./data/server/tag-api -f app/tag/cmd/api/etc/tag-api.yaml &
./data/server/article-rpc -f app/article/cmd/rpc/etc/article.yaml &
./data/server/article-api -f app/article/cmd/api/etc/article-api.yaml &
./data/server/article-mq -f app/article/cmd/mq/etc/article.yaml &

# 保持容器运行
wait