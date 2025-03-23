package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get articles failed")
	}
	if !exist {
		l.Logger.Errorf("article not found, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "get articles failed")
	}

	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("get article failed,error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed")
	}

	data := &pb.Article{}
	err = copier.Copy(data, article)
	if err != nil {
		return nil, err
	}

	// 返回成功响应
	l.Logger.Infof("get article successfully, id: %d", in.Id)

	return &pb.GetArticleResponse{
		Msg:  "获取文章成功",
		Data: data,
	}, nil
}
