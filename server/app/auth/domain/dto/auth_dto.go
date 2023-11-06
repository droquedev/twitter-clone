package auth_dto

type AuthLoginDTO struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
