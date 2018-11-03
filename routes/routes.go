package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Url         string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{"CreateUser", "POST", "/api/v1/users", createUser},
	{"Login", "POST", "/api/v1/login", loginUser},
}

func ToJSON(arg interface{}) []byte {
	json, err := json.MarshalIndent(arg, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create USER PATH")
	w.Write(ToJSON("create user path"))
}

func loginUser(w http.ResponseWriter, r *http.Request) {

}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Url).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
