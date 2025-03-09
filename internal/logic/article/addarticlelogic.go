package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增一篇文章
func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleRequest) (resp *types.Response, err error) {
	//1.检查标签是否存在
	_, err = l.svcCtx.TagModel.ExistTagByID(l.ctx, req.TagId)
	if err != nil {
		if err == model.ErrNotFound {
			return app.Response(e.ERROR_NOT_EXIST_TAG, nil), nil
		}
		l.Logger.Errorf("failed to check to tag existence :%v", err)
		return nil, fmt.Errorf("%s: %v", e.GetMsg(e.ERROR_EXIST_TAG_FAIL), err)
	}

	//2.创建文章
	article := &model.BlogArticle{
		TagId:      req.TagId,
		Title:      req.Title,
		Desc:       req.Desc,
		Content:    req.Content,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  time.Now().Unix(),
		State:      req.State,
		ModifiedOn: 0,
		DeletedOn:  0,
	}

	id, err := l.svcCtx.ArticleModel.Insert(l.ctx, article)
	if err != nil {
		l.Logger.Errorf("failed to insert article: %v", err)
		return nil, fmt.Errorf("%s: %v", e.GetMsg(e.ERROR_ADD_ARTICLE_FAIL), err)
	}

	l.Logger.Info("article added with ID: %d", id)
	return app.Response(e.SUCCESS, nil), nil
}
