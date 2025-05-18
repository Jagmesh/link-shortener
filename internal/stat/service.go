package stat

import (
	"errors"
	"link-shortener/entity/event"
	"link-shortener/pkg/bus"
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

func (s *Service) ListenClick(eventBus *bus.EventBus) error {
	for ev := range eventBus.Consume(event.ClickEventName) {
		id, ok := ev.Data.(uint)
		if !ok {
			return errors.New("event Data is not uint")
		}
		if s.Repository.CreateOrIncrementClick(id) != nil {
			return errors.New("failed to add click")
		}
	}
	return nil
}

func (s *Service) GetClicksNumberByDate(linksId []uint, from string, to string) (uint, error) {
	stat, err := s.Repository.GetClicksCountByDate(linksId, from, to)
	if err != nil {
		return 0, err
	}
	return stat, nil
}
