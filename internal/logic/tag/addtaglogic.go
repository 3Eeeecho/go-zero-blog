package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

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
	// 验证请求参数
	if req.Name == "" || req.State < 0 {
		return app.Response(e.INVALID_PARAMS, nil), nil
	}

	exist, err := l.svcCtx.TagModel.ExistTagByName(l.ctx, req.Name)
	if err != nil {
		return app.Response(e.ERROR_EXIST_TAG_FAIL, nil), err
	}

	//防止重复插入
	if exist {
		return app.ResponseMsg(e.ERROR, "已经存在相同标签", nil), nil
	}

	//  构造标签数据
	tag := &model.BlogTag{
		Name:      req.Name,
		State:     req.State,
		CreatedBy: req.CreatedBy,
		CreatedOn: time.Now().Unix(),
	}

	// 插入数据库
	err = l.svcCtx.TagModel.Insert(l.ctx, tag)
	if err != nil {
		l.Logger.Errorf("failed to add tag, req: %+v, error: %v", req, err)
		return app.Response(e.ERROR_ADD_TAG_FAIL, nil), err
	}

	// 构造成功响应
	l.Logger.Infof("tag added successfully, name: %s, state: %d", req.Name, req.State)
	return app.Response(e.SUCCESS, tag), nil
}
