package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditArticleLogic {
	return &EditArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章
func (l *EditArticleLogic) EditArticle(in *pb.EditArticleRequest) (*pb.ArticleCommonResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ArticleCommonResponse{}, nil
}
