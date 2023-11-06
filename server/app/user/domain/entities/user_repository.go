package user_entities

import "twitter-clone/server/app/shared/shared_domain"

type UserRepository interface {
	FindUserByUsername(username string) (*User, error)
	Persist(user *User) error
	Search(filters []shared_domain.Filters) ([]*User, error)
}
