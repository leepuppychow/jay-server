package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leepuppychow/jay_medtronic/database"
	_ "github.com/leepuppychow/jay_medtronic/env"
	"github.com/leepuppychow/jay_medtronic/routes"

	"github.com/gorilla/handlers"
)

func main() {
	router := routes.NewRouter()
	fmt.Println("Server running on port 8000")
	database.InitDB("production")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
