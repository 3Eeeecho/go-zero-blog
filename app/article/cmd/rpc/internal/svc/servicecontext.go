package svc

import (
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/usercenter"
	"github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config           config.Config
	ArticleModel     model.BlogArticleModel
	CommentModel     model.BlogCommentModel
	ArticleLikeModel model.BlogArticleLikeModel
	TagRpc           tagservice.TagService
	UserRpc          usercenter.Usercenter
	Redis            *redis.Redis
	MQConn           *amqp091.Connection
	MQChannel        *amqp091.Channel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	rdb := redis.MustNewRedis(c.CustomRedis)

	conn, err := amqp091.Dial(c.RabbitMQ.URL)
	if err != nil {
		panic("failed to connect RabbitMQ: " + err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		panic("failed to open RabbitMQ channel: " + err.Error())
	}

	_, err = ch.QueueDeclare(c.RabbitMQ.SubmissionQueue, true, false, false, false, nil)
	if err != nil {
		panic("failed to declare submission queue: " + err.Error())
	}

	return &ServiceContext{
		Config:           c,
		ArticleModel:     model.NewBlogArticleModel(db),
		CommentModel:     model.NewBlogCommentModel(db),
		ArticleLikeModel: model.NewBlogArticleLikeModel(db),
		TagRpc:           tagservice.NewTagService(zrpc.MustNewClient(c.TagServiceRpcConf)),
		UserRpc:          usercenter.NewUsercenter(zrpc.MustNewClient(c.UserServiceRpcConf)),
		Redis:            rdb,
		MQConn:           conn,
		MQChannel:        ch,
	}
}

func (s *ServiceContext) Close() {
	s.MQChannel.Close()
	s.MQConn.Close()
}
