package models

// type User struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required,min=8"`
// }

type UserLogin struct {
	Nombres  string `json:"nombres" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserRegister struct {
	Nombres         string `json:"nombres" bson:"nombres,omitempty" validate:"required"`
	FechaNacimiento string `json:"fecha" bson:"fecha,omitempty" validate:"required"`
	Telefono        uint   `json:"telefono" bson:"telefono,omitempty" validate:"required"` //numeros enteros
	Password        string `json:"password" bson:"password,omitempty" validate:"required"`
	Tipo            int    `json:"tipo" bson:"tipo,omitempty" validate:"required"` // 0 es un paciente y 1 es servicio
}

type UserNewPassword struct {
	Nombres  string `json:"nombres" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}
