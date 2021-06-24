package models

import "time"

type User struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	Photo      string    `json:"photo"`
	Company    string    `json:"company"`
	CreateDate time.Time `json:"createDate"`
	Birthday   time.Time `json:"birthday"`
}
