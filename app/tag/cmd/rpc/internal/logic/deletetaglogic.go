package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteTag 删除文章标签
func (l *DeleteTagLogic) DeleteTag(in *pb.DeleteTagRequest) (*pb.TagCommonResponse, error) {
	// 验证标签 ID 是否合法
	if in.Id <= 0 {
		l.Logger.Errorf("invalid tag id: %d", in.Id)
		return nil, xerr.NewErrCode(xerr.REQUEST_PARAM_ERROR)
	}

	// 检查标签是否存在
	exist, err := l.svcCtx.TagModel.ExistTagByID(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("check tag existence failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "check tag existence failed: %v", err)
	}

	// 如果标签不存在，返回错误
	if !exist {
		l.Logger.Errorf("tag not found, id: %d", in.Id)
		return nil, xerr.NewErrCode(xerr.ERROR_NOT_EXIST_TAG)
	}

	// 执行删除操作
	err = l.svcCtx.TagModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to delete tag, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete tag failed: %v", err)
	}

	// 记录成功日志并返回响应
	l.Logger.Infof("tag deleted successfully, id: %d", in.Id)
	return &pb.TagCommonResponse{
		Msg: "删除标签成功",
	}, nil
}
