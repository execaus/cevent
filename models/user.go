package models

type User struct {
	Email      string `json:"email"`
	Password   string `json:"-"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Photo      string `json:"photo"`
}
