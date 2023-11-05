package usecase

import (
	"errors"
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/entity"
	"twitter-clone/server/app/event"
	"twitter-clone/server/app/shared"
	"twitter-clone/server/pkg/event_bus"

	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepository entity.UserRepository
}

func NewUserUsecase(userRepo entity.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (u *UserUseCase) Create(createUserDTO dto.UserCreateDTO) (*entity.User, error) {

	user := &entity.User{
		ID:       uuid.NewString(),
		Username: createUserDTO.Username,
		Email:    createUserDTO.Email,
		Password: createUserDTO.Password,
	}

	filters := []shared.Filters{
		{Field: "username", Operator: "=", Value: createUserDTO.Username},
		{Field: "email", Operator: "=", Value: createUserDTO.Email},
	}

	users, err := u.userRepository.Search(filters)

	if err != nil {
		println(err.Error())
		return nil, err
	}

	if len(users) > 0 {
		return nil, errors.New("username or email already exists")
	}

	err2 := u.userRepository.Persist(user)

	if err2 != nil {
		return nil, err2
	}

	event_bus.Publish(event.CreateUserEvent(user))

	return user, nil
}
