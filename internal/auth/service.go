package auth

import "link-shortener/internal/user"

type Service struct {
	deps *AuthServiceDeps
}

type AuthServiceDeps struct {
	UserSerive *user.Service
}

func NewService(deps *AuthServiceDeps) *Service {
	return &Service{deps: deps}
}

func (s *Service) Register(email, password, name string) (*user.User, error) {
	return s.deps.UserSerive.Create(email, password, name)
}
