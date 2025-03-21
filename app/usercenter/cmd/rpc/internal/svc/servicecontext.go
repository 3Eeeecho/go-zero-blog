package svc

import (
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.BlogUserModel
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	rdb := redis.MustNewRedis(c.CustomRedis)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewBlogUserModel(db),
		Redis:     rdb,
	}
}
