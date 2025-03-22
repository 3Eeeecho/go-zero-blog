package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
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

func (l *UpdateUserRoleLogic) UpdateUserRole(in *userpb.UpdateUserRoleRequest) (*userpb.UpdateUserRoleResponse, error) {
	// 检查调用者是否为管理员（从上下文获取ID）
	if in.AdminId == 0 {
		l.Logger.Errorf("admin ID not found in context")
		return nil, errors.New("unauthorized: admin ID not provided")
	}

	admin, err := l.svcCtx.UserModel.FindOne(l.ctx, in.AdminId) // 假设 ID 是字符串
	if err != nil {
		l.Logger.Errorf("failed to find admin: %s, error: %v", in.AdminId, err)
		return nil, errors.New("admin not found")
	}
	if admin.Role != RoleAdmin {
		l.Logger.Errorf("permission denied for user: %s, role: %s", in.AdminId, admin.Role)
		return nil, errors.New("permission denied: requires admin role")
	}

	// 验证目标用户
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, int64(in.Id))
	if err != nil {
		l.Logger.Errorf("failed to find user: %d, error: %v", in.Id, err)
		return nil, errors.New("user not found")
	}

	// 检查角色有效性
	switch in.Role {
	case RoleUser, RoleAuthor, RoleAdmin:
		// 合法角色
	default:
		l.Logger.Errorf("invalid role: %s", in.Role)
		return nil, errors.New("invalid role value")
	}

	// 更新用户角色
	user.Role = in.Role
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("failed to update user role: %v", err)
		return nil, errors.New("failed to update role")
	}

	l.Logger.Infof("user %d role updated to %s by admin %s", in.Id, in.Role, in.AdminId)
	return &userpb.UpdateUserRoleResponse{
		Msg: "用户权限更改成功",
	}, nil
}
