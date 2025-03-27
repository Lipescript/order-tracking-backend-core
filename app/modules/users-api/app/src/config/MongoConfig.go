package config

import (
	"context"
	"log"

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

func InitDatabaseConection() (*mongo.Client, error) {
	clientURI := options.Client().ApplyURI(constants.DatabaseURI)
	client, err := mongo.Connect(context.Background(), clientURI)
	if err != nil {
		client.Disconnect(context.Background())
		return nil, err
	}

	if err := pingDatabase(client); err != nil {
		client.Disconnect(context.Background())
		return nil, err
	}
	log.Println("Connected to database: ", constants.DatabaseName)
	return client, nil
}

func pingDatabase(client *mongo.Client) error {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v. Disconnecting...", err)
		return err
	}
	return nil
}
