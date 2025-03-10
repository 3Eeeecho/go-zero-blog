package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

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
	// 设置分页默认值，若未提供则使用默认值
	pageNum := req.PageNum
	pageSize := req.PageSize
	if pageNum <= 0 {
		pageNum = 1 // 默认第1页
	}
	if pageSize <= 0 {
		pageSize = 10 // 默认每页10条
	}

	// 构造查询条件，支持按名称和状态过滤
	conditions := make(map[string]any)
	if req.Name != "" {
		conditions["name"] = req.Name
	}
	if req.State != 0 { // 假设 0 表示未指定状态
		conditions["state"] = req.State
	}
	l.Logger.Infof("query conditions: %v", conditions) // 调试条件

	// 获取标签列表
	tags, err := l.svcCtx.TagModel.GetAll(l.ctx, pageNum, pageSize, conditions)
	if err != nil {
		l.Logger.Errorf("get tags failed, page_num: %d, page_size: %d, conditions: %v, error: %v",
			pageNum, pageSize, conditions, err)
		return app.GetTagsResponse(e.ERROR_GET_TAGS_FAIL, nil, 0, pageNum, pageSize), err
	}

	// 获取标签总数
	total, err := l.svcCtx.TagModel.CountByCondition(l.ctx, conditions)
	if err != nil {
		l.Logger.Errorf("count tags failed,condition:%v,error:%v", conditions, err)
		return app.GetTagsResponse(e.ERROR_COUNT_TAG_FAIL, nil, 0, pageNum, pageSize), err
	}

	// 返回成功响应，包含标签列表和分页信息
	l.Logger.Infof("tags retrieved successfully, page_num: %d, page_size: %d, total: %d", pageNum, pageSize, total)
	return app.GetTagsResponse(e.SUCCESS, tags, total, pageNum, pageSize), nil
}
