package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(conn string) {
	fmt.Println("Database connection initialized", conn)
	DB, _ = sql.Open("postgres", conn)
	err := DB.Ping()
	if err != nil {
		log.Fatal("Cannot establish a connection with the database", err)
	}
}
