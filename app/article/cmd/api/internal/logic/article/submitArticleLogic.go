package article

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.ArticleCommonResponse{}

	err = copier.Copy(resp, submitArticleResp)
	if err != nil {
		l.Logger.Errorf("failed to copy submitArticleResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}

	return resp, nil
}
