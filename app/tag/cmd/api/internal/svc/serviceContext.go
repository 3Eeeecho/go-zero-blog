package svc

import (
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	TagServiceRpc tagservice.TagService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		TagServiceRpc: tagservice.NewTagService(zrpc.MustNewClient(c.TagServiceRpcConf)),
	}
}
