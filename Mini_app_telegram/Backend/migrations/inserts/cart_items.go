package inserts

import (
	"database/sql"
)

func insertCartItem(db *sql.DB, cartID int, productID int, quantityID int) error {
	_, err := db.Exec(`INSERT INTO cart_items (cart_id, product_id, quantity_id) VALUES ($1, $2, $3)`, cartID, productID, quantityID)
	return err
}
