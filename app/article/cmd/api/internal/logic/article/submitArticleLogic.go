package article

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitArticleLogic {
	return &SubmitArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitArticleLogic) SubmitArticle(req *types.SubmitArticleRequest) (resp *types.ArticleCommonResponse, err error) {
	submitArticleResp, err := l.svcCtx.ArticleServiceRpc.SubmitArticle(l.ctx, &articleservice.SubmitArticleRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("failed to submit article: %v", err)
		return nil, err
	}

	resp = &types.ArticleCommonResponse{}

	err = copier.Copy(resp, submitArticleResp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
