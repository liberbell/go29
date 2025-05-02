package models

import (
	"database/sql"
)

type DBmodels struct {
	DB *sql.DB
}

type Models struct {
	DB DBmodels
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBmodels{DB: db},
	}
}

type Widget struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
