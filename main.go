package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/leepuppychow/jay_medtronic/database"
	_ "github.com/leepuppychow/jay_medtronic/env"
	"github.com/leepuppychow/jay_medtronic/routes"

	"github.com/gorilla/handlers"
)

func main() {
	port := 3000
	router := routes.NewRouter()
	fmt.Println("Server running on port:", port)
	database.InitDB(os.Getenv("SERVER_ENV"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
