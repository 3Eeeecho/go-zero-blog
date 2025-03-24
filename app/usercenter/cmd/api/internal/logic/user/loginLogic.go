package user

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/usercenter"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
	loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.LoginResponse{} // 初始化 resp
	err = copier.Copy(resp, loginResp)
	if err != nil {
		l.Logger.Errorf("failed to copy loginResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
