package service

import (
	"cevent/models"
	"cevent/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) error
}

type Event interface {
}

type Service struct {
	Authorization
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
