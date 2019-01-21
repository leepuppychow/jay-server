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
		fmt.Println("Using production DB")
		connection = env.ProductionDB
	case "development":
		fmt.Println("Using development DB")
		connection = env.DevelopmentDB
	case "test":
		fmt.Println("Using test DB")
		connection = env.TestDB
	default:
		fmt.Println("Using default development DB")
		connection = env.DevelopmentDB
	}
	fmt.Println("Database connection initialized")
	DB, _ = sql.Open("postgres", connection)
	err := DB.Ping()
	if err != nil {
		log.Fatal("Cannot establish a connection with the database", err)
	}
}
