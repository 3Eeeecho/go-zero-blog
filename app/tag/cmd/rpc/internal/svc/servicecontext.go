package svc

import (
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/tag/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	TagModel model.BlogTagModel
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	rdb := redis.MustNewRedis(c.CustomRedis)

	return &ServiceContext{
		Config:   c,
		TagModel: model.NewBlogTagModel(db),
		Redis:    rdb,
	}
}
