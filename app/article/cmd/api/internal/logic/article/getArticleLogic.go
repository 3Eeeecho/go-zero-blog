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

func (l *GetArticleLogic) GetArticle(req *types.GetArticleRequest) (resp *types.GetArticleResponse, err error) {
	if req.Id <= 0 {
		l.Logger.Errorf("invalid param req.Id: %d ,err: %v", req.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR), "invalid article id: %d", req.Id)
	}

	getArticleResp, err := l.svcCtx.ArticleServiceRpc.GetArticle(l.ctx, &articleservice.GetArticleRequest{
		Id: int64(req.Id),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.GetArticleResponse{} // 初始化 resp
	err = copier.Copy(resp, getArticleResp)
	if err != nil {
		l.Logger.Errorf("failed to copy submitArticleResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
