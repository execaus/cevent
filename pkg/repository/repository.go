package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Event interface {
}

type Repository struct {
	Authorization
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
