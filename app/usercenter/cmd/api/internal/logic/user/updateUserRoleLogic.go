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

type UpdateUserRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户权限
func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRoleLogic) UpdateUserRole(req *types.UpdateUserRoleRequest) (resp *types.UpdateUserRoleResponse, err error) {
	adminId := ctxdata.GetUidFromCtx(l.ctx)
	updateUserRoleResp, err := l.svcCtx.UsercenterRpc.UpdateUserRole(l.ctx, &usercenter.UpdateUserRoleRequest{
		Id:      int64(req.Id),
		AdminId: adminId,
		Role:    req.Role,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UpdateUserRoleResponse{} // 初始化 resp
	err = copier.Copy(resp, updateUserRoleResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
