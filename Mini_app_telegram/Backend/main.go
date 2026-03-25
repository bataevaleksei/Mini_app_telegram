package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var sqlFiles embed.FS

func initSchema(ctx context.Context, db *pgxpool.Pool) error {
	entries, err := fs.ReadDir(sqlFiles, "sql")
	if err != nil {
		return err
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		names = append(names, e.Name())
	}
	sort.Strings(names)

	for _, name := range names {
		path := "sql/" + name

		queryBytes, err := sqlFiles.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", path, err)
		}

		if _, err := db.Exec(ctx, string(queryBytes)); err != nil {
			return fmt.Errorf("exec %s: %w", path, err)
		}

		log.Printf("applied %s", path)
	}

	return nil
}

func main() {
	r := gin.Default()

	// Простой healthcheck
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Запуск на :8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

	ctx := context.Background()
	dbURL := ""

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := initSchema(ctx, db); err != nil {
		log.Fatal(err)
	}

}
