package article

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/articleservice"
	"github.com/jinzhu/copier"

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
		return nil, errors.New("invalid tag id")
	}

	delArticleResp, err := l.svcCtx.ArticleServiceRpc.DeleteArticle(l.ctx, &articleservice.DeleteArticleRequest{
		Id: int64(req.Id),
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ArticleCommonResponse{} // 初始化 resp
	err = copier.Copy(resp, delArticleResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
