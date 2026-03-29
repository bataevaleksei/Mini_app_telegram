package handlers

import (
	"context"
	"database/sql"
	"telega_mini_app/migrations/inserts"
)

func NewCartItemHandler(ctx context.Context, db *sql.DB, userID, productID, quantity int64) error {

	return inserts.InsertCartItem(ctx, db, userID, productID, quantity)
}
