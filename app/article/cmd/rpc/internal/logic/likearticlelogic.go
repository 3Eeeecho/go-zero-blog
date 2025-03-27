package logic

import (
	"context"
	"fmt"
	"strconv"

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

	cacheKey := fmt.Sprintf("article:likes:%d", in.ArticleId)
	userIDStr := strconv.FormatInt(in.UserId, 10)

	// 先检查 Redis
	exists, err := l.svcCtx.Redis.Sismember(cacheKey, userIDStr)
	if err != nil {
		l.Logger.Errorf("failed to check like in redis, error: %v", err)
	} else if exists {
		l.Logger.Infof("cache hit for like articleId: %d", in.ArticleId)
		return nil, xerr.NewErrCodeMsg(xerr.REQUEST_PARAM_ERROR, "已点赞过此文章")
	}

	// Redis 未命中或失效时检查数据库
	if err != nil || !exists {
		dbExists, dbErr := l.svcCtx.ArticleLikeModel.Exists(l.ctx, in.ArticleId, in.UserId)
		if dbErr != nil {
			l.Logger.Errorf("failed to check like in db, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, dbErr)
			return nil, xerr.NewErrCode(xerr.DB_ERROR)
		}
		if dbExists {
			l.svcCtx.Redis.Sadd(cacheKey, userIDStr) // 修复缓存
			return nil, xerr.NewErrCodeMsg(xerr.REQUEST_PARAM_ERROR, "已点赞过此文章")
		}
	}

	// 先写 Redis
	_, err = l.svcCtx.Redis.Sadd(cacheKey, userIDStr)
	if err != nil {
		l.Logger.Errorf("failed to add like in redis, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// TODO 定时任务同步数据库
	like := &model.BlogArticleLike{ArticleId: in.ArticleId, UserId: in.UserId}
	err = l.svcCtx.ArticleLikeModel.Insert(l.ctx, like)
	if err != nil {
		l.svcCtx.Redis.Srem(cacheKey, userIDStr) // 回滚 Redis
		l.Logger.Errorf("failed to insert like, article_id: %d, user_id: %d, error: %v", in.ArticleId, in.UserId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	l.Logger.Infof("article liked successfully, article_id: %d, user_id: %d", in.ArticleId, in.UserId)
	return &pb.ArticleCommonResponse{Msg: "点赞成功"}, nil
}
