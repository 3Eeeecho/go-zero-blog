package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Dsn         string // 数据库连接字符串
		TablePrefix string // 表前缀
	}
	CustomRedis redis.RedisConf // Redis 配置，使用 go-zero 内置结构体
	JwtAuth     struct {
		AccessSecret string // JWT 密钥
		AccessExpire int
	}
	Crypto struct {
		Key string
	}
}
