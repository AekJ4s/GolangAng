package driving

import "golangbackend/internal/core/domain"

type UserService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
	GetProfile(username string) (*domain.User, error)
}
