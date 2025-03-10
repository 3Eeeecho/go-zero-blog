package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

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
	// 检查标签 ID 是否有效，小于等于 0 则为无效参数
	if req.Id <= 0 {
		l.Logger.Errorf("invalid tag id: %d", req.Id)
		return app.Response(e.INVALID_PARAMS, nil), nil
	}

	// 检查标签是否存在，避免修改不存在的标签
	exist, err := l.svcCtx.TagModel.ExistTagByID(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("check tag existence failed, id: %d, error: %v", req.Id, err)
		return app.Response(e.ERROR_EXIST_TAG_FAIL, nil), err
	}
	if !exist {
		l.Logger.Errorf("tag not found, id: %d", req.Id)
		return app.Response(e.ERROR_NOT_EXIST_TAG, nil), nil
	}

	// 构造更新数据，映射到 BlogTag 结构
	data := map[string]any{
		"name":        req.Name,
		"modified_by": req.ModifiedBy,
		"state":       req.State,
		"modified_on": time.Now().Unix(),
	}

	// 执行修改操作，调用模型层的 Edit 方法
	err = l.svcCtx.TagModel.Edit(l.ctx, req.Id, data)
	if err != nil {
		l.Logger.Errorf("failed to edit tag, id: %d, error: %v", req.Id, err)
		return app.Response(e.ERROR_EDIT_TAG_FAIL, nil), err
	}

	// 修改成功，记录日志并返回成功响应
	l.Logger.Infof("tag edited successfully, id: %d", req.Id)
	return app.Response(e.SUCCESS, nil), nil
}
