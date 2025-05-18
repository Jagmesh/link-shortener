package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Links    []Link
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func NewUser(email, password, name string) *User {
	return &User{
		Email:    email,
		Name:     name,
		Password: password,
	}
}
