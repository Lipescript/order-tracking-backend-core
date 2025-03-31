package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"users-api/app/src/constants"
)

type MongoConfig struct {
	Mongo        *mongo.Client
	InProduction bool
}

func NewMongoConfig() *MongoConfig {
	return &MongoConfig{}
}

func InitMongoConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	clientURI := options.Client().ApplyURI(constants.DatabaseURI)
	client, err := mongo.Connect(ctx, clientURI)
	defer cancel()

	if err != nil {
		return nil, err
	}

	if err := pingDatabase(client); err != nil {
		return nil, err
	}

	log.Printf(constants.MESSAGE_SUCCESSFUL_CONNECTION_TO_DATABASE, constants.DatabaseName)
	return client, nil
}

func CloseMongoConnection(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func pingDatabase(client *mongo.Client) error {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf(constants.MESSAGE_FAILED_CONNECTION_TO_DATABASE, err)
		return err
	}
	return nil
}
