package usecase

import (
	"errors"
	"twitter-clone/server/app/dto"
	"twitter-clone/server/app/entity"
	"twitter-clone/server/pkg/password"
)

type AuthUseCase struct {
	userRepository entity.UserRepository
}

func NewAuthUsecase(postRepo entity.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepository: postRepo,
	}
}

func (u *AuthUseCase) Login(createPostDTO dto.AuthLoginDTO) (*entity.User, error) {

	user, err := u.userRepository.FindUserByUsername(createPostDTO.UserName)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user or password incorrect")
	}

	if validPassword := password.Verify(user.Password, createPostDTO.Password); !validPassword {
		return nil, errors.New("user or password incorrect")
	}

	return user, nil
}
