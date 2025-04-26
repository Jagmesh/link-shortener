package bus

import (
	"link-shortener/pkg/logger"
)

type EventName string
type Event struct {
	Name EventName
	Data any
}
type EventBus struct {
	busMap map[EventName]chan *Event
}

var log = logger.GetWithScopes("EVENT_BUS")

func New() *EventBus {
	return &EventBus{
		busMap: make(map[EventName]chan *Event),
	}
}

func (eb *EventBus) Publish(event *Event) {
	log.Debugf("Event received: {%s}", event.Name)
	if eb.busMap[event.Name] == nil {
		eb.busMap[event.Name] = make(chan *Event)
	}
	eb.busMap[event.Name] <- event
}

func (eb *EventBus) Consume(eventName EventName) chan *Event {
	log.Debugf("Event consuming: {%s}", eventName)
	if eb.busMap[eventName] == nil {
		eb.busMap[eventName] = make(chan *Event)
	}
	return eb.busMap[eventName]
}
