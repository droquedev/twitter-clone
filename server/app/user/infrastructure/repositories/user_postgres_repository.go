package user_repositories

import (
	"database/sql"
	"twitter-clone/server/app/shared/shared_domain"
	user_entities "twitter-clone/server/app/user/domain/entities"
)

type UserPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) user_entities.UserRepository {
	return &UserPostgresRepository{db}
}

func (u *UserPostgresRepository) FindUserByUsername(username string) (*user_entities.User, error) {
	row := u.db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	user := new(user_entities.User)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserPostgresRepository) Search(filters []shared_domain.Filters) ([]*user_entities.User, error) {

	query := "SELECT * FROM users WHERE "

	filtersString, args := shared_domain.GenerateFilter(filters)

	query += filtersString

	sql, err := u.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err := sql.Query(args...)

	if err != nil {
		return make([]*user_entities.User, 0), err
	}

	var users []*user_entities.User

	for rows.Next() {
		user := new(user_entities.User)

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserPostgresRepository) Persist(user *user_entities.User) error {
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
