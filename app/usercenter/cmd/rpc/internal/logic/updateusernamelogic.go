package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUsernameLogic {
	return &UpdateUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUsername 更新用户名逻辑
func (l *UpdateUsernameLogic) UpdateUsername(in *userpb.UpdateUsernameRequest) (*userpb.UpdateUsernameResponse, error) {
	// 检查新用户名是否为空
	if in.NewUsername == "" {
		l.Logger.Errorf("new username is empty, in: %+v", in)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	// 根据用户 ID 查询用户信息
	user, err := l.svcCtx.UserModel.FindByUserId(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to find user, userId: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query user by id %d: %v", in.Id, err)
	}
	if user == nil {
		l.Logger.Errorf("user not found, userId: %d", in.Id)
		return nil, xerr.NewErrCode(xerr.USER_NOT_FOUND)
	}

	// 检查新用户名是否已被占用
	existingUser, err := l.svcCtx.UserModel.FindByUsername(l.ctx, in.NewUsername)
	if err != nil {
		l.Logger.Errorf("failed to check new username, error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to check username %s: %v", in.NewUsername, err)
	}
	if existingUser != nil {
		l.Logger.Errorf("new username already exists: %s", in.NewUsername)
		return nil, xerr.NewErrCode(xerr.USER_ALREADY_EXISTS)
	}

	// 更新用户名为新用户名
	user.Username = in.NewUsername
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("failed to update username, username: %s, error: %v", in.NewUsername, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update username failed: %v", err)
	}

	// 记录用户名更新成功的日志
	l.Logger.Infof("username updated successfully, new username: %s", in.NewUsername)
	// 返回更新用户名响应
	return &userpb.UpdateUsernameResponse{
		Msg: "成功修改用户名",
	}, nil
}
