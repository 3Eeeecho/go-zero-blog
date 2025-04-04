package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/pb"
	"github.com/3Eeeecho/go-zero-blog/app/article/model"
	"github.com/3Eeeecho/go-zero-blog/app/tag/cmd/rpc/tagservice"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
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
	// 检查权限
	hasPermission, err := l.svcCtx.ArticleModel.CheckPermission(l.ctx, in.Id, in.UserId)
	if err != nil {
		l.Logger.Errorf("failed to check permission for article %d by user %d: %v", in.Id, in.UserId, err)
		return nil, err
	}

	user, err := l.svcCtx.UserRpc.GetUserRole(l.ctx, &userpb.GetUserRoleRequest{
		Id: in.UserId,
	})
	if err != nil {
		l.Logger.Errorf("user %d is not admin: %v", in.UserId, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get user role failed: %v", err)
	}

	//既不是文章作者也不是管理员,不允许操作
	if !hasPermission && user.Role != "admin" {
		l.Logger.Errorf("user %d has no permission to delete article %d", in.UserId, in.Id)
		return nil, xerr.NewErrCode(xerr.ERROR_FORBIDDEN) // 100003: "权限不足"
	}

	//检查articleID
	originalArticle, err := l.svcCtx.ArticleModel.GetArticle(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("check article existence failed, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "exist article failed")
	}
	if originalArticle.Id <= 0 {
		l.Logger.Errorf("article not found, id: %d", in.Id)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND), "article not exist")
	}

	// 检查 tagID（如果提供）
	if in.TagId != 0 {
		//检查tagID
		foundResp, err := l.svcCtx.TagRpc.FoundTag(l.ctx, &tagservice.FoundTagRequest{
			Id: in.TagId,
		})
		if err != nil {
			l.Logger.Errorf("check tag existence failed, tag_id: %d, error: %v", in.TagId, err)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find tag failed")
		}
		if !foundResp.Found {
			l.Logger.Errorf("tag not found, tag_id: %d", in.TagId)
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ERROR_NOT_EXIST_TAG), "tag not exist")
		}
	}

	// 构造更新数据，只更新提供的字段
	updatedArticle := &model.BlogArticle{
		Id:         in.Id,
		TagId:      originalArticle.TagId,
		Title:      originalArticle.Title,
		Desc:       originalArticle.Desc,
		Content:    originalArticle.Content,
		State:      originalArticle.State,
		ModifiedBy: originalArticle.ModifiedBy,
		ModifiedOn: time.Now().Unix(),
	}

	// 检查并更新提供的字段
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
	if in.State != 0 { // 假设 State 为 0 表示未提供
		updatedArticle.State = in.State
	}
	if in.ModifiedBy != 0 {
		updatedArticle.ModifiedBy = in.ModifiedBy
	}

	err = l.svcCtx.ArticleModel.Update(l.ctx, in.Id, updatedArticle)
	if err != nil {
		l.Logger.Errorf("failed to edit article, id: %d, error: %v", in.Id, err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "edit article failed")
	}

	// 清理缓存
	l.svcCtx.Redis.Del(fmt.Sprintf("article:detail:%d", in.Id))
	if in.TagId != 0 {
		l.svcCtx.Redis.Del(fmt.Sprintf("article:list:tag_%d:page_*", in.TagId))
	} else {
		l.svcCtx.Redis.Del(fmt.Sprintf("article:list:tag_%d:page_*", originalArticle.TagId))
	}

	// 返回成功响应
	l.Logger.Infof("article edited successfully, id: %d", in.Id)

	return &pb.ArticleCommonResponse{
		Msg: "修改文章成功",
	}, nil
}
