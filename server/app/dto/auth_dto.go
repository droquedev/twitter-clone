package dto

type AuthLoginDTO struct {
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}
