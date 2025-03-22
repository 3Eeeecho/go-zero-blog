package logic

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/pb"

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

// 修改用户名
func (l *UpdateUsernameLogic) UpdateUsername(in *pb.UpdateUsernameRequest) (*pb.UpdateUsernameResponse, error) {
	if in.NewUsername == "" {
		l.Logger.Errorf("new username is empty, in: %+v", in)
		return nil, errors.New("用户名为空")
	}

	// 获取当前用户信息
	user, err := l.svcCtx.UserModel.FindByUserId(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to find user, userId: %s, error: %v", in.Id, err)
		return nil, err
	}
	if user == nil {
		l.Logger.Errorf("user not found, userId: %s", in.Id)
		return nil, errors.New("不存在当前用户")
	}

	// 检查新用户名是否已被占用
	existingUser, err := l.svcCtx.UserModel.FindByUsername(l.ctx, in.NewUsername)
	if err != nil {
		l.Logger.Errorf("failed to check new username, error: %v", err)
		return nil, err
	}
	if existingUser != nil {
		l.Logger.Errorf("new username already exists: %s", in.NewUsername)
		return nil, errors.New("用户名已被占用")
	}

	// 更新用户名
	user.Username = in.NewUsername
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("failed to update username, username: %s, error: %v", in.NewUsername, err)
		return nil, err
	}

	l.Logger.Info("username updated successfully")

	return &pb.UpdateUsernameResponse{
		Msg: "成功修改用户名",
	}, nil
}
