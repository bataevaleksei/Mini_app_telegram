package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	_ = godotenv.Load()

	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok || dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Простой healthcheck
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ctx := context.Background()

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// Запуск на :8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
