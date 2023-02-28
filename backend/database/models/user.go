package models

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLogin struct {
	Name     string `json:"name" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
