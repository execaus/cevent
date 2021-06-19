package repository

type Authorization interface {
}

type Event interface {
}

type Repository struct {
	Authorization
	Event
}

func NewRepository() *Repository {
	return &Repository{}
}
