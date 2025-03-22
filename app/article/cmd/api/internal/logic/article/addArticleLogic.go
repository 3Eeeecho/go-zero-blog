package article

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/jinzhu/copier"

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
	addArticleResp, err := l.svcCtx.ArticleServiceRpc.AddArticle(l.ctx, &articleservice.AddArticleRequest{
		TagId:     req.TagId,
		Title:     req.Title,
		Desc:      req.Desc,
		Content:   req.Content,
		CreatedBy: req.CreatedBy,
	})

	if err != nil {
		return nil, err
	}

	resp = &types.ArticleCommonResponse{} // 初始化 resp
	err = copier.Copy(resp, addArticleResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
