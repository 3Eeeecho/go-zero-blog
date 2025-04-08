package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"

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

// GetTags 获取标签列表
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

	cacheKey := fmt.Sprintf("tag:list:page_%d_%d", pageNum, pageSize)

	// 从 Redis 获取
	cached, err := l.svcCtx.Redis.Get(cacheKey)
	if err == nil && cached != "" {
		var resp pb.GetTagsResponse
		err := json.Unmarshal([]byte(cached), &resp)
		//正确解析
		if err == nil {
			l.Logger.Info("cache hit getTags!")
			return &resp, nil
		}
		l.Logger.Errorf("json unmarshal failed,err:%v", err)
	}

	if err != nil {
		//发生错误,记录下来
		l.Logger.Errorf("cache hit redis failed,err:%v", err)
	}

	// 构造查询条件，支持按名称和状态过滤
	conditions := make(map[string]any)
	// 假设 1 表示启用状态,默认为启用
	conditions["state"] = 1

	// 获取标签列表
	tags, err := l.svcCtx.TagModel.GetOnCondition(l.ctx, int(pageNum), int(pageSize), conditions)
	if err != nil {
		l.Logger.Errorf("get tags failed, page_num: %d, page_size: %d, conditions: %v, error: %v",
			pageNum, pageSize, conditions, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get tags failed: %v", err)
	}

	// 获取标签总数
	total := int64(len(tags))

	// 将标签数据转换为响应格式
	data := make([]*pb.Tag, len(tags))
	for i, tag := range tags {
		data[i] = &pb.Tag{
			Id:         tag.Id,
			Name:       tag.Name,
			State:      tag.State,
			CreatedBy:  tag.CreatedBy,
			ModifiedBy: tag.ModifiedBy,
		}
	}

	resp := &pb.GetTagsResponse{
		Msg:      "获取标签列表成功",
		Data:     data,
		Total:    total,
		PageNum:  pageNum,
		PageSize: pageSize,
	}

	// 存入 Redis
	jsonData, _ := json.Marshal(resp)
	l.svcCtx.Redis.Set(cacheKey, string(jsonData))

	// 记录成功日志并返回响应
	l.Logger.Infof("get tags successfully, page_num: %d, page_size: %d, total: %d", pageNum, pageSize, total)
	return resp, nil
}
