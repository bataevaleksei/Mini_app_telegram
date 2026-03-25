package inserts

import (
	"database/sql"
	"time"
)

func InsertOrder(db *sql.DB, userID int, cart_items_id int, total_price float64) error {
	_, err := db.Exec(`INSERT INTO orders (user_id, cart_items_id, total_price, status, created_at) VALUES ($1, $2, $3, $4, $5)`, userID, cart_items_id, total_price, "pending", time.Now())
	return err
}
