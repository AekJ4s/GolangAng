package repository

import "golangbackend/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
}
