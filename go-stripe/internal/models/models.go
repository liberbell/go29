package models

import "database/sql"

type DBmodels struct {
	DB *sql.DB
}

type Models struct {
	DB DBmodels
}
