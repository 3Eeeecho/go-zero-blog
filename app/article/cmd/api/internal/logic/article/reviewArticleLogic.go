package article

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/3Eeeecho/go-zero-blog/pkg/ctxdata"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReviewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewArticleLogic {
	return &ReviewArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReviewArticleLogic) ReviewArticle(req *types.ReviewArticleRequest) (resp *types.ArticleCommonResponse, err error) {
	userID := ctxdata.GetUidFromCtx(l.ctx)
	reviewArticleresp, err := l.svcCtx.ArticleServiceRpc.ReviewArticle(l.ctx, &articleservice.ReviewArticleRequest{
		Id:         req.Id,
		Approved:   req.Approved,
		ReviewedBy: userID,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.ArticleCommonResponse{}
	err = copier.Copy(resp, reviewArticleresp)
	if err != nil {
		l.Logger.Errorf("failed to copy reviewArticleresp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}

	return resp, nil
}
