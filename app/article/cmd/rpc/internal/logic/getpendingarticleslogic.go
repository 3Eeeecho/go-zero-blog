package logic

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPendingArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPendingArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPendingArticlesLogic {
	return &GetPendingArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPendingArticlesLogic) GetPendingArticles(in *pb.GetPendingArticlesRequest) (*pb.GetPendingArticlesResponse, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	user, err := l.svcCtx.UserRpc.GetUserRole(l.ctx)
	if err != nil || user.Role != "admin" {
		l.Logger.Errorf("user %d is not admin: %v", userID, err)
		return nil, errors.New("权限不足")
	}

	return &pb.GetPendingArticlesResponse{}, nil
}
