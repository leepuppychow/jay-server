package main

import (
	"fmt"
	_ "jay_medtronic/database"
	"jay_medtronic/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
