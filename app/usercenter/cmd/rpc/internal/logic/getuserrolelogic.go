package logic

import (
	"context"
	"fmt"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoleLogic {
	return &GetUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserRoleLogic) GetUserRole(in *userpb.GetUserRoleRequest) (*userpb.GetUserRoleResponse, error) {
	if in.Id < 0 {
		l.Logger.Errorf("failed to find user, userId: %d", in.Id)
		return nil, fmt.Errorf("invalid ID: %d", in.Id)
	}

	user, err := l.svcCtx.UserModel.FindByUserId(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to find user, userId: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(err, "查询用户失败, userId: %d", in.Id)
	}

	if user == nil {
		l.Logger.Errorf("user not found, userId: %s", in.Id)
		return nil, nil
	}

	return &userpb.GetUserRoleResponse{
		Role: user.Role,
	}, nil
}
