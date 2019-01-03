package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/leepuppychow/jay_medtronic/env"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	fmt.Println("Database connection initialized")
	DB, _ = sql.Open("postgres", env.DevelopmentDB)
	err := DB.Ping()
	if err != nil {
		log.Fatal("Cannot establish a connection with the database", err)
	}
}
