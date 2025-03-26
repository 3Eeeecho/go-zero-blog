package model

import (
	"context"

	"gorm.io/gorm"
)

type BlogArticleLike struct {
	Id        int64 `gorm:"primaryKey;column:id"`
	ArticleId int64 `gorm:"column:article_id"`
	UserId    int64 `gorm:"column:user_id"`
	CreatedOn int64 `gorm:"column:created_on"`
}

func (BlogArticleLike) TableName() string {
	return "article_likes"
}

type (
	BlogArticleLikeModel interface {
		Insert(ctx context.Context, ArticleLike *BlogArticleLike) error
		ExistByID(ctx context.Context, id int64) (bool, error)
		FindByID(ctx context.Context, id int64) (*BlogArticleLike, error)
		CountByCondition(ctx context.Context, maps any) (int64, error)
		Delete(ctx context.Context, articleId, userId int64) error
		Exists(ctx context.Context, articleId, userId int64) (bool, error)
		//CleanAll(ctx context.Context) (bool, error) 软删除暂不考虑，后面在添加
	}

	defaultArticleLikeModel struct {
		db *gorm.DB
	}
)

// NewBlogArticleLikeModel returns a model for the database table.
func NewBlogArticleLikeModel(db *gorm.DB) BlogArticleLikeModel {
	return &defaultArticleLikeModel{db: db}
}

func (m *defaultArticleLikeModel) Insert(ctx context.Context, ArticleLike *BlogArticleLike) error {
	return m.db.WithContext(ctx).Create(ArticleLike).Error
}

func (m *defaultArticleLikeModel) ExistByID(ctx context.Context, id int64) (bool, error) {
	var ArticleLike BlogArticleLike
	err := m.db.WithContext(ctx).Select("id").Where("id=?", id).First(&ArticleLike).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return ArticleLike.Id > 0, nil
}

func (m *defaultArticleLikeModel) FindByID(ctx context.Context, id int64) (*BlogArticleLike, error) {
	var ArticleLike BlogArticleLike
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&ArticleLike).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &ArticleLike, nil
}

func (m *defaultArticleLikeModel) CountByCondition(ctx context.Context, maps any) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&BlogArticleLike{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *defaultArticleLikeModel) Delete(ctx context.Context, articleId, userId int64) error {
	return m.db.WithContext(ctx).Where("article_id = ? AND user_id = ?", articleId, userId).Delete(&BlogArticleLike{}).Error
}

func (m *defaultArticleLikeModel) Exists(ctx context.Context, articleId, userId int64) (bool, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&BlogArticleLike{}).Where("article_id = ? AND user_id = ?", articleId, userId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
