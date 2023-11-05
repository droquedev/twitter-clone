package event_bus

import (
	"github.com/asaskevich/EventBus"
)

var bus EventBus.Bus

type EventData struct {
	Topic string
	Data  interface{}
}

func init() {
	bus = EventBus.New()
}

func Publish(event *EventData) {
	bus.Publish(event.Topic, event.Data)
}

func Subscribe(topic string, fn interface{}) error {
	return bus.Subscribe(topic, fn)
}
