package article

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/3Eeeecho/go-zero-blog/pkg/state"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesPendingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章列表
func NewGetArticlesPendingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesPendingLogic {
	return &GetArticlesPendingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesPendingLogic) GetArticlesPending(req *types.GetArticlesRequest) (resp *types.GetArticlesResponse, err error) {
	getArticlesPendingResp, err := l.svcCtx.ArticleServiceRpc.GetArticles(l.ctx, &articleservice.GetArticlesRequest{
		State:    int32(state.Pending),
		PageNum:  int64(req.PageNum),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	if getArticlesPendingResp == nil {
		l.Logger.Error("getArticlesResp is nil")
		return nil, errors.New("getArticlesResp is nil")
	}

	resp = &types.GetArticlesResponse{}

	err = copier.Copy(resp, getArticlesPendingResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
