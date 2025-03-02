package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func NewUser(email, password, name string) *User {
	return &User{
		Email:    email,
		Name:     name,
		Password: password,
	}
}
