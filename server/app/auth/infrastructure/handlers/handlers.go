package auth_handlers

import (
	auth_application "twitter-clone/server/app/auth/application"
)

type AuthHandler struct {
	authLogin *auth_application.AuthUseCases
}

func NewAuthHandler(authLogin *auth_application.AuthUseCases) *AuthHandler {
	return &AuthHandler{
		authLogin: authLogin,
	}
}
