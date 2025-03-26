package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewArticleLogic {
	return &ViewArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewArticleLogic) ViewArticle(in *pb.ViewArticleRequest) (*pb.ArticleCommonResponse, error) {
	// 检查文章是否存在
	exist, err := l.svcCtx.ArticleModel.ExistArticleByID(l.ctx, in.ArticleId)
	if err != nil {
		l.Logger.Errorf("failed to find article, id: %d, error: %v", in.ArticleId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	if !exist {
		l.Logger.Errorf("article not found, id: %d", in.ArticleId)
		return nil, xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND)
	}

	// 增加浏览量
	err = l.svcCtx.ArticleModel.IncrementViews(l.ctx, in.ArticleId)
	if err != nil {
		l.Logger.Errorf("failed to increment views, article_id: %d, error: %v", in.ArticleId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	l.Logger.Infof("article views incremented, article_id: %d", in.ArticleId)
	return &pb.ArticleCommonResponse{
		Msg: "浏览量增加成功",
	}, nil
}
