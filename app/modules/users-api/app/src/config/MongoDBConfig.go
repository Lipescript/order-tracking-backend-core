package config

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConfig struct {
	Mongo        *mongo.Client
	InProduction bool
}
