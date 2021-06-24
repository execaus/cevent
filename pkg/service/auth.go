package service

import (
	"cevent/models"
	"cevent/pkg/repository"
	"crypto/sha1"
	"fmt"
	"time"
)

const salt = "uhsdfginufgnysdfhi23lrhm8o4gn8ye"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) error {
	user.Password = generatePasswordHash(user.Password)
	user.Photo = "default-user-avatar.svg"
	user.CreateDate = time.Now()

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
