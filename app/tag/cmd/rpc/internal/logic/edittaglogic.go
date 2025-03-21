package logic

import (
	"context"
	"errors"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditTagLogic {
	return &EditTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章标签
func (l *EditTagLogic) EditTag(in *pb.EditTagRequest) (*pb.TagCommonResponse, error) {
	// 检查标签 ID 是否有效，小于等于 0 则为无效参数
	if in.Id <= 0 {
		l.Logger.Errorf("invalid tag id: %d", in.Id)
		return nil, errors.New("不合法的参数")
	}

	// 检查标签是否存在，避免修改不存在的标签
	exist, err := l.svcCtx.TagModel.ExistTagByID(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("check tag existence failed, id: %d, error: %v", in.Id, err)
		return nil, err
	}
	if !exist {
		l.Logger.Errorf("tag not found, id: %d", in.Id)
		return nil, errors.New("该标签不存在")
	}

	// 构造更新数据，映射到 BlogTag 结构
	data := map[string]any{
		"name":        in.Name,
		"modified_by": in.ModifiedBy,
		"state":       in.State,
		"modified_on": time.Now().Unix(),
	}

	// 执行修改操作，调用模型层的 Edit 方法
	err = l.svcCtx.TagModel.Edit(l.ctx, in.Id, data)
	if err != nil {
		l.Logger.Errorf("failed to edit tag, id: %d, error: %v", in.Id, err)
		return nil, err
	}

	// 修改成功，记录日志并返回成功响应
	l.Logger.Infof("tag edited successfully, id: %d", in.Id)
	return &pb.TagCommonResponse{
		Msg: "修改标签成功",
	}, nil
}
