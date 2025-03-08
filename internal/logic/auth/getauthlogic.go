package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取授权 Token
func NewGetAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthLogic {
	return &GetAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAuthLogic) GetAuth(req *types.GetAuthRequest) (resp *types.GetAuthResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
