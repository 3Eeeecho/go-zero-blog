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
	if req.State < 0 || req.State > 1 {
		l.Logger.Errorf("invalid tag state: %d", req.State)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR), "invalid tag state: %d", req.State)
	}

	addTagResp, err := l.svcCtx.TagServiceRpc.AddTag(l.ctx, &tagservice.AddTagRequest{
		Name:      req.Name,
		CreatedBy: req.CreatedBy,
		State:     int64(req.State),
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	resp = &types.Response{} // 初始化 resp
	err = copier.Copy(resp, addTagResp)
	if err != nil {
		l.Logger.Errorf("failed to copy addTagResp: %v", err)
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	return resp, nil
}
