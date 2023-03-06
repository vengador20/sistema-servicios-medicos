package routers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/controllers"
	"github.com/vengador20/sistema-servicios-medicos/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	Db *mongo.Client //*mongo.Client
}

func (r *Router) Router(router fiber.Router) {

	controller := controllers.Controllers{
		Client: r.Db,
	}
	router.Get("/hola", func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		defer cancel()

		mongo := database.Mongodb{
			Client: r.Db,
		}

		coll, err := mongo.Collection("users") //Collection("users")

		if err != nil {
			fmt.Println(err)
		}

		cur, err := coll.Find(ctx, bson.D{})
		defer cur.Close(ctx)

		if err != nil {
			fmt.Println(err)
		}

		var users []bson.D

		err = cur.All(ctx, &users)

		if err != nil {
			c.JSON(map[string]string{"error": err.Error()})
		}

		return c.JSON(&users)
	})

	router.Post("/login", controller.Login)

	router.Post("/register", controller.RegisterUser)
}
