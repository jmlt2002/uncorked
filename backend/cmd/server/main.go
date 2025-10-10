package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmlt2002/uncorked/backend/internal/api"
	_ "github.com/lib/pq"
)

func main() {
	// later fix to use env vars
	dsn := "host=localhost port=5432 user=s password=postgres dbname=uncorked sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := api.NewRouter(db)
	fmt.Println("Server started at :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
