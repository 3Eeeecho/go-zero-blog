package logic

import (
	"context"
	"errors"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"

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
	if err != nil || user.Role != "admin" {
		l.Logger.Errorf("user %d is not admin: %v", in.ReviewedBy, err)
		return nil, errors.New("权限不足")
	}

	article, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("article ID %d not found: %v", in.Id, err)
		return nil, err
	}

	if article.State != StatePending {
		return nil, errors.New("文章不在待审核状态")
	}

	updates := map[string]any{
		"state":       map[bool]int8{true: 2, false: 3}[in.Approved], // 1: 通过, 2: 拒绝
		"modified_by": in.ReviewedBy,
		"modified_on": time.Now().Unix(),
	}

	err = l.svcCtx.ArticleModel.Update(l.ctx, in.Id, updates)
	if err != nil {
		l.Logger.Errorf("failed to review article ID %d: %v", in.Id, err)
		return nil, err
	}

	l.Logger.Infof("article ID %d reviewed by %d, approved: %v", in.Id, in.ReviewedBy, in.Approved)
	return &pb.ReviewArticleResponse{Msg: "文章审核完成"}, nil
}
