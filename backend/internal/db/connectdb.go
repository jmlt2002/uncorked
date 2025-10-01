package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	host := "localhost"
	port := "5432"
	user := "user"
	password := "user"
	dbname := "uncorked"

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, nil
}
