package logic

import (
	"context"
	"fmt"
	"strconv"

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

	cacheKey := fmt.Sprintf("article:likes:%d", in.ArticleId)
	userIdStr := strconv.FormatInt(in.UserId, 10)

	// 优先检查 Redis
	exists, _ := l.svcCtx.Redis.Sismember(cacheKey, userIdStr)
	if !exists {
		// Redis 明确未点赞，检查数据库
		dbExists, dbErr := l.svcCtx.ArticleLikeModel.Exists(l.ctx, in.ArticleId, in.UserId)
		if dbErr != nil {
			l.Logger.Errorf("failed to check like in db, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, dbErr)
			return nil, xerr.NewErrCode(xerr.DB_ERROR)
		}
		if !dbExists {
			return nil, xerr.NewErrCodeMsg(xerr.REQUEST_PARAM_ERROR, "未点赞此文章")
		}
		// 数据库有记录，Redis 需要修复
		exists = true
	}

	// 先删数据库
	err = l.svcCtx.ArticleLikeModel.Delete(l.ctx, in.ArticleId, in.UserId)
	if err != nil {
		l.Logger.Errorf("failed to delete like, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 再删 Redis
	if exists {
		_, err = l.svcCtx.Redis.Srem(cacheKey, userIdStr)
		if err != nil {
			l.Logger.Errorf("failed to remove like in redis, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, err)
			// 不回滚数据库，记录日志即可
		}
		l.Logger.Infof("cache hit for unlike articleId: %d", in.ArticleId)
	}

	l.Logger.Infof("article unliked successfully, article_id: %d, user_id: %d", in.ArticleId, in.UserId)
	return &pb.ArticleCommonResponse{
		Msg: "取消点赞成功",
	}, nil
}
