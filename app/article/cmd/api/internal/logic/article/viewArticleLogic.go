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

type ViewArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 浏览文章
func NewViewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewArticleLogic {
	return &ViewArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewArticleLogic) ViewArticle(req *types.ViewArticleRequest) (resp *types.ViewArticleResponse, err error) {
	if req.Id <= 0 {
		l.Logger.Errorf("invalid param req.Id: %d ,err: %v", req.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR), "invalid article id: %d", req.Id)
	}
	userID := ctxdata.GetUidFromCtx(l.ctx)
	if userID <= 0 {
		l.Logger.Errorf("invalid param userID: %d,err: %v", userID, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR), "invalid user id: %d", userID)
	}

	viewArticleResp, err := l.svcCtx.ArticleServiceRpc.ViewArticle(l.ctx, &articleservice.ViewArticleRequest{
		ArticleId: req.Id,
		UserId:    userID,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.ViewArticleResponse{} // 初始化 resp
	err = copier.Copy(resp, viewArticleResp)
	if err != nil {
		l.Logger.Errorf("failed to copy submitArticleResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
