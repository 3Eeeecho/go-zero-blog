package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FoundTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFoundTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FoundTagLogic {
	return &FoundTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FoundTagLogic) FoundTag(in *pb.FoundTagRequest) (*pb.FoundTagResponse, error) {
	found, err := l.svcCtx.TagModel.ExistTagByID(l.ctx, in.Id)
	return &pb.FoundTagResponse{
		Found: found,
	}, err
}
