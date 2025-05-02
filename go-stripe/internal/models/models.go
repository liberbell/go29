package models

import (
	"database/sql"
	"time"
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

// Widget type
type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
