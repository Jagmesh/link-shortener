package user

import (
	"fmt"
	"link-shortener/entity/model"
	apperror "link-shortener/pkg/app-error"
)

type Service struct {
	*ServiceDeps
}

type ServiceDeps struct {
	Repository *Repository
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{deps}
}

func (s Service) FindByEmail(email string) (*model.User, error) {
	user := &model.User{Email: email}
	err := s.Repository.FindOne(user)

	if err != nil {
		return nil, apperror.NotFound(fmt.Sprintf("Failed to find user by '%s' email", email))
	}
	return user, nil
}

func (s Service) Create(email, password, name string) (*model.User, error) {
	if existingUser, _ := s.FindByEmail(email); existingUser != nil {
		return nil, apperror.Conflict(fmt.Sprintf("User with '%s' email already exists", email))
	}

	user := model.NewUser(email, password, name)
	_, err := s.Repository.Create(user)
	if err != nil {
		return nil, apperror.Internal("Failed to create User")
	}
	return user, nil
}
