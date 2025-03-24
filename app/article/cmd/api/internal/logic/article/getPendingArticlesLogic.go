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

type GetPendingArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取待审核文章列表
func NewGetPendingArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPendingArticlesLogic {
	return &GetPendingArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPendingArticlesLogic) GetPendingArticles(req *types.GetPendingArticlesRequest) (resp *types.GetPendingArticlesResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	getArticlesPendingResp, err := l.svcCtx.ArticleServiceRpc.GetPendingArticles(l.ctx, &articleservice.GetPendingArticlesRequest{
		UserId:   userId,
		PageNum:  int64(req.PageNum),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.GetPendingArticlesResponse{}
	err = copier.Copy(resp, getArticlesPendingResp)
	if err != nil {
		l.Logger.Errorf("failed to copy getArticlesPendingResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
