package service

import (
	"errors"
	"time"

	"golangbackend/infrastructure"
	"golangbackend/internal/core/domain"
	"golangbackend/internal/core/port/driven"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo driven.UserRepository
}

func NewUserService(repo driven.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(username, password string) error {
	existing, _ := s.repo.FindByUsername(username)
	if existing != nil {
		return errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.Save(&domain.User{
		Username: username,
		Password: string(hashed),
	})
}

func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(infrastructure.AppCfg.JWT.Secret))
}

func (s *UserService) GetProfile(username string) (*domain.User, error) {
	return s.repo.FindByUsername(username)
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(infrastructure.AppCfg.JWT.Secret), nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("username not found in token")
	}
	return username, nil
}
