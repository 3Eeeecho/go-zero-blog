package article

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/jinzhu/copier"

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
	editArticleResp, err := l.svcCtx.ArticleServiceRpc.EditArticle(l.ctx, &articleservice.EditArticleRequest{
		Id:         int64(req.Id),
		TagId:      int64(req.TagId),
		Title:      req.Title,
		Desc:       req.Desc,
		Content:    req.Content,
		ModifiedBy: req.ModifiedBy,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ArticleCommonResponse{} // 初始化 resp
	err = copier.Copy(resp, editArticleResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
