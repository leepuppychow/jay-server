package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay-server/database"
	"github.com/leepuppychow/jay-server/routes"
)

var router *mux.Router

func TestMain(m *testing.M) {
	dbConn := "user=leechow dbname=jay_test sslmode=disable host=localhost port=5432 timezone=utc"
	router = routes.NewRouter(false) // opt out of using Auth middleware
	database.InitDB(dbConn)
	code := m.Run()
	os.Exit(code)
}

func ExecuteQuery(query string) {
	_, err := database.DB.Exec(query)
	if err != nil {
		log.Fatal("Unable to execute test query", err)
	}
}

func TestPapersIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/papers", nil)
	respRecorder := httptest.NewRecorder()
	router.ServeHTTP(respRecorder, req)

	if respRecorder.Code != 200 {
		t.Errorf("GET to papers index failed")
	}
	fmt.Println(respRecorder.Body)
}

func TestPapersShow(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/papers/3", nil)
	respRecorder := httptest.NewRecorder()
	router.ServeHTTP(respRecorder, req)

	if respRecorder.Code != 200 {
		t.Errorf("GET to papers show failed")
	}

	// Non-existent paper
	req, _ = http.NewRequest("GET", "/api/v1/papers/1000", nil)
	respRecorder = httptest.NewRecorder()
	router.ServeHTTP(respRecorder, req)

	if respRecorder.Code != 400 {
		t.Errorf("GET to papers show failed")
	}
}
