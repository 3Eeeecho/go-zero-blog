#!/bin/bash

# 定义服务路径
RPC_PATH="/home/eecho/go/src/go-zero-blog/app/tag/cmd/rpc/tag.go"
API_PATH="/home/eecho/go/src/go-zero-blog/app/tag/cmd/api/tag.go"

# 定义配置文件路径
RPC_CONFIG="/home/eecho/go/src/go-zero-blog/app/tag/cmd/rpc/etc/tag.yaml"
API_CONFIG="/home/eecho/go/src/go-zero-blog/app/tag/cmd/api/etc/tag-api.yaml"

# 启动 rpc 服务
gnome-terminal --title="Tag RPC" -- bash -c "go run $RPC_PATH -f $RPC_CONFIG; exec bash"

# 启动 api 服务
gnome-terminal --title="Tag API" -- bash -c "go run $API_PATH -f $API_CONFIG; exec bash"

echo "Tag RPC 和 Tag API 服务已启动"

#------------------------------------------------------------------
# 定义服务路径
RPC_PATH="/home/eecho/go/src/go-zero-blog/app/usercenter/cmd/rpc/usercenter.go"
API_PATH="/home/eecho/go/src/go-zero-blog/app/usercenter/cmd/api/user.go"

# 定义配置文件路径
RPC_CONFIG="/home/eecho/go/src/go-zero-blog/app/usercenter/cmd/rpc/etc/usercenter.yaml"
API_CONFIG="/home/eecho/go/src/go-zero-blog/app/usercenter/cmd/api/etc/user-api.yaml"

# 启动 rpc 服务
gnome-terminal --title="User RPC" -- bash -c "go run $RPC_PATH -f $RPC_CONFIG; exec bash"

# 启动 api 服务
gnome-terminal --title="User API" -- bash -c "go run $API_PATH -f $API_CONFIG; exec bash"

echo "User RPC 和 User API 服务已启动"

#------------------------------------------------------------------
# 定义服务路径
RPC_PATH="/home/eecho/go/src/go-zero-blog/app/article/cmd/rpc/article.go"
API_PATH="/home/eecho/go/src/go-zero-blog/app/article/cmd/api/article.go"

# 定义配置文件路径
RPC_CONFIG="/home/eecho/go/src/go-zero-blog/app/article/cmd/rpc/etc/article.yaml"
API_CONFIG="/home/eecho/go/src/go-zero-blog/app/article/cmd/api/etc/article-api.yaml"

# 启动 rpc 服务
gnome-terminal --title="Article RPC" -- bash -c "go run $RPC_PATH -f $RPC_CONFIG; exec bash"

# 启动 api 服务
gnome-terminal --title="Article API" -- bash -c "go run $API_PATH -f $API_CONFIG; exec bash"

echo "Article RPC 和 Article API 服务已启动"