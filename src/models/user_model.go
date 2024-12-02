package models

import (
	"go-boilerplate-v2/src/dtos"
	"go-boilerplate-v2/src/pkg/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    int64     `json:"user_id" gorm:"primaryKey;column:user_id"`
	Email     string    `json:"email" gorm:"unique;not null"`
	FirstName string    `json:"first_name" gorm:"not null;column:first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPassword(u.Password)

	return
}

func (u *User) FillRegister(data dtos.RegisterParam) {
	u.Email = data.Email
	u.FirstName = data.FirstName
	u.LastName = data.LastName
	u.Phone = data.Phone
	u.Password = data.Password
}
