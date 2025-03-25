package model

import (
	"context"

	"gorm.io/gorm"
)

type BlogComment struct {
	Id        int64  `gorm:"primaryKey;column:id"`
	ArticleId int64  `gorm:"column:article_id"`
	UserId    int64  `gorm:"column:user_id"`
	Content   string `gorm:"column:content"`
	ParentId  int64  `gorm:"parent_id"`
	CreatedOn int64  `gorm:"column:created_on"`
}

func (BlogComment) TableName() string {
	return "comments"
}

type (
	BlogCommentModel interface {
		Insert(ctx context.Context, comment *BlogComment) (int64, error)
		ExistByID(ctx context.Context, id int64) (bool, error)
		FindByID(ctx context.Context, id int64) (*BlogComment, error)
		GetAll(ctx context.Context, pageNum, pageSize int) ([]*BlogComment, error)
		CountByCondition(ctx context.Context, maps any) (int64, error)
		Delete(ctx context.Context, id int64) *gorm.DB
		Update(ctx context.Context, id int64, data any) error
		//CleanAll(ctx context.Context) (bool, error) 软删除暂不考虑，后面在添加
	}

	defaultCommentModel struct {
		db *gorm.DB
	}
)

// NewBlogCommentModel returns a model for the database table.
func NewBlogCommentModel(db *gorm.DB) BlogCommentModel {
	return &defaultCommentModel{db: db}
}

func (m *defaultCommentModel) Insert(ctx context.Context, comment *BlogComment) (int64, error) {
	err := m.db.WithContext(ctx).Create(comment).Error
	if err != nil {
		return 0, err
	}
	return comment.Id, nil
}

func (m *defaultCommentModel) ExistByID(ctx context.Context, id int64) (bool, error) {
	var comment BlogComment
	err := m.db.WithContext(ctx).Select("id").Where("id=?", id).First(&comment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return comment.Id > 0, nil
}

func (m *defaultCommentModel) FindByID(ctx context.Context, id int64) (*BlogComment, error) {
	var comment BlogComment
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&comment).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (m *defaultCommentModel) GetAll(ctx context.Context, pageNum, pageSize int) ([]*BlogComment, error) {
	var comments []*BlogComment
	if pageNum > 0 && pageSize > 0 {
		offset := (pageNum - 1) * pageSize
		err := m.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(&comments).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := m.db.WithContext(ctx).Find(&comments).Error
		if err != nil {
			return nil, err
		}
	}
	return comments, nil
}

func (m *defaultCommentModel) CountByCondition(ctx context.Context, maps any) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&BlogComment{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *defaultCommentModel) Delete(ctx context.Context, id int64) *gorm.DB {
	return m.db.WithContext(ctx).Model(&BlogComment{}).Where("id = ?", id).Delete(&BlogComment{})
}

func (m *defaultCommentModel) Update(ctx context.Context, id int64, data any) error {
	return m.db.WithContext(ctx).Model(&BlogComment{}).Where("id = ?", id).Updates(data).Error
}
