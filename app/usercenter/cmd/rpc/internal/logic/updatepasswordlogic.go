package logic

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/util"
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

// 修改密码
func (l *UpdatePasswordLogic) UpdatePassword(in *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {

	// 获取当前用户信息
	user, err := l.svcCtx.UserModel.FindByUserId(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to find user, userId: %s, error: %v", in.Id, err)
		return nil, err
	}
	if user == nil {
		l.Logger.Errorf("user not found, userId: %s", in.Id)
		return nil, errors.New("user not found")
	}

	// 解密新密码
	key := []byte(l.svcCtx.Config.Crypto.Key)
	plainPassword, err := util.DecryptPassword(key, in.NewPassword)
	if err != nil {
		l.Logger.Errorf("failed to decrypt new password, error: %v", err)
		return nil, err
	}

	// 生成新的 bcrypt 哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("failed to hash new password, error: %v", err)
		return nil, err
	}

	// 更新密码
	user.Password = string(hashedPassword)
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		l.Logger.Errorf("failed to update password, username: %s, error: %v", user.Username, err)
		return nil, err
	}

	l.Logger.Infof("password updated successfully, username: %s", user.Username)
	return &pb.UpdatePasswordResponse{
		Msg: "成功更新密码",
	}, nil
}
