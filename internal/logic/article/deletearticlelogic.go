package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

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

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleRequest) (resp *types.Response, err error) {

	id := req.Id
	if id <= 0 {
		l.Logger.Errorf("invalid article id: %s, error: %v", id, err)
		return app.Response(e.INVALID_PARAMS, nil, nil)
	}

	err = l.svcCtx.ArticleModel.Delete(l.ctx, id)
	if err != nil {
		l.Logger.Errorf("failed to delete article, id: %d, error: %v", id, err)
		return app.Response(e.ERROR_DELETE_ARTICLE_FAIL, nil, err)
	}

	l.Logger.Infof("article deleted successfully, id: %d", id)
	return app.Response(e.SUCCESS, nil, nil)
}
