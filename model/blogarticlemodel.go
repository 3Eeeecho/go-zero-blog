package model

import (
	"context"

	"gorm.io/gorm"
)

type BlogArticle struct {
	Id         int     `gorm:"primaryKey;column:id"`
	TagId      int     `gorm:"column:tag_id;index"`
	Tag        BlogTag `gorm:"foreignKey:TagId"`
	Title      string  `gorm:"column:title"`
	Desc       string  `gorm:"column:desc"`
	Content    string  `gorm:"column:content"`
	State      int     `gorm:"column:state"`
	CreatedBy  string  `gorm:"column:created_by"`
	CreatedOn  int64   `gorm:"column:created_on"`
	ModifiedBy string  `gorm:"column:modified_by"`
	ModifiedOn int64   `gorm:"column:modified_on"`
	DeletedOn  int64   `gorm:"column:deleted_on"`
}

func (BlogArticle) TableName() string {
	return "blog_article"
}

type (
	BlogArticleModel interface {
		Insert(ctx context.Context, data *BlogArticle) (int, error)
		ExistArticleByID(ctx context.Context, id int) (bool, error)
		GetArticle(ctx context.Context, id int) (*BlogArticle, error)
		GetArticles(ctx context.Context, pageNum, pageSize int, maps any) ([]*BlogArticle, error)
		CountByCondition(ctx context.Context, maps any) (int64, error)
		Delete(ctx context.Context, id int) error
		Edit(ctx context.Context, id int, data any) error
		//CleanAll(ctx context.Context) (bool, error) 软删除暂不考虑，后面在添加
	}

	defaultArticleModel struct {
		db *gorm.DB
	}
)

// NewBlogArticleModel returns a model for the database table.
func NewBlogArticleModel(db *gorm.DB) BlogArticleModel {
	return &defaultArticleModel{db: db}
}

func (m *defaultArticleModel) Insert(ctx context.Context, data *BlogArticle) (int, error) {
	err := m.db.WithContext(ctx).Create(data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (m *defaultArticleModel) ExistArticleByID(ctx context.Context, id int) (bool, error) {
	var article BlogArticle
	err := m.db.WithContext(ctx).Select("id").Where("id=?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return article.Id > 0, nil
}

func (m *defaultArticleModel) GetArticle(ctx context.Context, id int) (*BlogArticle, error) {
	var article BlogArticle

	err := m.db.WithContext(ctx).Preload("Tag").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func (m *defaultArticleModel) GetArticles(ctx context.Context, pageNum, pageSize int, maps any) ([]*BlogArticle, error) {
	var articles []*BlogArticle
	query := m.db.WithContext(ctx).Preload("Tag")
	if maps != nil {
		query = query.Where(maps)
	}
	if pageNum > 0 && pageSize > 0 {
		offset := (pageNum - 1) * pageSize
		err := query.Offset(offset).Limit(pageSize).Find(&articles).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := query.Find(&articles).Error
		if err != nil {
			return nil, err
		}
	}
	return articles, nil
}

func (m *defaultArticleModel) CountByCondition(ctx context.Context, maps any) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&BlogArticle{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *defaultArticleModel) Delete(ctx context.Context, id int) error {
	return m.db.WithContext(ctx).Model(&BlogArticle{}).Where("id = ?", id).Delete(&BlogArticle{}).Error
}

func (m *defaultArticleModel) Edit(ctx context.Context, id int, data any) error {
	return m.db.WithContext(ctx).Model(&BlogArticle{}).Where("id = ?", id).Updates(data).Error
}
