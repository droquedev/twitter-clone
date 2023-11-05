package repository

import (
	"database/sql"
	"fmt"
	"twitter-clone/server/app/entity"
	"twitter-clone/server/app/shared"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) entity.UserRepository {
	return &UserPostgresRepository{db}
}

func (u *UserPostgresRepository) FindUserByUsername(username string) (*entity.User, error) {
	row := u.db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	user := new(entity.User)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserPostgresRepository) Search(filters []shared.Filters) ([]*entity.User, error) {

	query := "SELECT * FROM users WHERE "
	args := make([]interface{}, len(filters))

	for i, filter := range filters {
		if i > 0 {
			query += " AND "
		}

		query += fmt.Sprintf("%s %s $%d", filter.Field, filter.Operator, i+1)
		args[i] = filter.Value
	}

	sql, err := u.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err := sql.Query(args...)

	if err != nil {
		return make([]*entity.User, 0), err
	}

	var users []*entity.User

	for rows.Next() {
		user := new(entity.User)

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserPostgresRepository) Persist(user *entity.User) error {
	_, err := u.db.Exec(
		"INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4)",
		user.ID,
		user.Username,
		user.Email,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}
