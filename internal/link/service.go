package link

import (
	"link-shortener/model"
	apperror "link-shortener/pkg/app-error"
)

type Service struct {
	LinkServiceDeps
}

type LinkServiceDeps struct {
	Repository *Repository
}

func NewLinkService(deps LinkServiceDeps) *Service {
	return &Service{deps}
}

func (s *Service) Create(url string, userId uint) (*model.Link, error) {
	existingLink, _ := s.Find(&FindParams{url: url})
	if existingLink != nil {
		return nil, apperror.Conflict("Link already exists")
	}
	return s.Repository.Create(model.NewLink(url, userId))
}

func (s *Service) Find(params *FindParams) (*model.Link, error) {
	link, err := s.Repository.FindFirst(params)
	if err != nil {
		return nil, apperror.NotFound("Link not found")
	}
	return link, nil
}

func (s *Service) Delete(params *FindParams) error {
	link, err := s.Find(params)
	if err != nil {
		return apperror.NotFound("Link not found")
	}

	if err := s.Repository.Delete(link); err != nil {
		return apperror.Internal("Failed to delete link")
	}

	return nil
}
