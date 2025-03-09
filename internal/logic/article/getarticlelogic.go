package logic

import (
	"context"
	"fmt"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取单篇文章的详细信息
func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleLogic) GetArticle(req *types.GetArticleRequest) (resp *types.Response, err error) {
	exist, err := l.svcCtx.ArticleModel.ExistArticleByID(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("check article existence failed, id: %d, error: %v", req.Id, err)
		return app.Response(e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil), err
	}
	if !exist {
		l.Logger.Errorf("article not found, id: %d", req.Id)
		return app.Response(e.ERROR_NOT_EXIST_ARTICLE, nil), fmt.Errorf("article not found, id: %d", req.Id)
	}

	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, req.Id)
	if err != nil {
		return app.Response(e.ERROR_GET_ARTICLE_FAIL, nil), err
	}

	// 返回成功响应
	l.Logger.Infof("article retrieved successfully, id: %d", req.Id)
	return app.Response(e.SUCCESS, article), nil
}
