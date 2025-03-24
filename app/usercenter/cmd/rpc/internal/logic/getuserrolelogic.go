package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
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
	// 检查参数
	if in.Id < 0 {
		l.Logger.Errorf("invalid user ID: %d", in.Id)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR) // 100002: "请求参数错误"
	}

	// 查询用户
	user, err := l.svcCtx.UserModel.FindByUserId(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to find user, userId: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query user by id %d: %v", in.Id, err) // 100005
	}

	// 检查用户是否存在
	if user == nil {
		l.Logger.Errorf("user not found, userId: %d", in.Id)
		return nil, xerr.NewErrCode(xerr.USER_NOT_FOUND) // 101001: "该用户不存在"
	}

	// 返回用户角色
	l.Logger.Infof("get user role successfully, userId: %d, role: %s", in.Id, user.Role)
	return &userpb.GetUserRoleResponse{
		Role: user.Role,
	}, nil
}
