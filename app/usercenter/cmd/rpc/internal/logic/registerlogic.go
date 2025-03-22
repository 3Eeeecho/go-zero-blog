package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/usercenter"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/util"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	// 检查用户名和密码是否为空
	if in.Username == "" || in.Password == "" {
		return nil, errors.New("账户或密码格式错误")
	}

	// 检查用户是否已存在(用户名唯一)
	user, err := l.svcCtx.UserModel.FindByUsername(l.ctx, in.Username)
	if err != nil {
		l.Logger.Errorf("failed to check user existence, username: %s, error: %v", in.Username, err)
		return nil, errors.Wrapf(err, "查询用户失败, username: %s", in.Username)
	}
	if user != nil {
		l.Logger.Errorf("user already exists, username: %s", in.Username)
		return nil, errors.Wrapf(err, "用户已存在, username: %s", in.Username)
	}

	// 解密客户端传来的密码
	key := []byte(l.svcCtx.Config.Crypto.Key)
	plainPassword, err := util.DecryptPassword(key, in.Password)
	if err != nil {
		l.Logger.Errorf("failed to decrypt password, error: %v", err)
		return nil, err
	}

	// 生成 bcrypt 哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("failed to hash password, error: %v", err)
		return nil, err
	}

	newUser := &model.BlogUser{
		Username: in.Username,
		Password: string(hashedPassword),
		Role:     in.Role,
	}
	err = l.svcCtx.UserModel.Insert(l.ctx, newUser)

	if err != nil {
		l.Logger.Errorf("failed to insert user, username: %s, error: %v", in.Username, err)
		return nil, err
	}

	//生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: newUser.Id,
	})
	if err != nil {
		l.Logger.Errorf("generate token failed: %v", err)
	}

	l.Logger.Infof("user registered successfully, username: %s", in.Username)
	return &userpb.RegisterResponse{
		Token:   tokenResp.AccessToken,
		Expires: tokenResp.AccessExpire,
	}, nil
}
