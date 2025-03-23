package logic

import (
	"context"
	"encoding/json"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"

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
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Logger.Errorf("get article failed, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed,id: %d", in.Id)
	}
	if err == gorm.ErrRecordNotFound {
		l.Logger.Errorf("article not failed, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "get article failed,id: %d", in.Id)
	}

	//文章变为待审核状态
	data := map[string]interface{}{
		"state": StatePending,
	}
	err = l.svcCtx.ArticleModel.Update(l.ctx, in.Id, data)
	if err != nil {
		l.Logger.Errorf("failed to update article state: %v", err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	body, err := json.Marshal(article)
	if err != nil {
		l.Logger.Errorf("failed to marshal article: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
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
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}

	l.Logger.Infof("article submission queued for id %d", in.Id)
	return &pb.SubmitArticleResponse{Msg: "文章提交已进入队列"}, nil
}
