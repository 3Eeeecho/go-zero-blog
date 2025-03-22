package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除文章
func (l *DeleteArticleLogic) DeleteArticle(in *pb.DeleteArticleRequest) (*pb.ArticleCommonResponse, error) {
	err := l.svcCtx.ArticleModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to delete article, id: %d, error: %v", in.Id, err)
		return nil, err
	}

	l.Logger.Infof("article deleted successfully, id: %d", in.Id)
	return &pb.ArticleCommonResponse{
		Msg: "删除文章成功",
	}, nil
}
