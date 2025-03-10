package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

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
	if req.Id <= 0 {
		l.Logger.Errorf("Invalid tag id:%d", req.Id)
		return app.Response(e.INVALID_PARAMS, nil), nil
	}

	exist, err := l.svcCtx.TagModel.ExistTagByID(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("check tag existence failed, id: %d, error: %v", req.Id, err)
		return app.Response(e.ERROR_EXIST_TAG_FAIL, nil), err
	}

	if !exist {
		l.Logger.Errorf("tag not found, id: %d", req.Id)
		return app.Response(e.ERROR_NOT_EXIST_TAG, nil), nil
	}

	// 执行删除操作
	err = l.svcCtx.TagModel.Delete(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("failed to delete tag, id: %d, error: %v", req.Id, err)
		return app.Response(e.ERROR_DELETE_TAG_FAIL, nil), err
	}

	// 返回成功响应
	l.Logger.Infof("tag deleted successfully, id: %d", req.Id)
	return app.Response(e.SUCCESS, nil), nil
}
