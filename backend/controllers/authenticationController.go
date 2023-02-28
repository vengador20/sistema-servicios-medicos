package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
)

type Response struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

func (con *Controllers) Login(c *fiber.Ctx) error {
	fmt.Println("hola")

	var user models.User

	err := c.BodyParser(&user)

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//message de errores
	errors, err := config.ValidateUser(&user)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error", Errors: errors})
	}

	return c.Status(http.StatusOK).JSON(Response{Message: "exito"})
}

func (con *Controllers) ResetPassword(c *fiber.Ctx) error {
	return c.JSON("")
}
