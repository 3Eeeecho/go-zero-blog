package logic

import (
	"context"
	"fmt"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
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
	// 检查权限
	hasPermission, err := l.svcCtx.ArticleModel.CheckPermission(l.ctx, in.Id, in.UserId)
	if err != nil {
		l.Logger.Errorf("failed to check permission for article %d by user %d: %v", in.Id, in.UserId, err)
		return nil, err
	}

	user, err := l.svcCtx.UserRpc.GetUserRole(l.ctx, &userpb.GetUserRoleRequest{
		Id: in.UserId,
	})
	if err != nil {
		l.Logger.Errorf("user %d is not admin: %v", in.UserId, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get user role failed: %v", err)
	}

	//既不是文章作者也不是管理员,不允许操作
	if !hasPermission && user.Role != "admin" {
		l.Logger.Errorf("user %d has no permission to delete article %d", in.UserId, in.Id)
		return nil, xerr.NewErrCode(xerr.ERROR_FORBIDDEN) // 100003: "权限不足"
	}

	result := l.svcCtx.ArticleModel.Delete(l.ctx, in.Id)
	if result.Error != nil {
		l.Logger.Errorf("failed to delete article, id: %d, error: %v", in.Id, result.Error)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete article failed")
	}

	if result.RowsAffected == 0 {
		l.Logger.Errorf("article not exist, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "articles not exist")
	}

	l.svcCtx.Redis.Del(fmt.Sprintf("article:detail:%d", in.Id))
	l.svcCtx.Redis.Del("article:list:tag_*:page_*") // 粗粒度失效

	l.Logger.Infof("article deleted successfully, id: %d", in.Id)
	return &pb.ArticleCommonResponse{
		Msg: "删除文章成功",
	}, nil
}
