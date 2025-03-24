package database

import (
	"users-api/app/src/config"
)

type Repository struct {
	App *config.MongoConfig
	DB  MongodbRepo
}

var Repo *Repository

func NewRepo(app *config.MongoConfig, db *MongodbRepo) *Repository {
	return &Repository{
		App: app,
		DB:  *db,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
