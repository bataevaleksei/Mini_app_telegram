package inserts

import (
	"database/sql"
)

func InsertProduct(db *sql.DB, title string, description string, price float64) error {
	_, err := db.Exec(`INSERT INTO products (title, description, price) VALUES ($1, $2, $3)`, title, description, price)
	return err

}
