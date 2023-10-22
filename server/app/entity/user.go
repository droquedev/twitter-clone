package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string
}

type UserRepository interface {
	FindUserByUsername(username string) (*User, error)
}
