package user

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/usercenter"
	"github.com/3Eeeecho/go-zero-blog/pkg/ctxdata"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改密码
func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordRequest) (resp *types.UpdatePasswordResponse, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	updatePwResp, err := l.svcCtx.UsercenterRpc.UpdatePassword(l.ctx, &usercenter.UpdatePasswordRequest{
		Id:          userId,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UpdatePasswordResponse{} // 初始化 resp
	err = copier.Copy(resp, updatePwResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
