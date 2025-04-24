package logic

import (
	"context"
	"encoding/json"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"

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
// TODO 只需要返回文章详情即可,并不是给用户请求的
func (l *GetArticleLogic) GetArticle(in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	cacheKey := l.svcCtx.CacheKeys.GetDetailCacheKey(in.Id)

	// 从 Redis 获取缓存
	cached, err := l.svcCtx.Redis.GetCtx(l.ctx, cacheKey)
	if err == nil && cached != "" {
		var detailCache pb.GetArticleResponse
		if err = json.Unmarshal([]byte(cached), &detailCache); err != nil {
			l.Logger.Errorf("failed to unmarshal cached data, key: %s, error: %v", cacheKey, err)
			// 继续查询数据库，不直接返回错误
		} else {
			l.Logger.Infof("cache hit for article, key: %s", cacheKey)
			return &detailCache, nil
		}
	}

	// 缓存未命中，查询数据库
	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			l.Logger.Errorf("article not found, id: %d", in.Id)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "get article failed")
		}
		l.Logger.Errorf("get article failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed")
	}

	data := &pb.Article{}
	if err := copier.Copy(data, article); err != nil {
		l.Logger.Errorf("copy article to pb failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "copy article failed")
	}

	resp := &pb.GetArticleResponse{
		Data: data,
	}

	// 异步更新缓存
	go l.updateDetailCache(cacheKey, resp)

	// 返回成功响应
	l.Logger.Infof("get article successfully, id: %d", in.Id)
	return resp, nil
}

func (l *GetArticleLogic) updateDetailCache(cacheKey string, resp *pb.GetArticleResponse) {
	// 存入 Redis，设置过期时间
	data, err := json.Marshal(resp)
	if err != nil {
		l.Logger.Errorf("failed to marshal response, key: %s, error: %v", cacheKey, err)
		return
	}

	// 设置缓存并设置过期时间
	if err := l.svcCtx.Redis.Setex(cacheKey, string(data), int(l.svcCtx.CacheTTL.GetDetailCacheTTL().Seconds())); err != nil {
		l.Logger.Errorf("failed to set cache, key: %s, error: %v", cacheKey, err)
	}
}
