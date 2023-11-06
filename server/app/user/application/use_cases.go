package user_application

import (
	user_entities "twitter-clone/server/app/user/domain/entities"
)

type UserUseCases struct {
	userRepository user_entities.UserRepository
}

func NewUserUseCases(userRepo user_entities.UserRepository) *UserUseCases {
	return &UserUseCases{
		userRepository: userRepo,
	}
}
