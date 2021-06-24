package repository

import (
	"cevent/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) error {
	query := fmt.Sprintf(`INSERT INTO %s ("Email", "PasswordHash", "Name", "Surname", "CreateDate", "Birthday") 
		VALUES($1, $2, $3, $4, $5, $6)`, userTable)
	_, err := r.db.Exec(query,
		user.Email,
		user.Password,
		user.Name,
		user.Surname,
		user.CreateDate,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
