package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/leepuppychow/jay_medtronic/database"
	"github.com/leepuppychow/jay_medtronic/routes"

	"github.com/gorilla/handlers"
)

var DevelopmentDB = "user=leechow dbname=jay_medtronic_development sslmode=disable host=localhost port=5432 timezone=utc"

func main() {
	port := 3000
	router := routes.NewRouter()
	fmt.Println("Server running on port:", port)

	dbConn, ok := os.LookupEnv("DB_CONN")
	if !ok {
		fmt.Println("Using development DB")
		dbConn = DevelopmentDB
	}
	database.InitDB(dbConn)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
