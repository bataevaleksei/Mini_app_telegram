package handlers

import (
	"database/sql"
	"telega_mini_app/migrations/inserts"
)

func NewProductHandler(db *sql.DB, name string, description string, price float64) error {
	return inserts.InsertProduct(db, name, description, price)
}
