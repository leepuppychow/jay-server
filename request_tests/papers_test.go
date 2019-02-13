package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay-server/database"
	"github.com/leepuppychow/jay-server/models"
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

func RunTestRequest(verb, route string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(verb, route, body)
	respRecorder := httptest.NewRecorder()
	router.ServeHTTP(respRecorder, req)
	return respRecorder
}

func TestPapersIndex(t *testing.T) {
	respRecorder := RunTestRequest("GET", "/api/v1/papers", nil)
	var papers []models.Paper
	json.Unmarshal(respRecorder.Body.Bytes(), &papers)

	if respRecorder.Code != 200 {
		t.Errorf("GET to papers index failed")
	}
	if len(papers) != 3 {
		t.Errorf("GET to papers index failed")
	}
}

func TestPapersShow(t *testing.T) {
	respRecorder := RunTestRequest("GET", "/api/v1/papers/3", nil)
	var paper models.Paper
	json.Unmarshal(respRecorder.Body.Bytes(), &paper)

	if respRecorder.Code != 200 || paper.Id != 3 {
		t.Errorf("GET to papers show failed")
	}

	if reflect.TypeOf(paper.Title).Kind() != reflect.String {
		t.Errorf("GET to papers show failed")
	}

	// Non-existent paper
	respRecorder = RunTestRequest("GET", "/api/v1/papers/1000", nil)

	if respRecorder.Code != 400 {
		t.Errorf("GET to papers show failed")
	}
}
