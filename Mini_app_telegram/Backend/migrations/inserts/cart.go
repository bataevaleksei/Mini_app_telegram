package inserts

import (
	"database/sql"
)

func InsertCart(db *sql.DB, userID int) error {
	_, err := db.Exec(`INSERT INTO carts (user_id) VALUES (1)`)
	return err
}
