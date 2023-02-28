package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vengador20/sistema-servicios-medicos/database"
	"github.com/vengador20/sistema-servicios-medicos/routers"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(recover.New())

	// env, err := config.GetEnviroment()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //insertamos la variable de entorno
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	// db, err := database.NewMongodb()

	// defer db.Disconnect(ctx)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// route := routers.Router{
	// 	Db: db,
	// }
	db, err := database.Connect() //NewMongodb()

	if err != nil {
		fmt.Println(err)
	}

	defer db.Disconnect(ctx)

	app.Get("/hola", func(c *fiber.Ctx) error {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		defer cancel()

		//db := database.DbConnect()

		// if err != nil {
		// 	fmt.Println(err)
		// }

		mongo := database.Mongodb{
			Client: db,
		}

		coll, err := mongo.Collection("users") //db.Collection("users")

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

	rt := routers.Router{
		Db: db,
	}

	app.Route("/api", rt.Router)

	app.Listen(":3000")
}
