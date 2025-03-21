package tag

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取标签列表
func NewGetTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagsLogic {
	return &GetTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagsLogic) GetTags(req *types.GetTagsRequest) (resp *types.GetTagsResponse, err error) {
	editTagResp, err := l.svcCtx.TagServiceRpc.GetTags(l.ctx, &tagservice.GetTagsRequest{
		Name:     req.Name,
		State:    int64(req.State),
		PageNum:  int64(req.PageNum),
		PageSize: int64(req.PageSize),
	})

	if err != nil {
		return nil, err
	}

	resp = &types.GetTagsResponse{} // 初始化 resp
	err = copier.Copy(resp, editTagResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
