package database

import (
	"context"
	"fmt"
	"time"

	"github.com/vengador20/sistema-servicios-medicos/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var (
// 	once sync.Once
// 	mg   *Mongodb
// )

type Mongodb struct {
	Client *mongo.Client
	//url string
}

// nueva conexión a Mongodb
func Connect() (*mongo.Client, error) {
	// env, err := config.GetEnviroment()

	// once.Do(func() {
	// 	mg = &Mongodb{
	// 		url: env,
	// 	}
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// return mg, nil

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	env, err := config.GetEnviroment()
	if err != nil {
		fmt.Println(env)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env))

	//retorna un error
	if err != nil {
		return nil, err
	}

	//retorna la conexión
	return client, nil

}

// func (m *Mongodb) getDatabase() (*mongo.Client, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.url))

// 	//retorna un error
// 	if err != nil {
// 		return nil, err
// 	}

// 	//retorna la conexión
// 	return client, nil
// }

func (m *Mongodb) Collection(name string) (*mongo.Collection, error) {
	//sistema-medico
	coll := m.Client.Database("uguia").Collection(name)

	//retorna la colección
	return coll, nil
}

func (m *Mongodb) Disconnect(ctx context.Context) {
	//db, _ := m.getDatabase()

	m.Client.Disconnect(ctx)
}
