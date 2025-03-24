package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/util"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePassword 更新用户密码逻辑
func (l *UpdatePasswordLogic) UpdatePassword(in *userpb.UpdatePasswordRequest) (*userpb.UpdatePasswordResponse, error) {
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

	// 获取解密密钥
	key := []byte(l.svcCtx.Config.Crypto.Key)
	// 解密新密码
	plainPassword, err := util.DecryptPassword(key, in.NewPassword)
	if err != nil {
		l.Logger.Errorf("failed to decrypt new password, error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "decrypt new password failed: %v", err)
	}

	// 使用 bcrypt 生成新密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("failed to hash new password, error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "hash new password failed: %v", err)
	}

	// 更新用户密码
	user.Password = string(hashedPassword)
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("failed to update password, username: %s, error: %v", user.Username, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update password failed: %v", err)
	}

	// 记录密码更新成功的日志
	l.Logger.Infof("password updated successfully, username: %s", user.Username)
	// 返回更新密码响应
	return &userpb.UpdatePasswordResponse{
		Msg: "成功更新密码",
	}, nil
}
