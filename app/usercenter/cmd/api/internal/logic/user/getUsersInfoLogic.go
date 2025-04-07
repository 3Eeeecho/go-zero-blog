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

type GetUsersInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersInfoLogic {
	return &GetUsersInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersInfoLogic) GetUsersInfo(req *types.GetUsersInfoRequest) (resp *types.GetUsersInfoResponse, err error) {
	usersInfoResp, err := l.svcCtx.UsercenterRpc.GetUsersInfo(l.ctx, &usercenter.BatchGetUsersInfoRequest{
		Ids: req.Ids,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.GetUsersInfoResponse{} // 初始化 resp
	err = copier.Copy(resp, usersInfoResp)
	if err != nil {
		l.Logger.Errorf("failed to copy usersInfoResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
