package user_application

import (
	"errors"
	"twitter-clone/server/app/shared/shared_domain"
	user_entities "twitter-clone/server/app/user/domain/entities"
	user_events "twitter-clone/server/app/user/domain/events"
	"twitter-clone/server/pkg/event_bus"
)

func (u *UserUseCases) UserCreator(user *user_entities.User) error {

	filters := []shared_domain.Filters{
		{Field: "username", Operator: "=", Value: user.Username, Logic: "OR"},
		{Field: "email", Operator: "=", Value: user.Email},
	}

	users, err := u.userRepository.Search(filters)

	if err != nil {
		println(err.Error())
		return err
	}

	if len(users) > 0 {
		return errors.New("username or email already exists")
	}

	err2 := u.userRepository.Persist(user)

	if err2 != nil {
		return err2
	}

	event_bus.Publish(user_events.UserCreatedEvent(user))

	return nil
}
