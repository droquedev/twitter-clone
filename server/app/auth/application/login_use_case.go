package auth_application

import (
	"errors"
	auth_dto "twitter-clone/server/app/auth/domain/dto"
	user_entities "twitter-clone/server/app/user/domain/entities"
	"twitter-clone/server/pkg/password"
)

func (u *AuthUseCases) Login(createPostDTO auth_dto.AuthLoginDTO) (*user_entities.User, error) {

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
