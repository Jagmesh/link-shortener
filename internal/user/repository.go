package user

import (
	"errors"
	"link-shortener/pkg/database"
)

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindOne(user *User) error {
	res := r.db.Where("email = ?", user.Email).First(&user)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no user found")
	}
	return nil
}

func (r *Repository) Create(user *User) (*User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
