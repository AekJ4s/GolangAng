package repository

import (
	"database/sql"
	"golangbackend/domain/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (username, password) VALUES (?, ?)",
		user.Username, user.Password,
	)
	return err
}

func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	row := r.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username = ?",
		username,
	)
	user := &entity.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
