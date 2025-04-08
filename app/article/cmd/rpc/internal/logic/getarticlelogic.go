package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单篇文章的详细信息
func (l *GetArticleLogic) GetArticle(in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	cacheKey := fmt.Sprintf("article:detail:%d", in.Id)

	// 从 Redis 获取缓存
	cached, err := l.svcCtx.Redis.Get(cacheKey)
	if err == nil && cached != "" {
		var resp pb.GetArticleResponse
		if err := json.Unmarshal([]byte(cached), &resp); err != nil {
			l.Logger.Errorf("failed to unmarshal cached data, key: %s, error: %v", cacheKey, err)
			// 继续查询数据库，不直接返回错误
		} else {
			l.Logger.Infof("cache hit for article, key: %s", cacheKey)
			return &resp, nil
		}
	} else if err != redis.Nil {
		l.Logger.Errorf("failed to get from redis, key: %s, error: %v", cacheKey, err)
	}

	// 未命中缓存或 Redis 错误，查询数据库
	if err != redis.Nil { // redis.Nil 表示缓存不存在，不记录为错误
		l.Logger.Errorf("failed to get from redis, key: %s, error: %v", cacheKey, err)
	}

	exist, err := l.svcCtx.ArticleModel.ExistArticleByID(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("check article existence failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get articles failed")
	}
	if !exist {
		l.Logger.Errorf("article not found, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "get articles failed")
	}

	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("get article failed,error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed")
	}

	data := &pb.Article{}
	if err := copier.Copy(data, article); err != nil {
		l.Logger.Errorf("copy article to pb failed, error: %v", err)
		return nil, err
	}

	resp := &pb.GetArticleResponse{
		Msg:  "获取文章成功",
		Data: data,
	}

	// 存入 Redis，设置过期时间（例如 1 小时）
	jsonData, err := json.Marshal(resp) // 存储整个响应结构体
	if err != nil {
		l.Logger.Errorf("failed to marshal response, error: %v", err)
	} else {
		// 设置缓存，TTL 为 3600 秒（1小时）
		if err := l.svcCtx.Redis.Setex(cacheKey, string(jsonData), 3600); err != nil {
			l.Logger.Errorf("failed to set redis cache, key: %s, error: %v", cacheKey, err)
		} else {
			l.Logger.Infof("cache set successfully, key: %s", cacheKey)
		}
	}

	// 返回成功响应
	l.Logger.Infof("get article successfully, id: %d", in.Id)
	return resp, nil
}
