package article

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/jinzhu/copier"

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
	getArticlesResp, err := l.svcCtx.ArticleServiceRpc.GetArticles(l.ctx, &articleservice.GetArticlesRequest{
		TagId:    req.TagId,
		State:    req.State,
		PageNum:  int64(req.PageNum),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	if getArticlesResp == nil {
		l.Logger.Error("getArticlesResp is nil")
		return nil, errors.New("getArticlesResp is nil")
	}

	resp = &types.GetArticlesResponse{}

	err = copier.Copy(resp, getArticlesResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
