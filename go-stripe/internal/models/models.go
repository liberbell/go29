package models

import (
	"context"
	"database/sql"
	"time"
)

type DBmodel struct {
	DB *sql.DB
}

type Models struct {
	DB DBmodel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBmodel{DB: db},
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

type Order struct {
	ID            int
	WidgetID      int
	TransactionID int
	StatusID      int
	Quantity      int
	Amount        int
}

func (m *DBmodel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget
	row := m.DB.QueryRowContext(ctx, "SELECT id, name FROM widgets WHERE id = ?", id)
	err := row.Scan(&widget.ID, &widget.Name)
	if err != nil {
		return widget, err
	}
	return widget, nil
}
