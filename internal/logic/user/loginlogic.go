package user

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"
	"github.com/3Eeeecho/go-zero-blog/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取授权 Token
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 检查用户名和密码是否为空
	if req.Username == "" || req.Password == "" {
		l.Logger.Errorf("invalid username or password, req: %+v", req)
		return app.LoginResponse(e.INVALID_PARAMS, ""), nil
	}

	// 查询用户信息
	user, err := l.svcCtx.UserModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		l.Logger.Errorf("failed to find user, username: %s, error: %v", req.Username, err)
		return app.LoginResponse(e.ERROR_USER_CHECK_TOKEN_FAIL, ""), err
	}

	//TODO 实际应加密校验
	if user == nil || user.Password != req.Password { // 简单比较
		l.Logger.Errorf("invalid username or password, req: %+v", req)
		return app.LoginResponse(e.ERROR_USER, ""), nil
	}

	tokenKey := "token:" + req.Username

	// 先尝试从 Redis 获取 token
	if cachedToken, err := l.svcCtx.Redis.Get(tokenKey); err == nil && cachedToken != "" {
		// 检查 token 是否有效
		if claims, err := util.ParseToken(l.svcCtx.Config, cachedToken); err == nil && claims.Username == req.Username {
			// 获取剩余过期时间
			if ttl, err := l.svcCtx.Redis.Ttl(tokenKey); err == nil && ttl > 0 {
				l.Logger.Infof("user login successful, using cached token, username: %s, token: %s, expires: %d", req.Username, cachedToken, ttl)
				resp := app.LoginResponse(e.SUCCESS, cachedToken)
				resp.Expires = int(ttl)
				return resp, nil
			}
		}
	}

	// 如果 Redis 中没有有效 token，则生成新 token
	expiration := l.svcCtx.Config.User.AccessExpire
	token, err := util.GenerateToken(l.svcCtx.Config, req.Username, expiration)
	if err != nil {
		l.Logger.Errorf("generate token failed: %v", err)
		return app.LoginResponse(e.ERROR_USER_GENERATE_TOKEN, ""), err
	}

	// 存储 token 到 Redis
	err = l.svcCtx.Redis.Setex(tokenKey, token, int(expiration))
	if err != nil {
		l.Logger.Errorf("failed to store token in redis, token: %s, error: %v", token, err)
		return app.LoginResponse(e.ERROR_USER_STORE_TOKEN_FAIL, ""), err
	}

	// 返回成功响应
	l.Logger.Infof("user login successful, username: %s, token: %s, expires: %d", req.Username, token, expiration)
	resp = app.LoginResponse(e.SUCCESS, token)
	resp.Expires = int(expiration)
	return resp, nil
}
