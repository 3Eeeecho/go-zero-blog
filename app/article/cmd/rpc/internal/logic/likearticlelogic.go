package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeArticleLogic {
	return &LikeArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeArticleLogic) LikeArticle(in *pb.LikeArticleRequest) (*pb.ArticleCommonResponse, error) {
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

	// 检查是否已点赞
	exists, err := l.svcCtx.ArticleLikeModel.Exists(l.ctx, in.ArticleId, in.UserId)
	if err != nil {
		l.Logger.Errorf("failed to check like status, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	if exists {
		l.Logger.Errorf("user has not liked this article, article_id: %d, user_id: %d", in.ArticleId, in.UserId)
		return nil, xerr.NewErrCodeMsg(xerr.REQUEST_PARAM_ERROR, "已点赞过此文章")
	}

	// 添加点赞记录
	like := &model.BlogArticleLike{
		ArticleId: in.ArticleId,
		UserId:    in.UserId,
		CreatedOn: time.Now().Unix(),
	}
	err = l.svcCtx.ArticleLikeModel.Insert(l.ctx, like)
	if err != nil {
		l.Logger.Errorf("failed to insert like, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	l.Logger.Infof("article liked successfully, article_id: %d, user_id: %d", in.ArticleId, in.UserId)
	return &pb.ArticleCommonResponse{
		Msg: "点赞成功",
	}, nil
}
