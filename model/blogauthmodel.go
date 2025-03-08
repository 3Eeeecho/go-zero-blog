package model

import (
	"context"

	"gorm.io/gorm"
)

type BlogAuth struct {
	ID       int    `gorm:"primary_key;column:id" `
	Username string `gorm:"column:username" `
	Password string `gorm:"column:password" `
}

func (BlogAuth) TableName() string {
	return "blog_auth"
}

type (
	// BlogAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogAuthModel.
	BlogAuthModel interface {
		//TODO 使用哈希（如 bcrypt）存储密码，并在验证时比较哈希值
		CheckAuth(ctx context.Context, username, password string) (bool, error)
	}

	defaultBlogAuthModel struct {
		db *gorm.DB
	}
)

// NewBlogAuthModel returns a model for the database table.
func NewBlogAuthModel(db *gorm.DB) BlogAuthModel {
	return &defaultBlogAuthModel{db: db}
}

func (m *defaultBlogAuthModel) CheckAuth(ctx context.Context, username, password string) (bool, error) {
	var auth BlogAuth
	err := m.db.WithContext(ctx).Select("id").Where(BlogAuth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return auth.ID > 0, nil
}
