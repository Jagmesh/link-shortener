package user

import apperror "link-shortener/pkg/app-error"

type Service struct {
	deps *UserServiceDeps
}

type UserServiceDeps struct {
	Repository *Repository
}

func NewService(deps *UserServiceDeps) *Service {
	return &Service{deps: deps}
}

func (s Service) Create(email, password, name string) (*User, error) {
	user := NewUser(email, password, name)
	_, err := s.deps.Repository.Create(user)
	if err != nil {
		return nil, apperror.Internal("Failed to create User")
	}
	return user, nil
}
