package svc

import (
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.BlogArticleModel
	TagRpc       tagservice.TagService
	Redis        *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	rdb := redis.MustNewRedis(c.CustomRedis)

	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewBlogArticleModel(db),
		TagRpc:       tagservice.NewTagService(zrpc.MustNewClient(c.TagServiceRpcConf)),
		Redis:        rdb,
	}
}
