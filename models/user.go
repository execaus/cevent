package models

import "time"

type User struct {
	Email      string    `json:"email" binding:"required"`
	Password   string    `json:"-" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Surname    string    `json:"surname" binding:"required"`
	Patronymic string    `json:"patronymic"`
	Photo      string    `json:"photo"`
	Company    string    `json:"company"`
	CreateDate time.Time `json:"createDate"`
	Birthday   time.Time `json:"birthday"`
}
