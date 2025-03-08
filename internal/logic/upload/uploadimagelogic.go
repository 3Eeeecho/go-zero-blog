package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpLoadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传图片
func NewUpLoadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpLoadImageLogic {
	return &UpLoadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpLoadImageLogic) UpLoadImage(req *types.UpLoadImageRequest) (resp *types.UpLoadImageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
