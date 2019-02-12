package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/leepuppychow/jay-server/database"
	"github.com/leepuppychow/jay-server/routes"

	"github.com/gorilla/handlers"
)

func main() {
	port := 3000
	router := routes.NewRouter()
	log.Println("Server running on port:", port)

	dbConn, ok := os.LookupEnv("DB_CONN")
	if !ok {
		log.Fatal("Unable to connect to database")
	}
	database.InitDB(dbConn)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
