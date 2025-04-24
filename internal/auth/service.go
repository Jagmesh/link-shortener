package auth

import (
	"link-shortener/model"
	apperror "link-shortener/pkg/app-error"
)

type userService interface {
	Create(email, password, name string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type Service struct {
	deps *AuthServiceDeps
}

type AuthServiceDeps struct {
	UserSerive userService
}

func NewService(deps *AuthServiceDeps) *Service {
	return &Service{deps: deps}
}

func (s *Service) Login(email, password string) (*model.User, error) {
	user, err := s.deps.UserSerive.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, apperror.BadRequest("Wrong password")
	}

	return user, nil
}

func (s *Service) Register(email, password, name string) (*model.User, error) {
	return s.deps.UserSerive.Create(email, password, name)
}
