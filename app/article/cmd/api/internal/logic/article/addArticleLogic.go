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

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增一篇文章
func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleRequest) (resp *types.ArticleCommonResponse, err error) {
	created_by := ctxdata.GetUidFromCtx(l.ctx)
	addArticleResp, err := l.svcCtx.ArticleServiceRpc.AddArticle(l.ctx, &articleservice.AddArticleRequest{
		TagId:     req.TagId,
		Title:     req.Title,
		Desc:      req.Desc,
		Content:   req.Content,
		CreatedBy: created_by,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.ArticleCommonResponse{} // 初始化 resp
	err = copier.Copy(resp, addArticleResp)
	if err != nil {
		l.Logger.Errorf("failed to copy addArticleResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
