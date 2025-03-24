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

type EditArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改文章
func NewEditArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditArticleLogic {
	return &EditArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditArticleLogic) EditArticle(req *types.EditArticleRequest) (resp *types.ArticleCommonResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	editArticleResp, err := l.svcCtx.ArticleServiceRpc.EditArticle(l.ctx, &articleservice.EditArticleRequest{
		Id:         int64(req.Id),
		TagId:      int64(req.TagId),
		Title:      req.Title,
		Desc:       req.Desc,
		Content:    req.Content,
		ModifiedBy: userId,
		UserId:     userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.ArticleCommonResponse{} // 初始化 resp
	err = copier.Copy(resp, editArticleResp)
	if err != nil {
		l.Logger.Errorf("failed to copy editArticleResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
