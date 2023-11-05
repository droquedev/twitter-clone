package event

import (
	"twitter-clone/server/app/entity"
	"twitter-clone/server/pkg/event_bus"
)

func CreateUserEvent(user *entity.User) *event_bus.EventData {
	return &event_bus.EventData{
		Topic: "USER_CREATED",
		Data:  user,
	}
}
