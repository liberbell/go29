package models

import "database/sql"

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
