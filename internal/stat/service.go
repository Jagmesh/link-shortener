package stat

type Service struct {
	deps *StatServiceDeps
}

type StatServiceDeps struct {
	Repository *Repository
}

func NewService(deps *StatServiceDeps) *Service {
	return &Service{deps: deps}
}

func (s Service) AddClick(linkId uint) error {
	return s.deps.Repository.CreateOrIncrementClick(linkId)
}
