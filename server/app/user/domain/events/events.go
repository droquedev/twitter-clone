package user_events

import (
	user_domain_entities "twitter-clone/server/app/user/domain/entities"
	"twitter-clone/server/pkg/event_bus"
)

func UserCreatedEvent(user *user_domain_entities.User) *event_bus.EventData {
	return &event_bus.EventData{
		Topic: "USER_CREATED",
		Data:  user,
	}
}
