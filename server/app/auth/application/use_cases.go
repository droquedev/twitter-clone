package auth_application

import (
	user_entities "twitter-clone/server/app/user/domain/entities"
)

type AuthUseCases struct {
	userRepository user_entities.UserRepository
}

func NewAuthLogin(userRepo user_entities.UserRepository) *AuthUseCases {
	return &AuthUseCases{
		userRepository: userRepo,
	}
}
