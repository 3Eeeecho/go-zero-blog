package tag

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 导出标签信息
func NewExportTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportTagLogic {
	return &ExportTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportTagLogic) ExportTag(req *types.ExportTagRequest) (resp *types.ExportTagResponse, err error) {
	exportTagResp, err := l.svcCtx.TagServiceRpc.ExportTag(l.ctx, &tagservice.ExportTagRequest{
		Name:  req.Name,
		State: int64(req.State),
	})

	if err != nil {
		return nil, err
	}

	resp = &types.ExportTagResponse{} // 初始化 resp
	err = copier.Copy(resp, exportTagResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
