package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeArtilceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnlikeArtilceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeArtilceLogic {
	return &UnlikeArtilceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnlikeArtilceLogic) UnlikeArtilce(in *pb.UnlikeArticleRequest) (*pb.ArticleCommonResponse, error) {
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
	if !exists {
		l.Logger.Errorf("user has not liked this article, article_id: %d, user_id: %d", in.ArticleId, in.UserId)
		return nil, xerr.NewErrCodeMsg(xerr.REQUEST_PARAM_ERROR, "未点赞此文章")
	}

	// 删除点赞记录
	err = l.svcCtx.ArticleLikeModel.Delete(l.ctx, in.ArticleId, in.UserId)
	if err != nil {
		l.Logger.Errorf("failed to delete like, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	l.Logger.Infof("article unliked successfully, article_id: %d, user_id: %d", in.ArticleId, in.UserId)
	return &pb.ArticleCommonResponse{
		Msg: "取消点赞成功",
	}, nil
}
