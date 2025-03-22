package model

import (
	"context"

	"gorm.io/gorm"
)

type (
	BlogUsersModel interface {
		Insert(ctx context.Context, user *BlogUser) error
		FindOne(ctx context.Context, id int64) (*BlogUser, error)
		FindByUsername(ctx context.Context, username string) (*BlogUser, error)
		FindByUserId(ctx context.Context, id int64) (*BlogUser, error)
		Update(ctx context.Context, user *BlogUser) error
	}

	defaultUserModel struct {
		db *gorm.DB
	}

	BlogUser struct {
		Id       int64  `gorm:"column:id;primaryKey"`
		Username string `gorm:"column:username"`
		Password string `gorm:"column:password"`
		Role     string `gorm:"column:role"`
	}
)

func (BlogUser) TableName() string {
	return "blog_users"
}

func NewBlogUsersModel(db *gorm.DB) BlogUsersModel {
	return &defaultUserModel{db: db}
}

func (m *defaultUserModel) Insert(ctx context.Context, user *BlogUser) error {
	return m.db.WithContext(ctx).Create(user).Error
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

func (m *defaultUserModel) Update(ctx context.Context, user *BlogUser) error {
	return m.db.WithContext(ctx).Save(user).Error
}

func (m *defaultUserModel) FindByUserId(ctx context.Context, id int64) (*BlogUser, error) {
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
