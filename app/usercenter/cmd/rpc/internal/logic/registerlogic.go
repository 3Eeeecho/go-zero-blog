package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/usercenter"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/util"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
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

// Register 用户注册逻辑
func (l *RegisterLogic) Register(in *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	// 检查用户名和密码是否为空
	if in.Username == "" || in.Password == "" || in.Nickname == "" {
		l.Logger.Errorf("invalid params, req: %+v", in)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	// 检查用户是否已存在（用户名唯一）
	user, err := l.svcCtx.UserModel.FindByUsername(l.ctx, in.Username)
	if err != nil {
		l.Logger.Errorf("failed to check user existence, username: %s, error: %v", in.Username, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query user by username %s: %v", in.Username, err)
	}
	if user != nil {
		l.Logger.Errorf("user already exists, username: %s", in.Username)
		return nil, xerr.NewErrCode(xerr.USER_ALREADY_EXISTS)
	}

	// 获取解密密钥
	key := []byte(l.svcCtx.Config.Crypto.Key)
	// 解密客户端传来的密码
	plainPassword, err := util.DecryptPassword(key, in.Password)
	if err != nil {
		l.Logger.Errorf("failed to decrypt password, error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "decrypt password failed: %v", err)
	}

	// 使用 bcrypt 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("failed to hash password, error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "hash password failed: %v", err)
	}

	// 创建新用户对象
	newUser := &model.BlogUser{
		Username: in.Username,
		Password: string(hashedPassword),
		Role:     "user",
		Nickname: in.Nickname,
	}
	// 将新用户插入数据库
	err = l.svcCtx.UserModel.Insert(l.ctx, newUser)
	if err != nil {
		l.Logger.Errorf("failed to insert user, username: %s, error: %v", in.Username, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert user failed: %v", err)
	}

	// 创建 GenerateTokenLogic 实例以生成 Token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	// 调用 GenerateToken 方法生成 Token
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: newUser.Id,
	})
	if err != nil {
		l.Logger.Errorf("generate token failed: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "generate token failed: %v", err)
	}

	// 记录注册成功的日志
	l.Logger.Infof("user registered successfully, username: %s", in.Username)
	// 返回注册响应
	return &userpb.RegisterResponse{
		Token:   tokenResp.AccessToken,
		Expires: tokenResp.AccessExpire,
	}, nil
}
