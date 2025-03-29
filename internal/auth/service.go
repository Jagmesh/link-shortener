package auth

import (
	"errors"
	"link-shortener/internal/user"
)

type Service struct {
	deps *AuthServiceDeps
}

type AuthServiceDeps struct {
	UserSerive *user.Service
}

func NewService(deps *AuthServiceDeps) *Service {
	return &Service{deps: deps}
}

func (s *Service) Login(email, password string) (*user.User, error) {
	user, err := s.deps.UserSerive.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func (s *Service) Register(email, password, name string) (*user.User, error) {
	return s.deps.UserSerive.Create(email, password, name)
}
