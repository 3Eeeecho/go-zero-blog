package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章列表
func NewGetArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesLogic {
	return &GetArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesLogic) GetArticles(req *types.GetArticlesRequest) (resp *types.GetArticlesResponse, err error) {
	// 设置分页默认值
	pageNum := req.PageNum
	pageSize := req.PageSize
	if pageNum <= 0 {
		pageNum = 1 // 默认第1页
	}
	if pageSize <= 0 {
		pageSize = 10 // 默认每页10条
	}

	// 构造过滤条件
	maps := make(map[string]any)
	if req.TagId != 0 {
		maps["tag_id"] = req.TagId
	}
	if req.State != 0 {
		maps["state"] = req.State
	}

	articles, err := l.svcCtx.ArticleModel.GetArticles(l.ctx, pageNum, pageSize, maps)
	if err != nil {
		l.Logger.Errorf("get articles failed, page_num: %d, page_size: %d, maps: %v, error: %v",
			pageNum, pageSize, maps, err)
		return app.GetArticlesResponse(e.ERROR_GET_ARTICLES_FAIL, nil, 0, pageNum, pageSize), err
	}

	//文章总数
	total, err := l.svcCtx.ArticleModel.CountByCondition(l.ctx, maps)
	if err != nil {
		l.Logger.Errorf("count articles failed,condition:%v,error:%v", maps, err)
		return app.GetArticlesResponse(e.ERROR_COUNT_ARTICLE_FAIL, nil, 0, pageNum, pageSize), nil
	}

	// 返回成功响应
	l.Logger.Infof("articles retrieved successfully, page_num: %d, page_size: %d", pageNum, pageSize)
	return app.GetArticlesResponse(e.SUCCESS, articles, total, pageNum, pageSize), nil
}
