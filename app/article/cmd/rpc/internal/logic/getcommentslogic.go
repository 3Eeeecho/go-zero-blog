package logic

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取评论列表
func (l *GetCommentsLogic) GetComments(in *pb.GetCommentsRequest) (*pb.GetCommentsResponse, error) {
	exist, err := l.svcCtx.ArticleModel.ExistArticleByID(l.ctx, in.ArticleId)
	if err != nil {
		l.Logger.Errorf("exist article_id failed,err: %v", err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	if !exist {
		l.Logger.Errorf("article not exist")
		return nil, xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND)
	}

	// 设置分页默认值
	pageNum := in.PageNum
	pageSize := in.PageSize
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	comments, err := l.svcCtx.CommentModel.GetAll(l.ctx, int(in.PageNum), int(in.PageSize))
	if err != nil {
		l.Logger.Errorf("failed to get comments, article_id: %d, error: %v", in.ArticleId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 构建树形结构
	//TODO 如果子评论的父评论在列表中靠后（未处理），
	//commentMap[c.ParentId] 会失败，导致子评论未加入树
	commentMap := make(map[int64]*pb.Comment)
	var topComments []*pb.Comment
	for _, c := range comments {
		comment := &pb.Comment{
			Id:        c.Id,
			ArticleId: c.ArticleId,
			UserId:    c.UserId,
			Content:   c.Content,
			ParentId:  c.ParentId,
			Children:  []*pb.Comment{},
		}
		commentMap[c.Id] = comment
		if c.ParentId == 0 {
			topComments = append(topComments, comment)
		} else if parent, ok := commentMap[c.ParentId]; ok {
			parent.Children = append(parent.Children, comment)
		}
	}

	// 分页处理顶级评论
	//TODO 当前分页在内存中处理，如果评论量很大，应在数据库查询时分页顶级评论
	total := int64(len(topComments))
	start := (pageNum - 1) * pageSize
	end := start + pageSize
	if start >= total {
		start = total
	}
	if end > total {
		end = total
	}
	pagedComments := topComments[start:end]

	l.Logger.Infof("comments retrieved, article_id: %d, total: %d, page_num: %d, page_size: %d",
		in.ArticleId, total, pageNum, pageSize)
	return &pb.GetCommentsResponse{
		Msg:      "获取评论列表成功",
		Comments: pagedComments,
		Total:    total,
	}, nil
}
