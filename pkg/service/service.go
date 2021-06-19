package service

import "cevent/pkg/repository"

type Authorization interface {
}

type Event interface {
}

type Service struct {
	Authorization
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
