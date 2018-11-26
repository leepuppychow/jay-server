package routes

import (
	"jay_medtronic/controllers"
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
	{"CreateUser", "POST", "/api/v1/users", users.CreateUser},
	{"Login", "POST", "/api/v1/login", users.LoginUser},

	{"test", "GET", "/api/v1/test", users.Test},
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
