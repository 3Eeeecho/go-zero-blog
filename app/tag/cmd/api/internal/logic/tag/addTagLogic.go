package tag

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增文章标签
func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTagLogic) AddTag(req *types.AddTagRequest) (resp *types.Response, err error) {
	addTagResp, err := l.svcCtx.TagServiceRpc.AddTag(l.ctx, &tagservice.AddTagRequest{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		State:     int64(req.State),
	})

	if err != nil {
		return nil, err
	}

	resp = &types.Response{} // 初始化 resp
	err = copier.Copy(resp, addTagResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
