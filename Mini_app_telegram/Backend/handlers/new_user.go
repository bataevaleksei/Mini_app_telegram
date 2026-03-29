package handlers

import (
	"database/sql"

	"telega_mini_app/migrations/inserts"
)

func NewUserHander(db *sql.DB, id int) error {
	return inserts.InsertUser(db, id)
}
