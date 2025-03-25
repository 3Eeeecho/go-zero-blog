package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加评论
func (l *AddCommentLogic) AddComment(in *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	// 检查文章是否存在
	exist, err := l.svcCtx.ArticleModel.ExistArticleByID(l.ctx, in.ArticleId)
	if err != nil {
		l.Logger.Errorf("failed to find article, id: %d, error: %v", in.ArticleId, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	if !exist {
		l.Logger.Errorf("article not exist, id: %d, error: %v", in.ArticleId, err)
		return nil, xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND)
	}

	if in.ParentId != 0 {
		// 查询父评论
		parentComment, err := l.svcCtx.CommentModel.FindByID(l.ctx, in.ParentId)
		if err != nil {
			l.Logger.Errorf("failed to find parent comment, id: %d, error: %v", in.ParentId, err)
			return nil, xerr.NewErrCode(xerr.DB_ERROR)
		}
		if parentComment == nil {
			l.Logger.Errorf("parent comment not found, parent_id: %d", in.ParentId)
			return nil, xerr.NewErrCode(xerr.ERROR_NOT_EXIST_COMMENT)
		}
		// 确保父评论是顶级评论（parent_id = 0）
		if parentComment.ParentId != 0 {
			l.Logger.Errorf("cannot reply to a sub-comment, parent_id: %d has parent: %d", in.ParentId, parentComment.ParentId)
			return nil, xerr.NewErrCodeMsg(xerr.REQUEST_PARAM_ERROR, "不能回复子评论")
		}
	}

	// 构造评论数据
	comment := &model.BlogComment{
		ArticleId: in.ArticleId,
		UserId:    in.UserId,
		Content:   in.Content,
		ParentId:  in.ParentId,
		CreatedOn: time.Now().Unix(),
	}

	// 插入数据库
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		l.Logger.Errorf("failed to insert comment, req: %+v, error: %v", in, err)
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	l.Logger.Infof("comment added successfully, article_id: %d, user_id: %d", in.ArticleId, in.UserId)
	return &pb.AddCommentResponse{
		Msg: "评论添加成功",
	}, nil
}
