package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"

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

// 查看文章详情
func (l *ViewArticleLogic) ViewArticle(in *pb.ViewArticleRequest) (*pb.ViewArticleResponse, error) {
	cacheKey := l.svcCtx.CacheKeys.GetDetailCacheKey(in.ArticleId)
	var resp *pb.ViewArticleResponse

	// 从 Redis 获取缓存
	cached, err := l.svcCtx.Redis.GetCtx(l.ctx, cacheKey)
	if err == nil && cached != "" {
		var detailCache pb.ViewArticleResponse
		if err = json.Unmarshal([]byte(cached), &detailCache); err != nil {
			l.Logger.Errorf("failed to unmarshal cached data, key: %s, error: %v", cacheKey, err)
			// 继续查询数据库，不直接返回错误
		} else {
			l.Logger.Infof("cache hit for article, key: %s", cacheKey)
			resp = &detailCache
		}
	}

	if resp == nil {
		// 缓存未命中，查询数据库
		article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.ArticleId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				l.Logger.Errorf("article not found, id: %d", in.ArticleId)
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "get article failed")
			}
			l.Logger.Errorf("get article failed, id: %d, error: %v", in.ArticleId, err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed")
		}

		data := &pb.Article{}
		if err := copier.Copy(data, article); err != nil {
			l.Logger.Errorf("copy article to pb failed, id: %d, error: %v", in.ArticleId, err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "copy article failed")
		}

		resp = &pb.ViewArticleResponse{
			Data: data,
		}

		// 异步更新缓存
		go l.updateDetailCache(cacheKey, resp)
	}

	// 文章浏览量加一
	go l.incrementViewCount(in.ArticleId, in.UserId)

	l.Logger.Infof("article views incremented, article_id: %d", in.ArticleId)
	return resp, nil
}

func (l *ViewArticleLogic) updateDetailCache(cacheKey string, resp *pb.ViewArticleResponse) {
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

func (l *ViewArticleLogic) incrementViewCount(articleId int64, userId int64) {
	// 用户访问记录的key
	userViewKey := fmt.Sprintf("article:user_view:%d", articleId)

	// 检查用户是否已经访问过该文章
	if userId > 0 {
		userIdStr := fmt.Sprintf("%d", userId)
		exists, err := l.svcCtx.Redis.Sismember(userViewKey, userIdStr)
		if err == nil && exists {
			// 用户已经访问过，不增加计数
			l.Logger.Infof("user %d already viewed article %d, skip counting", userId, articleId)
			return
		}

		// 记录用户访问
		if userId > 0 {
			_, err = l.svcCtx.Redis.Sadd(userViewKey, userIdStr)
			if err != nil {
				l.Logger.Errorf("failed to record user view in Redis, article_id: %d, user_id: %d, error: %v",
					articleId, userId, err)
			}

			// 设置过期时间，例如一天后过期，允许用户再次增加浏览量
			l.svcCtx.Redis.Expire(userViewKey, 86400) // 24小时
		}
	}

	// 增加Redis中的浏览量计数
	viewKey := "article:view_count"
	timeKey := "article:view_time"
	currentTime := time.Now().Unix()

	if _, err := l.svcCtx.Redis.Hincrby(viewKey, fmt.Sprintf("%d", articleId), 1); err != nil {
		l.Logger.Errorf("failed to increment view count in Redis, article_id: %d, error: %v", articleId, err)

		// Redis失败时直接写入数据库
		if err := l.svcCtx.ArticleModel.IncrementViews(articleId, 1); err != nil {
			l.Logger.Errorf("failed to increment view count in DB, article_id: %d, error: %v", articleId, err)
		}
		return
	}

	// 获取当前计数值
	countStr, err := l.svcCtx.Redis.Hget(viewKey, fmt.Sprintf("%d", articleId))
	if err != nil {
		l.Logger.Errorf("failed to get view count from Redis, article_id: %d, error: %v", articleId, err)
		return
	}
	count, _ := strconv.ParseInt(countStr, 10, 64)

	// 获取最后更新时间
	lastUpdateTimeStr, err := l.svcCtx.Redis.Hget(timeKey, fmt.Sprintf("%d", articleId))
	var lastUpdateTime int64 = 0
	if err == nil && lastUpdateTimeStr != "" {
		lastUpdateTime, _ = strconv.ParseInt(lastUpdateTimeStr, 10, 64)
	}

	if lastUpdateTime == 0 {
		// 首次更新时设置最后更新时间
		if err := l.svcCtx.Redis.Hset(timeKey, fmt.Sprintf("%d", articleId), fmt.Sprintf("%d", currentTime)); err != nil {
			l.Logger.Errorf("failed to set last update time in Redis, article_id: %d, error: %v", articleId, err)
		}
	}

	// 阈值为1分钟 (60秒)
	timeThreshold := int64(60)

	// 当计数达到一定阈值（此处为10）或者距离上次更新时间超过1分钟，更新数据库并重置Redis计数
	if count >= 10 || (lastUpdateTime > 0 && currentTime-lastUpdateTime > timeThreshold) {
		if err := l.svcCtx.ArticleModel.IncrementViews(articleId, count); err != nil {
			l.Logger.Errorf("failed to increment view count in DB, article_id: %d, count: %d, error: %v",
				articleId, count, err)
			return
		}

		// 重置Redis计数
		if err := l.svcCtx.Redis.Hset(viewKey, fmt.Sprintf("%d", articleId), "0"); err != nil {
			l.Logger.Errorf("failed to reset view count in Redis, article_id: %d, error: %v", articleId, err)
		}

		// 更新最后更新时间
		if err := l.svcCtx.Redis.Hset(timeKey, fmt.Sprintf("%d", articleId), fmt.Sprintf("%d", currentTime)); err != nil {
			l.Logger.Errorf("failed to update last update time in Redis, article_id: %d, error: %v", articleId, err)
		}

		l.Logger.Infof("Synced view count to DB and reset Redis, article_id: %d, count: %d", articleId, count)
	}
}
