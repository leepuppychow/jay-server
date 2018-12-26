package main

import (
	"fmt"
	_ "github.com/leepuppychow/jay_medtronic/database"
	"github.com/leepuppychow/jay_medtronic/env"
	"github.com/leepuppychow/jay_medtronic/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	router := routes.NewRouter()
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{env.DevelopmentDomain}))(router)))
}
