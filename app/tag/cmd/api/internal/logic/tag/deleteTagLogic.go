package tag

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/api/internal/types"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章标签
func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagLogic) DeleteTag(req *types.DeleteTagRequest) (resp *types.Response, err error) {
	delTagResp, err := l.svcCtx.TagServiceRpc.DeleteTag(l.ctx, &tagservice.DeleteTagRequest{
		Id: int64(req.Id),
	})

	if err != nil {
		return nil, err
	}

	resp = &types.Response{} // 初始化 resp
	err = copier.Copy(resp, delTagResp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
