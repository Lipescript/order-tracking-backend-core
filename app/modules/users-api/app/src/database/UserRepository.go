package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"users-api/app/src/config"
	"users-api/app/src/domain"

)

type MongodbRepo struct {
	MongodbClient *mongo.Client
}

func CreateNewMongoDbRepo(a *config.MongoConfig) *MongodbRepo {
	return &MongodbRepo{
		MongodbClient: a.Mongo,
	}
}

func (mr *MongodbRepo) AddUser(ctx context.Context, b domain.User) error {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")
	insertResult, err := booksCollection.InsertOne(ctx, b)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Inserted a single User: ", insertResult.InsertedID)
	return nil
}

func (mr *MongodbRepo) DisconnectDB(ctx context.Context) error {
	err := mr.MongodbClient.Disconnect(ctx)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Connection to MongoDB has been closed.")
	return nil
}
