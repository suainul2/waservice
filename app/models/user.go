package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email,isUnique=users"`
	Password string `json:"password" validate:"required"`
	Role     int    `json:"role" validate:"numeric"`
	Token    string `json:"token" gorm:"unique"`
}

func (u *User) GetRole() map[string]int {
	return map[string]int{
		"admin":  0,
		"client": 1,
	}
}
