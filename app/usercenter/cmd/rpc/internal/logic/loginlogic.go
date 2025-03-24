package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/usercenter"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/util"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginLogic) Login(in *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	// 检查用户名和密码是否为空
	if in.Username == "" || in.Password == "" {
		l.Logger.Errorf("invalid username or password, req: %+v", in)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	// 查询用户信息
	user, err := l.svcCtx.UserModel.FindByUsername(l.ctx, in.Username)
	if err != nil {
		l.Logger.Errorf("failed to find user, username: %s, error: %v", in.Username, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to query user by username %s: %v", in.Username, err)
	}

	if user == nil {
		l.Logger.Errorf("user not found, username: %s", in.Username)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND), "user not found, username: %s", in.Username)
	}

	//TODO 解密和校验包装成函数
	// 采用AES256加密校验
	// 解密客户端传来的密码
	key := []byte(l.svcCtx.Config.Crypto.Key)
	plainPassword, err := util.DecryptPassword(key, in.Password)
	if err != nil {
		l.Logger.Errorf("failed to decrypt password, error: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "decrypt password failed: %v", err)
	}

	// 校验密码（bcrypt 比较）
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPassword))
	l.Logger.Info("plain password:", plainPassword)
	if err != nil {
		l.Logger.Errorf("password mismatch, username: %s", in.Username)
		return nil, xerr.NewErrCode(xerr.ERROR_INVALID_PASSWORD)
	}

	//生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		l.Logger.Errorf("generate token failed: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "generate token failed: %v", err)
	}

	l.Logger.Infof("user registered successfully, username: %s", in.Username)
	return &userpb.LoginResponse{
		Token:   tokenResp.AccessToken,
		Expires: tokenResp.AccessExpire,
	}, nil
}
