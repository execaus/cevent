package models

import "time"

type User struct {
	ID         int       `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"passwordHash"`
	Name       string    `json:"name" db:"name"`
	Surname    string    `json:"surname" db:"surname"`
	Patronymic string    `json:"patronymic" db:"patronymic"`
	Photo      string    `json:"photo" db:"photo"`
	Company    string    `json:"company" db:"company"`
	CreateDate time.Time `json:"createDate" db:"createDate"`
	Birthday   time.Time `json:"birthday" db:"birthday"`
}
