package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/utils"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增一篇文章
func (l *AddArticleLogic) AddArticle(in *pb.AddArticleRequest) (*pb.ArticleCommonResponse, error) {
	// 1.检查标签是否存在
	resp, err := l.svcCtx.TagRpc.FoundTag(l.ctx, &tagservice.FoundTagRequest{
		Id: in.TagId,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		l.Logger.Errorf("failed to check to tag existence :%v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find tag failed")
	}

	if !resp.Found {
		l.Logger.Errorf("该标签不存在 :%v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_NOT_EXIST_TAG), "tag not exist")
	}

	// 2.创建文章
	article := &model.BlogArticle{
		TagId:      in.TagId,
		Title:      in.Title,
		Desc:       in.Desc,
		Content:    in.Content,
		CreatedBy:  in.CreatedBy,
		CreatedOn:  time.Now().Unix(),
		State:      utils.StateDraft,
		ModifiedOn: 0,
		DeletedOn:  0,
	}

	id, err := l.svcCtx.ArticleModel.Insert(l.ctx, article)
	if err != nil {
		l.Logger.Errorf("failed to insert article: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert article failed")
	}

	l.Logger.Info("article added with ID: %d", id)

	return &pb.ArticleCommonResponse{
		Msg: "添加文章成功",
	}, nil
}
