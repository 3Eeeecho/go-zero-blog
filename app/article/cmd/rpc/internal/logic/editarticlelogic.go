package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/utils"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditArticleLogic {
	return &EditArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改文章
func (l *EditArticleLogic) EditArticle(in *pb.EditArticleRequest) (*pb.ArticleCommonResponse, error) {
	// 1. 检查文章是否存在
	originalArticle, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("get article failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get article failed")
	}
	if originalArticle.Id <= 0 {
		l.Logger.Errorf("article not found, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "article not exist")
	}

	// 2. 检查权限
	if err := utils.CheckArticlePermission(l.ctx, l.svcCtx, in.Id, in.UserId); err != nil {
		return nil, err
	}

	// 3. 更新文章
	if err := l.updateArticle(in, originalArticle); err != nil {
		return nil, err
	}

	// 4. 清理缓存（延迟双删）
	fmt.Println("in.Id", in.Id)
	fmt.Println("in.TagId", in.TagId)
	utils.CleanArticleCache(l.svcCtx, in.Id, in.TagId)

	// 返回成功响应
	l.Logger.Infof("article edited successfully, id: %d", in.Id)
	return &pb.ArticleCommonResponse{
		Msg: "修改文章成功",
	}, nil
}

// 更新文章
func (l *EditArticleLogic) updateArticle(in *pb.EditArticleRequest, originalArticle *model.BlogArticle) error {
	// 复制原始文章数据
	updatedArticle := &model.BlogArticle{}
	if err := copier.Copy(updatedArticle, originalArticle); err != nil {
		l.Logger.Errorf("copy article failed, id: %d, error: %v", originalArticle.Id, err)
		return errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "copy article failed")
	}

	// 更新提供的字段
	if in.TagId != 0 {
		updatedArticle.TagId = in.TagId
	}
	if in.Title != "" {
		updatedArticle.Title = in.Title
	}
	if in.Desc != "" {
		updatedArticle.Desc = in.Desc
	}
	if in.Content != "" {
		updatedArticle.Content = in.Content
	}
	if in.State != 0 {
		updatedArticle.State = in.State
	}
	if in.ModifiedBy != 0 {
		updatedArticle.ModifiedBy = in.ModifiedBy
	}

	// 设置更新时间
	updatedArticle.ModifiedOn = time.Now().Unix()

	// 执行更新
	if err := l.svcCtx.ArticleModel.Update(l.ctx, in.Id, updatedArticle); err != nil {
		l.Logger.Errorf("update article failed, id: %d, error: %v", in.Id, err)
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update article failed")
	}

	return nil
}
