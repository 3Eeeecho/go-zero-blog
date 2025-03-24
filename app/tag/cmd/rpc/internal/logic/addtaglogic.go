package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/tag/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"

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

// AddTag 新增文章标签
func (l *AddTagLogic) AddTag(in *pb.AddTagRequest) (*pb.TagCommonResponse, error) {
	// 验证请求参数是否合法
	if in.Name == "" || in.State < 0 {
		l.Logger.Errorf("invalid parameters, name: %s, state: %d", in.Name, in.State)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	// 检查标签名称是否已存在
	exist, err := l.svcCtx.TagModel.ExistTagByName(l.ctx, in.Name)
	if err != nil {
		l.Logger.Errorf("failed to check tag existence by name %s: %v", in.Name, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "check tag existence failed: %v", err)
	}

	// 如果标签已存在，返回错误
	if exist {
		l.Logger.Errorf("tag already exists, name: %s", in.Name)
		return nil, xerr.NewErrCode(xerr.ERROR_EXIST_TAG)
	}

	// 构造标签数据
	tag := &model.BlogTag{
		Name:      in.Name,
		State:     in.State,
		CreatedBy: in.CreatedBy,
		CreatedOn: time.Now().Unix(),
	}

	// 将标签插入数据库
	err = l.svcCtx.TagModel.Insert(l.ctx, tag)
	if err != nil {
		l.Logger.Errorf("failed to add tag, in: %+v, error: %v", in, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert tag failed: %v", err)
	}

	// 记录成功日志并返回响应
	l.Logger.Infof("tag added successfully, name: %s, state: %d", in.Name, in.State)
	return &pb.TagCommonResponse{
		Msg: "添加标签成功",
	}, nil
}
