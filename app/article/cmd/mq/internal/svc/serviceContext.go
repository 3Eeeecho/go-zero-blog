package svc

import (
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/mq/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config            config.Config
	ArticleServiceRpc articleservice.ArticleService
	ArticleModel      model.BlogArticleModel
	MQConn            *amqp091.Connection
	MQChannel         *amqp091.Channel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 MySQL 连接
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 初始化 RabbitMQ 连接
	conn, err := amqp091.Dial(c.RabbitMQ.URL)
	if err != nil {
		panic("failed to connect RabbitMQ: " + err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		panic("failed to open RabbitMQ channel: " + err.Error())
	}

	// 声明队列
	_, err = ch.QueueDeclare(
		c.RabbitMQ.SubmissionQueue, // 队列名称
		true,                       // 持久化
		false,                      // 自动删除
		false,                      // 独占
		false,                      // 不等待
		nil,                        // 额外参数
	)
	if err != nil {
		panic("failed to declare submission queue: " + err.Error())
	}
	return &ServiceContext{
		Config:            c,
		ArticleServiceRpc: articleservice.NewArticleService(zrpc.MustNewClient(c.ArticleServiceRpcConf)),
		ArticleModel:      model.NewBlogArticleModel(db),
		MQConn:            conn,
		MQChannel:         ch,
	}
}
