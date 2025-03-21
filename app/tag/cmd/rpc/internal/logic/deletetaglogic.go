package logic

import (
	"context"
	"errors"

	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/pb"

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

// 删除文章标签
func (l *DeleteTagLogic) DeleteTag(in *pb.DeleteTagRequest) (*pb.TagCommonResponse, error) {
	if in.Id <= 0 {
		l.Logger.Errorf("Invalid tag id:%d", in.Id)
		return nil, errors.New("不合法的参数")
	}
	exist, err := l.svcCtx.TagModel.ExistTagByID(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("check tag existence failed, id: %d, error: %v", in.Id, err)
		return nil, err
	}

	if !exist {
		l.Logger.Errorf("tag not found, id: %d", in.Id)
		return nil, errors.New("该标签不存在")
	}

	// 执行删除操作
	err = l.svcCtx.TagModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("failed to delete tag, id: %d, error: %v", in.Id, err)
		return nil, err
	}

	// 返回成功响应
	l.Logger.Infof("tag deleted successfully, id: %d", in.Id)
	return &pb.TagCommonResponse{
		Msg: "删除标签成功",
	}, nil
}
