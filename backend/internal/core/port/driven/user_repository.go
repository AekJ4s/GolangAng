package driven

import "golangbackend/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByUsername(username string) (*domain.User, error)
}
