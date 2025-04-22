package utils

import (
	"context"
	"time"

	"github.com/3Eeeecho/go-zero-blog/app/article/cmd/rpc/internal/svc"
	"github.com/3Eeeecho/go-zero-blog/app/usercenter/cmd/rpc/userpb"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	StateDraft     int32 = 0 // 草稿
	StatePending   int32 = 1 // 待审核
	StatePublished int32 = 2 // 审核成功
	StateRejected  int32 = 3 // 审核失败
)

// CheckArticlePermission 检查用户是否有权限操作文章
func CheckArticlePermission(ctx context.Context, svcCtx *svc.ServiceContext, articleId, userId int64) error {
	// 检查是否是文章作者
	hasPermission, err := svcCtx.ArticleModel.CheckPermission(ctx, articleId, userId)
	if err != nil {
		logx.WithContext(ctx).Errorf("check permission failed, article_id: %d, user_id: %d, error: %v",
			articleId, userId, err)
		return err
	}
	if hasPermission {
		return nil
	}

	// 检查是否是管理员
	user, err := svcCtx.UserRpc.GetUserRole(ctx, &userpb.GetUserRoleRequest{
		Id: userId,
	})
	if err != nil {
		logx.WithContext(ctx).Errorf("get user role failed, user_id: %d, error: %v", userId, err)
		return errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get user role failed")
	}

	if user.Role != "admin" {
		logx.WithContext(ctx).Errorf("user %d has no permission to operate article %d", userId, articleId)
		return xerr.NewErrCode(xerr.ERROR_FORBIDDEN)
	}

	return nil
}

// 清理文章相关缓存
func CleanArticleCache(svcCtx *svc.ServiceContext, articleId, tagId int64) {
	// 异步执行缓存清理
	go func() {
		// 创建新的 context，不继承父 context 的取消信号
		cleanCtx := context.Background()
		// 创建新的 logger
		// 调用 cleanCacheWithDelay 函数，传入新的 context、svcCtx、articleId、tagId 和 logger
		logger := logx.WithContext(cleanCtx)
		cleanCacheWithDelay(cleanCtx, svcCtx, articleId, tagId, logger)
	}()
}

// 延迟双删清理缓存
func cleanCacheWithDelay(ctx context.Context, svcCtx *svc.ServiceContext, articleId, tagId int64, logger logx.Logger) {
	// 第一次删除
	deleteCache(ctx, svcCtx, articleId, tagId, logger)

	// 延迟第二次删除
	time.Sleep(1 * time.Second)
	deleteCache(ctx, svcCtx, articleId, tagId, logger)
}

// 删除缓存
func deleteCache(ctx context.Context, svcCtx *svc.ServiceContext, articleId, tagId int64, logger logx.Logger) {
	// 删除文章详情缓存
	detailKey := svcCtx.CacheKeys.GetDetailCacheKey(articleId)
	if _, err := svcCtx.Redis.DelCtx(ctx, detailKey); err != nil {
		logger.Errorf("delete detail cache failed, key: %s, error: %v", detailKey, err)
	}

	// 删除文章列表缓存
	// 1. 如果有 tag_id，删除该标签的列表缓存
	if tagId != 0 {
		if err := deleteListCacheByTag(ctx, svcCtx, tagId, logger); err != nil {
			logger.Errorf("delete list cache failed for tag %d, error: %v", tagId, err)
		}
	}

	// 2. 删除所有标签的列表缓存
	if err := deleteListCacheByTag(ctx, svcCtx, 0, logger); err != nil {
		logger.Errorf("delete all list cache failed, error: %v", err)
	}
}

// 删除指定标签的所有列表缓存
func deleteListCacheByTag(ctx context.Context, svcCtx *svc.ServiceContext, tagId int64, logger logx.Logger) error {
	// 构建匹配模式
	pattern := svcCtx.CacheKeys.GetListCachePattern(tagId)

	// 使用SCAN命令遍历所有匹配的键
	var cursor uint64 = 0
	for {
		keys, nextCursor, err := svcCtx.Redis.ScanCtx(ctx, cursor, pattern, 100)
		if err != nil {
			return errors.Wrapf(err, "scan redis keys failed, pattern: %s", pattern)
		}

		// 删除找到的键
		if len(keys) > 0 {
			if _, err := svcCtx.Redis.DelCtx(ctx, keys...); err != nil {
				logger.Errorf("delete keys failed, keys: %v, error: %v", keys, err)
			}
		}

		// 更新游标
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return nil
}
