package database

import (
	"users-api/app/src/config"
	"users-api/app/src/database"
)

type Repository struct {
	App *config.MongoConfig
	DB  database.MongodbRepo
}

var Repo *Repository

func NewRepo(app *config.MongoConfig, db *database.MongodbRepo) *Repository {
	return &Repository{
		App: app,
		DB:  *db,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
