package logic

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/internal/types"
	"github.com/3Eeeecho/go-zero-blog/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/app"
	"github.com/3Eeeecho/go-zero-blog/pkg/e"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改文章
func NewEditArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditArticleLogic {
	return &EditArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditArticleLogic) EditArticle(req *types.EditArticleRequest) (resp *types.Response, err error) {
	//检查articleID
	exist, err := l.svcCtx.ArticleModel.ExistArticleByID(l.ctx, req.Id)
	if err != nil {
		l.Logger.Errorf("check article existence failed, id: %d, error: %v", req.Id, err)
		return app.Response(e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil), err
	}
	if !exist {
		l.Logger.Errorf("article not found, id: %d", req.Id)
		return app.Response(e.ERROR_NOT_EXIST_ARTICLE, nil), nil
	}

	//检查tagID
	exist, err = l.svcCtx.TagModel.ExistTagByID(l.ctx, req.TagId)
	if err != nil {
		l.Logger.Errorf("check tag existence failed, tag_id: %d, error: %v", req.TagId, err)
		return app.Response(e.ERROR_EXIST_TAG_FAIL, nil), err
	}
	if !exist {
		l.Logger.Errorf("tag not found, tag_id: %d", req.TagId)
		return app.Response(e.ERROR_NOT_EXIST_TAG, nil), nil
	}

	// 构造更新数据
	//TODO CoverImageUrl参数的处理
	article := &model.BlogArticle{
		Id:         req.Id,
		TagId:      req.TagId,
		Title:      req.Title,
		Desc:       req.Desc,
		Content:    req.Content,
		State:      req.State,
		ModifiedBy: req.ModifiedBy,
		ModifiedOn: time.Now().Unix(), // 更新修改时间
	}

	err = l.svcCtx.ArticleModel.Edit(l.ctx, req.Id, article)
	if err != nil {
		l.Logger.Errorf("failed to edit article, id: %d, error: %v", req.Id, err)
		return app.Response(e.ERROR_EDIT_ARTICLE_FAIL, nil), err
	}

	// 返回成功响应
	l.Logger.Infof("article edited successfully, id: %d", req.Id)
	return app.Response(e.SUCCESS, nil), nil
}
