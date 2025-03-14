package model

import (
	"context"

	"gorm.io/gorm"
)

type (
	BlogUserModel interface {
		Insert(ctx context.Context, data *BlogUser) error
		FindOne(ctx context.Context, id int64) (*BlogUser, error)
		FindByUsername(ctx context.Context, username string) (*BlogUser, error)
	}

	defaultUserModel struct {
		db *gorm.DB
	}

	BlogUser struct {
		Id       int64  `gorm:"column:id;primaryKey"`
		Username string `gorm:"column:username"`
		Password string `gorm:"column:password"`
	}
)

func (BlogUser) TableName() string {
	return "blog_user"
}

func NewBlogUserModel(db *gorm.DB) BlogUserModel {
	return &defaultUserModel{db: db}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *BlogUser) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*BlogUser, error) {
	var user BlogUser
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 未找到记录，返回 nil
		}
		return nil, err
	}
	return &user, nil
}

func (m *defaultUserModel) FindByUsername(ctx context.Context, username string) (*BlogUser, error) {
	var user BlogUser
	err := m.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 未找到记录，返回 nil
		}
		return nil, err
	}
	return &user, nil
}
