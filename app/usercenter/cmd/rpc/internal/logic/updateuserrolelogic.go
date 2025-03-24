package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	RoleUser   = "user"
	RoleAuthor = "author"
	RoleAdmin  = "admin"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserRole 更新用户角色逻辑
func (l *UpdateUserRoleLogic) UpdateUserRole(in *userpb.UpdateUserRoleRequest) (*userpb.UpdateUserRoleResponse, error) {
	// 检查调用者 AdminId 是否有效
	if in.AdminId == 0 {
		l.Logger.Errorf("admin ID not found in context")
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	// 查询管理员信息
	admin, err := l.svcCtx.UserModel.FindOne(l.ctx, in.AdminId)
	if err != nil {
		l.Logger.Errorf("failed to find admin: %d, error: %v", in.AdminId, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query admin by id %d: %v", in.AdminId, err)
	}
	if admin == nil || admin.Role != RoleAdmin {
		l.Logger.Errorf("permission denied for user: %d, role: %s", in.AdminId, admin.Role)
		return nil, xerr.NewErrCode(xerr.ERROR_FORBIDDEN)
	}

	// 查询目标用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(in.Id))
	if err != nil {
		l.Logger.Errorf("failed to find user: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query user by id %d: %v", in.Id, err)
	}
	if user == nil {
		l.Logger.Errorf("user not found: %d", in.Id)
		return nil, xerr.NewErrCode(xerr.USER_NOT_FOUND)
	}

	// 检查新角色的有效性
	switch in.Role {
	case RoleUser, RoleAuthor, RoleAdmin:
		// 合法角色
	default:
		l.Logger.Errorf("invalid role: %s", in.Role)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	// 更新用户角色
	user.Role = in.Role
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("failed to update user role: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update user role failed: %v", err)
	}

	// 记录角色更新成功的日志
	l.Logger.Infof("user %d role updated to %s by admin %d", in.Id, in.Role, in.AdminId)
	// 返回更新角色响应
	return &userpb.UpdateUserRoleResponse{
		Msg: "用户权限更改成功",
	}, nil
}
