package database

import (
	"context"
	"fmt"
	"time"

	"github.com/vengador20/sistema-servicios-medicos/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Client *mongo.Client
	//url string
}

// crea una nueva conexión a Mongodb
func Connect() (*mongo.Client, error) {
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

func (m *Mongodb) Collection(name string) (*mongo.Collection, error) {
	//sistema-medico
	coll := m.Client.Database("uguia").Collection(name)

	//retorna la colección
	return coll, nil
}

func (m *Mongodb) Disconnect(ctx context.Context) {
	m.Client.Disconnect(ctx)
}
