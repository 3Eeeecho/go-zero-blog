package model

import (
	"context"

	"github.com/3Eeeecho/go-zero-blog/app/tag/model"
	"github.com/3Eeeecho/go-zero-blog/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type BlogArticle struct {
	Id         int64         `gorm:"primaryKey;column:id"`
	TagId      int64         `gorm:"column:tag_id;index"`
	Tag        model.BlogTag `gorm:"foreignKey:TagId"`
	Title      string        `gorm:"column:title"`
	Desc       string        `gorm:"column:desc"`
	Content    string        `gorm:"column:content"`
	State      int32         `gorm:"column:state"`
	CreatedBy  int64         `gorm:"column:created_by"`
	CreatedOn  int64         `gorm:"column:created_on"`
	ModifiedBy int64         `gorm:"column:modified_by"`
	ModifiedOn int64         `gorm:"column:modified_on"`
	DeletedOn  int64         `gorm:"column:deleted_on"`
}

func (BlogArticle) TableName() string {
	return "articles"
}

type (
	BlogArticleModel interface {
		Insert(ctx context.Context, data *BlogArticle) (int64, error)
		ExistArticleByID(ctx context.Context, id int64) (bool, error)
		GetArticle(ctx context.Context, id int64) (*BlogArticle, error)
		GetArticles(ctx context.Context, pageNum, pageSize int, maps any) ([]*BlogArticle, error)
		CountByCondition(ctx context.Context, maps any) (int64, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, id int64, data any) error
		CheckPermission(ctx context.Context, articleId, operatorId int64) (bool, error)
		IncrementViews(id int64, count int64) error
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

func (m *defaultArticleModel) Insert(ctx context.Context, data *BlogArticle) (int64, error) {
	err := m.db.WithContext(ctx).Create(data).Error
	if err != nil {
		return 0, err
	}
	return data.Id, nil
}

func (m *defaultArticleModel) ExistArticleByID(ctx context.Context, id int64) (bool, error) {
	var article BlogArticle
	err := m.db.WithContext(ctx).Select("id").Where("id=?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return article.Id > 0, nil
}

func (m *defaultArticleModel) GetArticle(ctx context.Context, id int64) (*BlogArticle, error) {
	var article BlogArticle

	err := m.db.WithContext(ctx).Preload("Tag").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, err
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

func (m *defaultArticleModel) Delete(ctx context.Context, id int64) error {
	//TODO 事务(删除文章,一并删除评论和浏览量)
	return m.db.WithContext(ctx).Model(&BlogArticle{}).Where("id = ?", id).Delete(&BlogArticle{}).Error
}

func (m *defaultArticleModel) Update(ctx context.Context, id int64, data any) error {
	return m.db.WithContext(ctx).Model(&BlogArticle{}).Where("id = ?", id).Updates(data).Error
}

// CheckOperatorPermission 检查操作用户是否有权限（创建者或管理员）
func (m *defaultArticleModel) CheckPermission(ctx context.Context, articleId, operatorId int64) (bool, error) {
	// 查询文章
	var article BlogArticle
	result := m.db.WithContext(ctx).Select("id, created_by").Where("id = ?", articleId).First(&article)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, xerr.NewErrCode(xerr.ARTICLE_NOT_FOUND) // 102001: "该文章不存在"
		}
		return false, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to get article id %d: %v", articleId, result.Error)
	}

	// 检查是否为创建者
	if article.CreatedBy == operatorId {
		return true, nil
	}

	return false, nil
}

func (m *defaultArticleModel) IncrementViews(id int64, count int64) error {
	return m.db.Model(&BlogArticle{}).Where("id =?", id).Update("views", gorm.Expr("views + ?", count)).Error
}
