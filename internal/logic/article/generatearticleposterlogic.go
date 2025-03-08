package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateArticlePosterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成文章海报
func NewGenerateArticlePosterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateArticlePosterLogic {
	return &GenerateArticlePosterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateArticlePosterLogic) GenerateArticlePoster() (resp *types.GenerateArticlePosterResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
