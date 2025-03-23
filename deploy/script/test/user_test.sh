#!/bin/bash

# 定义颜色
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # 无颜色

# 项目根目录
BASE_DIR="/home/eecho/go/src/go-zero-blog"

# 函数：启动服务
start_service() {
    local service_name=$1
    local rpc_name=$2
    local api_name=$3

    # 根据服务名称设置文件名
    local rpc_file="${rpc_name:-$service_name}"
    local api_file="${api_name:-$service_name}"

    local rpc_path="${BASE_DIR}/app/${service_name}/cmd/rpc/${rpc_file}.go"
    local api_path="${BASE_DIR}/app/${service_name}/cmd/api/${api_file}.go"
    local rpc_config="${BASE_DIR}/app/${service_name}/cmd/rpc/etc/${rpc_file}.yaml"
    local api_config="${BASE_DIR}/app/${service_name}/cmd/api/etc/${api_file}-api.yaml"

    # 检查文件是否存在
    for file in "$rpc_path" "$api_path" "$rpc_config" "$api_config"; do
        if [ ! -f "$file" ]; then
            echo -e "${RED}错误: 文件 $file 不存在${NC}"
            return 1
        fi
    done

    # 启动 RPC 服务
    gnome-terminal --title="${service_name^} RPC" -- bash -c "go run '$rpc_path' -f '$rpc_config' || echo 'RPC 启动失败'; exec bash" &

    # 启动 API 服务
    gnome-terminal --title="${service_name^} API" -- bash -c "go run '$api_path' -f '$api_config' || echo 'API 启动失败'; exec bash" &

    echo -e "${GREEN}${service_name^} RPC 和 ${service_name^} API 服务已启动${NC}"
}

# 函数：启动 MQ Worker
start_mq_worker() {
    local mq_path="${BASE_DIR}/app/article/cmd/mq/article.go"
    local mq_config="${BASE_DIR}/app/article/cmd/mq/etc/article.yaml"

    # 检查文件是否存在
    for file in "$mq_path" "$mq_config"; do
        if [ ! -f "$file" ]; then
            echo -e "${RED}错误: 文件 $file 不存在${NC}"
            return 1
        fi
    done

    # 启动 MQ Worker
    gnome-terminal --title="Article MQ Worker" -- bash -c "go run '$mq_path' -f '$mq_config' || echo 'MQ Worker 启动失败'; exec bash" &

    echo -e "${GREEN}Article MQ Worker 已启动${NC}"
}

# 主函数：启动所有服务
main() {
    echo "开始启动服务..."
    echo "----------------------------------------"

    # 启动 tag 服务
    start_service "tag" "tag" "tag"
    if [ $? -ne 0 ]; then
        echo -e "${RED}启动 tag 服务失败${NC}"
        exit 1
    fi
    sleep 1

    # 启动 usercenter 服务（API 为 user.go，RPC 为 usercenter.go）
    start_service "usercenter" "usercenter" "user"
    if [ $? -ne 0 ]; then
        echo -e "${RED}启动 usercenter 服务失败${NC}"
        exit 1
    fi
    sleep 1

    # 启动 article 服务
    start_service "article" "article" "article"
    if [ $? -ne 0 ]; then
        echo -e "${RED}启动 article 服务失败${NC}"
        exit 1
    fi
    sleep 1

    # 启动 MQ Worker
    start_mq_worker
    if [ $? -ne 0 ]; then
        echo -e "${RED}启动 Article MQ Worker 失败${NC}"
        exit 1
    fi
    sleep 1

    echo "----------------------------------------"
    echo -e "${GREEN}所有服务启动完成${NC}"
}

# 执行主函数
main