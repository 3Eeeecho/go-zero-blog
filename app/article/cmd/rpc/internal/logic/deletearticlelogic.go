package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/utils"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"

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
	// 1. 检查权限
	if err := utils.CheckArticlePermission(l.ctx, l.svcCtx, in.Id, in.UserId); err != nil {
		return nil, err
	}

	// 2. 获取文章信息（用于删除缓存）
	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("get article failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed")
	}

	// 3. 删除文章
	if err := l.svcCtx.ArticleModel.Delete(l.ctx, in.Id); err != nil {
		l.Logger.Errorf("delete article failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete article failed")
	}

	// 4. 清理缓存
	utils.CleanArticleCache(l.svcCtx, in.Id, article.TagId)

	// 返回成功响应
	l.Logger.Infof("article deleted successfully, id: %d", in.Id)
	return &pb.ArticleCommonResponse{
		Msg: "删除文章成功",
	}, nil
}
