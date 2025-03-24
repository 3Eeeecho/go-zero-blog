package tag

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改文章标签
func NewEditTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditTagLogic {
	return &EditTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditTagLogic) EditTag(req *types.EditTagRequest) (resp *types.Response, err error) {
	editTagResp, err := l.svcCtx.TagServiceRpc.EditTag(l.ctx, &tagservice.EditTagRequest{
		Id:         req.Id,
		Name:       req.Name,
		ModifiedBy: req.ModifiedBy,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.Response{} // 初始化 resp
	err = copier.Copy(resp, editTagResp)
	if err != nil {
		l.Logger.Errorf("failed to copy editTagResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
