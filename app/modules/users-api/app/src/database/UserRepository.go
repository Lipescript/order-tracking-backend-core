package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"users-api/app/src/config"
)

type MongodbRepo struct {
	MongodbClient *mongo.Client
}

func CreateNewMongoDbRepo(mongoconfig *config.MongoConfig) *MongodbRepo {

	client, _ := config.InitMongoConnection()
	mongoconfig.Mongo = client

	return &MongodbRepo{
		MongodbClient: mongoconfig.Mongo,
	}
}

// TODO implement in order-tracking-orchestror
func InsertData(client *mongo.Client, databaseName, collectionName string, document any) (*mongo.InsertOneResult, error) {
	bsonDocument, err := convertToBson(document)
	if err != nil {
		return nil, err
	}

	collection := client.Database(databaseName).Collection(collectionName)
	result, err := collection.InsertOne(context.Background(), bsonDocument)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func convertToBson(document any) (any, error) {
	bsonDocument, err := bson.Marshal(document)
	if err != nil {
		return nil, err
	}
	return bsonDocument, nil
}
