package svc

import (
	"github.com/3Eeeecho/go-zero-blog/internal/config"
	"github.com/3Eeeecho/go-zero-blog/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config       config.Config
	ArticleModel model.BlogArticleModel
	TagModel     model.BlogTagModel
	UserModel    model.BlogUserModel
	Redis        *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	rdb := redis.MustNewRedis(c.Redis)
	// cacheConf := cache.CacheConf{
	// 	{
	// 		RedisConf: c.Redis, //缓存存储后端
	// 		Weight:    100,     // 节点的权重，用于负载均衡
	// 	},
	// }

	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewBlogArticleModel(db),
		TagModel:     model.NewBlogTagModel(db),
		UserModel:    model.NewBlogUserModel(db),
		Redis:        rdb,
	}
}
