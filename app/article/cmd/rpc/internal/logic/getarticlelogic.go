package logic

import (
	"context"
	"fmt"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单篇文章的详细信息
func (l *GetArticleLogic) GetArticle(in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
	exist, err := l.svcCtx.ArticleModel.ExistArticleByID(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("check article existence failed, id: %d, error: %v", in.Id, err)
		return nil, err
	}
	if !exist {
		l.Logger.Errorf("article not found, id: %d", in.Id)
		return nil, fmt.Errorf("article not found, id: %d", in.Id)
	}

	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	data := &pb.Article{}
	err = copier.Copy(data, article)
	if err != nil {
		return nil, err
	}

	// 返回成功响应
	l.Logger.Infof("article retrieved successfully, id: %d", in.Id)

	return &pb.GetArticleResponse{
		Msg:  "获取文章成功",
		Data: data,
	}, nil
}
