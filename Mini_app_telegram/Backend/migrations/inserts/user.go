package inserts

import (
	"database/sql"
)

func InsertUser(db *sql.DB, id int) error {
	_, err := db.Exec(`INSERT INTO users (id) VALUES ($1)`, id)
	return err

}
