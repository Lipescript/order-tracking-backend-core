package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"users-api/app/src/config"
	"users-api/app/src/constants"

)

type MongodbRepo struct {
	MongodbClient *mongo.Client
}

func CreateNewMongoDbRepo(a *config.MongoConfig) *MongodbRepo {

	client, _ := initDatabaseConection()
	a.Mongo = client

	return &MongodbRepo{
		MongodbClient: a.Mongo,
	}
}

func InsertData(client *mongo.Client, databaseName, collectionName string, document any) (*mongo.InsertOneResult, error) {
	bsonDocument, err := convertToBson(document)

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

func initDatabaseConection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(constants.DatabaseURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := pingDatabase(client); err != nil {
		return nil, err
	}
	log.Println("Connected to the database ", client.Database(constants.DatabaseName).Name())
	return client, nil
}

func pingDatabase(client *mongo.Client) error {
	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}
	return nil
}
