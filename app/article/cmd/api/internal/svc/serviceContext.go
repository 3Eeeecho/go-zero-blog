package svc

import (
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	ArticleServiceRpc articleservice.ArticleService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		ArticleServiceRpc: articleservice.NewArticleService(zrpc.MustNewClient(c.ArticleServiceRpcConf)),
	}
}
