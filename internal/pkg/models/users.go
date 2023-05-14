package models

import "time"

type UserModel struct {
	ID          int64     `gorm:"column:id;primary_key"`
	Username    string    `gorm:"column:username;not null"`
	Password    string    `gorm:"column:password;not null"`
	CreatedTime time.Time `gorm:"column:created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time"`
}

func (u *UserModel) TableName() string {
	return "users"
}
