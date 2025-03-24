package model

import (
	"context"

	"gorm.io/gorm"
)

type BlogTag struct {
	Id         int64  `gorm:"primaryKey;column:id"`
	Name       string `gorm:"column:name"`
	State      int64  `gorm:"column:state"`
	CreatedBy  string `gorm:"column:created_by"`
	CreatedOn  int64  `gorm:"column:created_on"`
	ModifiedBy string `gorm:"column:modified_by"`
	ModifiedOn int64  `gorm:"column:modified_on"`
	DeletedOn  int64  `gorm:"column:deleted_on"`
}

func (BlogTag) TableName() string {
	return "blog_tags"
}

type (
	BlogTagModel interface {
		Insert(ctx context.Context, data *BlogTag) error
		InsertBatch(ctx context.Context, tags []*BlogTag) error
		ExistTagByID(ctx context.Context, id int64) (bool, error)
		ExistTagByName(ctx context.Context, name string) (bool, error)
		GetAll(ctx context.Context, pageNum, pageSize int, maps any) ([]*BlogTag, error)
		CountByCondition(ctx context.Context, maps any) (int64, error)
		Delete(ctx context.Context, id int64) error
		Edit(ctx context.Context, id int64, data any) error
		//CleanAll(ctx context.Context) (bool, error) 软删除暂不考虑，后面在添加
	}

	defaultTagModel struct {
		db *gorm.DB
	}
)

// NewBlogTagModel returns a model for the database table.
func NewBlogTagModel(db *gorm.DB) BlogTagModel {
	return &defaultTagModel{db: db}
}

func (m *defaultTagModel) Insert(ctx context.Context, data *BlogTag) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultTagModel) InsertBatch(ctx context.Context, data []*BlogTag) error {
	if len(data) == 0 {
		return nil // 没有数据直接返回
	}

	// 使用 GORM 的 Create 方法进行批量插入
	err := m.db.WithContext(ctx).Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *defaultTagModel) ExistTagByID(ctx context.Context, id int64) (bool, error) {
	var tag BlogTag
	err := m.db.WithContext(ctx).Select("id").Where("id=?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return tag.Id > 0, nil
}

func (m *defaultTagModel) ExistTagByName(ctx context.Context, name string) (bool, error) {
	var tag BlogTag
	err := m.db.WithContext(ctx).Select("id").Where("name=?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return tag.Id > 0, nil
}

func (m *defaultTagModel) GetAll(ctx context.Context, pageNum, pageSize int, maps any) ([]*BlogTag, error) {
	var (
		tags []*BlogTag
		err  error
	)
	query := m.db.WithContext(ctx)

	// 应用过滤条件
	if maps != nil {
		query = query.Where(maps)
	}

	// 处理分页
	if pageNum > 0 && pageSize > 0 {
		offset := (pageNum - 1) * pageSize // 修正偏移量
		err = query.Offset(offset).Limit(pageSize).Find(&tags).Error
	} else {
		err = query.Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (m *defaultTagModel) CountByCondition(ctx context.Context, maps any) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&BlogTag{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *defaultTagModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Model(&BlogTag{}).Where("id = ?", id).Delete(&BlogTag{}).Error
}

func (m *defaultTagModel) Edit(ctx context.Context, id int64, data any) error {
	return m.db.WithContext(ctx).Model(&BlogTag{}).Where("id = ?", id).Updates(data).Error
}
