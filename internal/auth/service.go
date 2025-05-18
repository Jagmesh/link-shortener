package auth

import (
	"link-shortener/entity/model"
	apperror "link-shortener/pkg/app-error"
)

type userService interface {
	Create(email, password, name string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type Service struct {
	*ServiceDeps
}

type ServiceDeps struct {
	UserService userService
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{deps}
}

func (s *Service) Login(email, password string) (*model.User, error) {
	user, err := s.UserService.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, apperror.BadRequest("Wrong password")
	}

	return user, nil
}

func (s *Service) Register(email, password, name string) (*model.User, error) {
	return s.UserService.Create(email, password, name)
}
