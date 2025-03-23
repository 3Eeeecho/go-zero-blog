package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/mq/internal/config"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/mq/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/state"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/article.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	fmt.Printf("Starting RabbitMQ at %s...\n", c.RabbitMQ.URL)

	// 开始消费队列
	msgs, err := svcCtx.MQChannel.Consume(
		c.RabbitMQ.SubmissionQueue, // 队列名称
		"",                         // 消费者标签
		true,                       // 自动确认
		false,                      // 独占
		false,                      // 不本地
		false,                      // 不等待
		nil,                        // 额外参数
	)
	if err != nil {
		log.Fatalf("failed to consume submission queue: %v", err)
	}

	// 处理消息
	//TODO 现在是自动审核,全部通过,可以拓展逻辑,让管理员手动审核
	for msg := range msgs {
		var article *model.BlogArticle
		if err := json.Unmarshal(msg.Body, &article); err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			continue
		}
		article.State = int32(state.Approved)

		// 处理文章（例如保存到数据库）
		err := svcCtx.ArticleModel.Update(ctx, article.Id, article)
		if err != nil {
			log.Printf("failed to save article: %v", err)
			continue
		}
		log.Printf("article ID %d processed and saved", article.Id)
	}
}
