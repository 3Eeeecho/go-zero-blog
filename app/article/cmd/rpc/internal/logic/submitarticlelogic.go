package logic

import (
	"context"
	"encoding/json"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/rabbitmq/amqp091-go"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitArticleLogic {
	return &SubmitArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitArticleLogic) SubmitArticle(in *pb.SubmitArticleRequest) (*pb.SubmitArticleResponse, error) {
	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("get article failed, id: %d", in.Id)
		return nil, err
	}

	body, err := json.Marshal(article)
	if err != nil {
		l.Logger.Errorf("failed to marshal article: %v", err)
		return nil, err
	}

	err = l.svcCtx.MQChannel.PublishWithContext(
		l.ctx,
		"",
		l.svcCtx.Config.RabbitMQ.SubmissionQueue,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		l.Logger.Errorf("failed to publish to submission queue: %v", err)
		return nil, err
	}

	l.Logger.Infof("article submission queued for id %d", in.Id)
	return &pb.SubmitArticleResponse{Msg: "文章提交已进入队列"}, nil
}
