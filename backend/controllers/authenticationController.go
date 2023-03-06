package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

const TABLEUSER = "users"

func (con *Controllers) Login(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var user models.UserLogin

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//message de errores
	errors, err := config.ValidateUser(&user)

	//los mensajes de error cuando los campos son requeridos etc...
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error", Errors: errors})
	}

	db := database.Mongodb{
		Client: con.Client,
	}

	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error"})
	}

	var userRes models.UserLogin

	//filtrar para la busqueda del usuario si existe
	filter := bson.D{{Key: "nombres", Value: user.Nombres}}

	err = coll.FindOne(ctx, filter).Decode(&userRes)

	//usuario no existe
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "Nombre o contraseña inválido"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRes.Password), []byte(user.Password))

	//contraseña invalida
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "Nombre o contraseña inválido"})
	}

	return c.Status(http.StatusOK).JSON(Response{Message: "exito"})
}

func (con *Controllers) RegisterUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var user models.UserRegister

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//message de errores
	errors, err := config.ValidateUser(&user)

	//los mensajes de error cuando los campos son requeridos etc...
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error", Errors: errors})
	}

	db := database.Mongodb{
		Client: con.Client,
	}

	//verificar que el usuario no existe
	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var userRes bson.M

	filter := bson.M{"nombres": user.Nombres}
	err = coll.FindOne(ctx, filter).Decode(&userRes)

	//si es nulo  usuario ya existe
	if err == nil {
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "Usuario ya existe"})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	// asignamos el password anterior
	// al password encriptado
	user.Password = string(password)

	_, err = coll.InsertOne(ctx, user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.Status(http.StatusOK).JSON(Response{Message: "Usuario creado"})
}

func (con *Controllers) ResetPassword(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var user models.UserNewPassword

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	db := database.Mongodb{
		Client: con.Client,
	}

	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var userRes bson.M

	filter := bson.M{"nombres": user.Nombres}
	err = coll.FindOne(ctx, filter).Decode(&userRes)

	//usuario no existe
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "Usuario no existe"})
	}

	filterPassword := bson.M{"nombres": user.Nombres}

	update := bson.M{"password": user.Password}

	_, err = coll.UpdateOne(ctx, filterPassword, update)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//enviar correo
	//smtp.PlainAuth("","vengadorba6@gmail.com","tobimoto2000","")

	return c.Status(http.StatusOK).JSON(Response{Message: "Contraseña actualizada"})
}
