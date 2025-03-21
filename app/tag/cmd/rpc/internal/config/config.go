package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Dsn         string
		TablePrefix string
	}
	App struct { // 应用相关配置
		PageSize       int
		ExportSavePath string
	}
	JwtAuth struct {
		AccessSecret string
	}
	CustomRedis redis.RedisConf
}
