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

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章
func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleRequest) (resp *types.ArticleCommonResponse, err error) {
	if req.Id <= 0 {
		l.Logger.Errorf("invalid param req.Id: %d ,err: %v", req.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR), "invalid article id: %d", req.Id)

	}
	userId := ctxdata.GetUidFromCtx(l.ctx)

	delArticleResp, err := l.svcCtx.ArticleServiceRpc.DeleteArticle(l.ctx, &articleservice.DeleteArticleRequest{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.ArticleCommonResponse{} // 初始化 resp
	err = copier.Copy(resp, delArticleResp)
	if err != nil {
		l.Logger.Errorf("failed to copy delArticleResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
