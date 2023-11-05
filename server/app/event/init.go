package event

import (
	"twitter-clone/server/app/entity"
	"twitter-clone/server/pkg/event_bus"
)

func InitSubscriptions() {
	event_bus.Subscribe("USER_CREATED", func(user *entity.User) {
		println("Should send email to: ", user.Email)
	})
}
