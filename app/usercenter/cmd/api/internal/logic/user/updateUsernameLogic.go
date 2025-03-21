package user

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/usercenter"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户名
func NewUpdateUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUsernameLogic {
	return &UpdateUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUsernameLogic) UpdateUsername(req *types.UpdateUsernameRequest) (resp *types.UpdateUsernameResponse, err error) {
	updateNameResp, err := l.svcCtx.UsercenterRpc.UpdatePassword(l.ctx, &usercenter.UpdatePasswordRequest{
		NewPassword: req.NewUsername,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UpdateUsernameResponse{} // 初始化 resp
	err = copier.Copy(resp, updateNameResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
