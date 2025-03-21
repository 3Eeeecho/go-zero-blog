package logic

import (
	"context"
	"errors"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/tag/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增文章标签
func (l *AddTagLogic) AddTag(in *pb.AddTagRequest) (*pb.TagCommonResponse, error) {
	// 验证请求参数
	if in.Name == "" || in.State < 0 {
		return nil, errors.New("不合法的参数")
	}

	exist, err := l.svcCtx.TagModel.ExistTagByName(l.ctx, in.Name)
	if err != nil {
		return nil, err
	}

	//防止重复插入
	if exist {
		return nil, errors.New("已经存在相同标签")
	}

	//  构造标签数据
	tag := &model.BlogTag{
		Name:      in.Name,
		State:     int(in.State),
		CreatedBy: in.CreatedBy,
		CreatedOn: time.Now().Unix(),
	}

	// 插入数据库
	err = l.svcCtx.TagModel.Insert(l.ctx, tag)
	if err != nil {
		l.Logger.Errorf("failed to add tag, in: %+v, error: %v", in, err)
		return nil, err
	}

	// 构造成功响应
	l.Logger.Infof("tag added successfully, name: %s, state: %d", in.Name, in.State)

	return &pb.TagCommonResponse{
		Msg: "添加标签成功",
	}, nil
}
