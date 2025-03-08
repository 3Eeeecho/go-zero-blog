package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导入标签信息
func NewImportTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportTagLogic {
	return &ImportTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportTagLogic) ImportTag(req *types.ImportTagRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
