package event

import "link-shortener/pkg/bus"

const ClickEventName bus.EventName = "STAT_CLICK_EVENT"

func NewStatClickEvent(linkId uint) *bus.Event {
	return &bus.Event{
		Name: ClickEventName,
		Data: linkId,
	}
}
