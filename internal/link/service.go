package link

import (
	"link-shortener/entity/model"
	apperror "link-shortener/pkg/app-error"
)

type Service struct {
	*ServiceDeps
}

type ServiceDeps struct {
	Repository *Repository
}

func NewLinkService(deps *ServiceDeps) *Service {
	return &Service{deps}
}

func (s *Service) Create(url string, userId uint) (*model.Link, error) {
	existingLink, _ := s.FindOne(&FindParams{Url: url})
	if existingLink != nil {
		return nil, apperror.Conflict("Link already exists")
	}
	return s.Repository.Create(model.NewLink(url, userId))
}

func (s *Service) FindAll(params *FindParams) ([]model.Link, error) {
	links, err := s.Repository.FindAll(params)
	if err != nil || len(links) == 0 {
		return nil, apperror.NotFound("Links not found")
	}
	return links, nil
}

func (s *Service) FindOne(params *FindParams) (*model.Link, error) {
	link, err := s.Repository.FindOne(params)
	if err != nil {
		return nil, apperror.NotFound("Link not found")
	}
	return link, nil
}

func (s *Service) Delete(params *FindParams) error {
	link, err := s.FindOne(params)
	if err != nil {
		return apperror.NotFound("Link not found")
	}

	if err := s.Repository.Delete(link); err != nil {
		return apperror.Internal("Failed to delete link")
	}

	return nil
}
