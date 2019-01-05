package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/controllers"
)

type Route struct {
	Name        string
	Method      string
	Url         string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{"CreateUser", "POST", "/api/v1/users", controllers.CreateUser},
	{"Login", "POST", "/api/v1/login", controllers.LoginUser},
	{"PapersIndex", "GET", "/api/v1/papers", controllers.PapersIndex},
	// {"CreatePaper", "POST", "/api/v1/papers", controllers.CreatePaper},
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
