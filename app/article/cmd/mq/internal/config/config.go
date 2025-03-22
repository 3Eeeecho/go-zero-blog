package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		Dsn string
	}
	RabbitMQ struct {
		URL             string
		SubmissionQueue string
	}
	ArticleServiceRpcConf zrpc.RpcClientConf
}
