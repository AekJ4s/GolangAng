package usecase

import (
	"errors"
	"time"

	"golangbackend/domain/entity"
	domainRepo "golangbackend/domain/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("it02-secret-key-2024")

type userUseCase struct {
	userRepo domainRepo.UserRepository
}

func NewUserUseCase(repo domainRepo.UserRepository) *userUseCase {
	return &userUseCase{userRepo: repo}
}

func (u *userUseCase) Register(username, password string) error {
	existing, _ := u.userRepo.FindByUsername(username)
	if existing != nil {
		return errors.New("username already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username: username,
		Password: string(hashed),
	}
	return u.userRepo.Create(user)
}

func (u *userUseCase) Login(username, password string) (string, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (u *userUseCase) GetProfile(username string) (*entity.User, error) {
	return u.userRepo.FindByUsername(username)
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
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
