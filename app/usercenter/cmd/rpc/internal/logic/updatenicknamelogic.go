package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNicknameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNicknameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNicknameLogic {
	return &UpdateNicknameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户名
func (l *UpdateNicknameLogic) UpdateNickname(in *userpb.UpdateNicknameRequest) (*userpb.UpdateNicknameResponse, error) {
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

	//TODO 检查用户名违规

	// 更新用户名为新用户名
	user.Nickname = in.NewNickname
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("failed to update nickname, nickname: %s, error: %v", in.NewNickname, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update nickname failed: %v", err)
	}

	// 记录用户名更新成功的日志
	l.Logger.Infof("nickname updated successfully, new nickname: %s", in.NewNickname)
	// 返回更新用户名响应
	return &userpb.UpdateNicknameResponse{
		Msg: "更改昵称成功",
	}, nil
}
