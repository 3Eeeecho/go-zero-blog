package logic

import (
	"context"
	"encoding/json"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/utils"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesLogic {
	return &GetArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文章列表
func (l *GetArticlesLogic) GetArticles(in *pb.GetArticlesRequest) (*pb.GetArticlesResponse, error) {
	// 尝试从 Redis 获取
	cacheKey := l.svcCtx.CacheKeys.GetListCacheKey(in.TagId, int(in.PageNum), int(in.PageSize))
	cached, err := l.svcCtx.Redis.Get(cacheKey)
	if err == nil && cached != "" {
		var resp pb.GetArticlesResponse
		json.Unmarshal([]byte(cached), &resp)
		l.Logger.Infof("cache hit for articles, key: %s", cacheKey)
		return &resp, nil
	}

	// 构造过滤条件
	maps := make(map[string]any)
	if in.TagId != 0 {
		maps["tag_id"] = in.TagId
	}
	maps["state"] = utils.StatePublished

	articles, err := l.svcCtx.ArticleModel.GetArticles(l.ctx, int(in.PageNum), int(in.PageSize), maps)
	if err != nil {
		l.Logger.Errorf("get articles failed, page_num: %d, page_size: %d, maps: %v, error: %v",
			in.PageNum, in.PageSize, maps, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get articles failed")
	}

	data := make([]*pb.Article, len(articles))
	for i, article := range articles {
		data[i] = &pb.Article{}
		if err := copier.Copy(data[i], article); err != nil {
			l.Logger.Errorf("copy article to pb failed, id: %d, error: %v", article.Id, err)
		}
	}

	resp := &pb.GetArticlesResponse{
		Data:     data,
		Total:    int64(len(articles)),
		PageNum:  in.PageNum,
		PageSize: in.PageSize,
	}

	go l.updateArticlesCache(cacheKey, resp)

	// 返回成功响应
	l.Logger.Infof("get articles successfully, page_num: %d, page_size: %d", in.PageNum, in.PageSize)
	return resp, nil
}

func (l *GetArticlesLogic) updateArticlesCache(cacheKey string, resp *pb.GetArticlesResponse) {
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
