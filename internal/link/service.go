package link

import apperror "link-shortener/pkg/app-error"

type Service struct {
	LinkServiceDeps
}

type LinkServiceDeps struct {
	Repository *Repository
}

func NewLinkService(deps LinkServiceDeps) *Service {
	return &Service{deps}
}

func (s *Service) Create(url string) (*Link, error) {
	existingLink, _ := s.Find(&FindParams{url: url})
	if existingLink != nil {
		return nil, apperror.Conflict("Link already exists")
	}
	return s.Repository.Create(NewLink(url))
}

func (s *Service) Find(params *FindParams) (*Link, error) {
	return s.Repository.FindFirst(params)
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
