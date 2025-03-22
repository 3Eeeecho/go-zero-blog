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
		return nil, errors.New("invalid article id")
	}

	getArticleResp, err := l.svcCtx.ArticleServiceRpc.GetArticle(l.ctx, &articleservice.GetArticleRequest{
		Id: int64(req.Id),
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GetArticleResponse{} // 初始化 resp
	err = copier.Copy(resp, getArticleResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
