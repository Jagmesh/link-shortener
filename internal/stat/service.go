package stat

import (
	"errors"
	"link-shortener/entity/event"
	"link-shortener/pkg/bus"
)

type Service struct {
	deps *ServiceDeps
}

type ServiceDeps struct {
	Repository *Repository
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{deps: deps}
}

func (s Service) AddClick(linkId uint) error {
	return s.deps.Repository.CreateOrIncrementClick(linkId)
}

func (s Service) ListenClick(eventBus *bus.EventBus) error {
	for ev := range eventBus.Consume(event.ClickEventName) {
		id, ok := ev.Data.(uint)
		if !ok {
			return errors.New("event Data is not uint")
		}
		if s.deps.Repository.CreateOrIncrementClick(id) != nil {
			return errors.New("failed to add click")
		}
	}
	return nil
}
