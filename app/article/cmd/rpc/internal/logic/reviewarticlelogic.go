package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReviewArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReviewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReviewArticleLogic {
	return &ReviewArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReviewArticleLogic) ReviewArticle(in *pb.ReviewArticleRequest) (*pb.ReviewArticleResponse, error) {
	user, err := l.svcCtx.UserRpc.GetUserRole(l.ctx, &userpb.GetUserRoleRequest{
		Id: in.ReviewedBy,
	})
	if err != nil {
		l.Logger.Errorf("user %d is not admin: %v", in.ReviewedBy, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get user role failed: %v", err)
	}
	if user.Role != "admin" {
		l.Logger.Errorf("user %d is not admin", in.ReviewedBy)
		return nil, xerr.NewErrCode(xerr.ERROR_FORBIDDEN)
	}

	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Logger.Errorf("get article failed, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed,id: %d", in.Id)
	}
	if err == gorm.ErrRecordNotFound {
		l.Logger.Errorf("article not failed, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "get article failed,id: %d", in.Id)
	}
	if article.State != StatePending {
		l.Logger.Errorf("article ID %d is not in pending state: current state %d", in.Id, article.State)
		return nil, xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "文章不在待审核状态")
	}

	updates := map[string]any{
		"state":       map[bool]int8{true: 2, false: 3}[in.Approved], // 1: 通过, 2: 拒绝
		"modified_by": in.ReviewedBy,
		"modified_on": time.Now().Unix(),
	}

	err = l.svcCtx.ArticleModel.Update(l.ctx, in.Id, updates)
	if err != nil {
		l.Logger.Errorf("update article failed, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update article failed,id: %d", in.Id)
	}

	l.Logger.Infof("article ID %d reviewed by %d, approved: %v", in.Id, in.ReviewedBy, in.Approved)
	return &pb.ReviewArticleResponse{Msg: "文章审核完成"}, nil
}
