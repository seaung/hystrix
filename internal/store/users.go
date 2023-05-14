package store

import (
	"context"
	"errors"

	"github.com/seaung/hystrix/internal/pkg/models"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *models.UserModel) error
	GetUser(ctx context.Context, username string) (*models.UserModel, error)
	UpdateUser(ctx context.Context, user *models.UserModel) error
	DeleteUser(ctx context.Context, username string) error
}

type users struct {
	db *gorm.DB
}

// 确保users都全部实现UserStore接口
var _ UserStore = (*users)(nil)

func newUser(db *gorm.DB) *users {
	return &users{db}
}

func (u *users) CreateUser(ctx context.Context, user *models.UserModel) error {
	return u.db.Create(&user).Error
}

func (u *users) GetUser(ctx context.Context, username string) (*models.UserModel, error) {
	var user models.UserModel
	if err := u.db.Where("usesrname = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *users) UpdateUser(ctx context.Context, user *models.UserModel) error {
	return u.db.Save(user).Error
}

func (u *users) DeleteUser(ctx context.Context, username string) error {
	err := u.db.Where("username = ?", username).Delete(&models.UserModel{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
