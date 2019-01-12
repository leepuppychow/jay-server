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
	{"CheckToken", "GET", "/api/v1/checktoken", controllers.CheckToken},
	{"CreateUser", "POST", "/api/v1/users", controllers.CreateUser},
	{"Login", "POST", "/api/v1/login", controllers.LoginUser},
	{"PapersIndex", "GET", "/api/v1/papers", controllers.PapersIndex},
	{"CreatePaper", "POST", "/api/v1/papers", controllers.CreatePaper},
	{"UpdatePaper", "PATCH", "/api/v1/papers/{id}", controllers.UpdatePaper},
	{"UpdatePaper", "PUT", "/api/v1/papers/{id}", controllers.UpdatePaper},
	{"DeletePaper", "DELETE", "/api/v1/papers/{id}", controllers.DeletePaper},
	{"AuthorsIndex", "POST", "/api/v1/authors", controllers.AuthorsIndex},
	{"CreateAuthor", "POST", "/api/v1/authors", controllers.CreateAuthor},
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
