package inserts

import (
	"context"
	"database/sql"
)

func InsertCartItem(ctx context.Context, db *sql.DB, userID, productID, quantity int64) error {
	var cartID int64

	err := db.QueryRowContext(ctx, `
		SELECT cart_id
		FROM carts
		WHERE user_id = $1
	`, userID).Scan(&cartID)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, `
		INSERT INTO cart_items (cart_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, product_id)
		DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity
	`, cartID, productID, quantity)

	return err
}
