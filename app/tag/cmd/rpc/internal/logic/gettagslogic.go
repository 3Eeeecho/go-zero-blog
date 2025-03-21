package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagsLogic {
	return &GetTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取标签列表
func (l *GetTagsLogic) GetTags(in *pb.GetTagsRequest) (*pb.GetTagsResponse, error) {
	// 设置分页默认值，若未提供则使用默认值
	pageNum := in.PageNum
	pageSize := in.PageSize
	if pageNum <= 0 {
		pageNum = 1 // 默认第1页
	}
	if pageSize <= 0 {
		pageSize = int64(l.svcCtx.Config.App.PageSize) // 使用默认配置
	}

	// 构造查询条件，支持按名称和状态过滤
	conditions := make(map[string]any)
	if in.Name != "" {
		conditions["name"] = in.Name
	}
	if in.State != 0 { // 假设 0 表示未指定状态
		conditions["state"] = in.State
	}

	// 获取标签列表
	tags, err := l.svcCtx.TagModel.GetAll(l.ctx, int(pageNum), int(pageSize), conditions)
	if err != nil {
		l.Logger.Errorf("get tags failed, page_num: %d, page_size: %d, conditions: %v, error: %v",
			pageNum, pageSize, conditions, err)
		return nil, err
	}

	// 获取标签总数
	total, err := l.svcCtx.TagModel.CountByCondition(l.ctx, conditions)
	if err != nil {
		l.Logger.Errorf("count tags failed,condition:%v,error:%v", conditions, err)
		return nil, err
	}

	// 返回成功响应，包含标签列表和分页信息
	l.Logger.Infof("tags retrieved successfully, page_num: %d, page_size: %d, total: %d", pageNum, pageSize, total)

	data := make([]*pb.Tag, len(tags))
	err = copier.Copy(&data, tags)
	if err != nil {
		l.Logger.Errorf("copy tags failed, error: %v", err)
		return nil, err
	}

	//TODO 返回的state不对
	return &pb.GetTagsResponse{
		Msg:      "获取标签列表成功",
		Data:     data,
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
	}, nil
}
