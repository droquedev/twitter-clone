package entity

import "twitter-clone/server/app/shared"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserRepository interface {
	FindUserByUsername(username string) (*User, error)
	Persist(user *User) error
	Search(filters []shared.Filters) ([]*User, error)
}
