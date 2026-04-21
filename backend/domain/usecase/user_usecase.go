package usecase

import "golangbackend/domain/entity"

type UserUseCase interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
	GetProfile(username string) (*entity.User, error)
}
