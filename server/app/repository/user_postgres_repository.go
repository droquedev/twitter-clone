package repository

import (
	"database/sql"
	"twitter-clone/server/app/entity"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func (u *UserPostgresRepository) FindUserByUsername(username string) (*entity.User, error) {
	row := u.db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	user := new(entity.User)

	err := row.Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserPostgresRepository(db *sql.DB) entity.UserRepository {
	return &UserPostgresRepository{db}
}
