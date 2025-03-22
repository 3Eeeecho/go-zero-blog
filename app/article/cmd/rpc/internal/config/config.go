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
		PageSize int
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int
	}
	CustomRedis        redis.RedisConf
	TagServiceRpcConf  zrpc.RpcClientConf
	UserServiceRpcConf zrpc.RpcClientConf
}
