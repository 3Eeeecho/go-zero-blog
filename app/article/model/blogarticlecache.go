package model

import (
	"fmt"
	"time"
)

// 文章列表缓存结构
type ArticleListCache struct {
	IDs        []int64 `json:"ids"`         // 文章ID列表
	Total      int64   `json:"total"`       // 总记录数
	UpdateTime int64   `json:"update_time"` // 更新时间
	TagID      int64   `json:"tag_id"`      // 标签ID（如果有）
	PageNum    int     `json:"page_num"`    // 页码
	PageSize   int     `json:"page_size"`   // 每页数量
}

// 缓存Key相关方法
type CacheKeys struct{}

// 获取文章列表缓存Key
func (k *CacheKeys) GetListCacheKey(tagId int64, pageNum, pageSize int) string {
	if tagId > 0 {
		return fmt.Sprintf("article:list:tag_%d:page_%d_%d", tagId, pageNum, pageSize)
	}
	return fmt.Sprintf("article:list:page_%d_%d", pageNum, pageSize)
}

// 获取文章详情缓存Key
func (k *CacheKeys) GetDetailCacheKey(articleId int64) string {
	return fmt.Sprintf("article:detail:%d", articleId)
}

// 获取文章点赞数缓存Key
func (k *CacheKeys) GetLikeCountKey(articleId int64) string {
	return fmt.Sprintf("article:likes:count:%d", articleId)
}

// 获取文章点赞集合缓存Key
func (k *CacheKeys) GetLikeSetKey(articleId int64) string {
	return fmt.Sprintf("article:likes:set:%d", articleId)
}

// 获取文章列表缓存Key的匹配模式
func (k *CacheKeys) GetListCachePattern(tagId int64) string {
	if tagId > 0 {
		return fmt.Sprintf("article:list:tag_%d:page_*_*", tagId)
	}
	return "article:list:page_*_*"
}

// 缓存过期时间相关方法
type CacheTTL struct{}

// 获取文章列表缓存过期时间
func (t *CacheTTL) GetListCacheTTL() time.Duration {
	return 1 * time.Hour
}

// 获取文章详情缓存过期时间
func (t *CacheTTL) GetDetailCacheTTL() time.Duration {
	return 24 * time.Hour
}

// 获取点赞相关缓存过期时间
func (t *CacheTTL) GetLikeCacheTTL() time.Duration {
	return 7 * 24 * time.Hour
}
