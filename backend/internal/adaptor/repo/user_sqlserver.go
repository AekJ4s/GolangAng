package repo

import (
	"database/sql"
	"golangbackend/internal/core/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Save(user *domain.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (username, password) VALUES (@username, @password)",
		sql.Named("username", user.Username),
		sql.Named("password", user.Password),
	)
	return err
}

func (r *UserRepo) FindByUsername(username string) (*domain.User, error) {
	row := r.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username = @username",
		sql.Named("username", username),
	)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
