package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database"
	"github.com/vengador20/sistema-servicios-medicos/routers"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(recover.New())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()
	db, err := database.Connect()

	if err != nil {
		fmt.Println(err)
	}

	defer db.Disconnect(ctx)

	rt := routers.Router{
		Db: db,
	}

	app.Get("/cookie", func(c *fiber.Ctx) error {

		token, err := config.NewToken("efrain@gmail.com")

		if err != nil {
			return c.JSON("error")
		}

		cookie := fiber.Cookie{
			// Name:     "Authorized",
			// Value:    token,
			// Path:     "/",
			// SameSite: "Strict",
			// Secure:   true,
			// HTTPOnly: true,
			// Expires:  time.Now().Add(time.Hour * 200),
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   true,
			Path:     "/",
			//cookie.Expires = time.Now().Add(24 * time.Hour)
			SameSite: "none",
		}

		c.Cookie(&cookie)

		return c.JSON("cookie")
	})

	//utilizar middleware personalizado
	//valida si el jwt no es modificado
	//app.Use("/api", middleware.ValidateJwt)

	app.Route("/api", rt.Router)

	app.Listen(":3000")
}
