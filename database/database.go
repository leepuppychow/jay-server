package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/leepuppychow/jay_medtronic/env"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(environment string) {
	var connection string
	switch environment {
	case "production":
		connection = env.ProductionDB
	case "development":
		connection = env.DevelopmentDB
	case "test":
		connection = env.TestDB
	default:
		connection = env.DevelopmentDB
	}
	fmt.Println("Database connection initialized")
	DB, _ = sql.Open("postgres", connection)
	err := DB.Ping()
	if err != nil {
		log.Fatal("Cannot establish a connection with the database", err)
	}
}
