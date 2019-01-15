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
	// User endpoints
	{"CheckToken", "GET", "/api/v1/checktoken", controllers.CheckToken},
	{"CreateUser", "POST", "/api/v1/users", controllers.CreateUser},
	{"Login", "POST", "/api/v1/login", controllers.LoginUser},
	// Papers endpoints
	{"CreatePaper", "POST", "/api/v1/papers", controllers.CreatePaper},
	{"PapersIndex", "GET", "/api/v1/papers", controllers.PapersIndex},
	{"PapersShow", "GET", "/api/v1/papers/{id}", controllers.PapersShow},
	{"UpdatePaper", "PATCH", "/api/v1/papers/{id}", controllers.UpdatePaper},
	{"UpdatePaper", "PUT", "/api/v1/papers/{id}", controllers.UpdatePaper},
	{"DeletePaper", "DELETE", "/api/v1/papers/{id}", controllers.DeletePaper},
	// Studies endpoints
	{"StudiesIndex", "GET", "/api/v1/studies", controllers.StudiesIndex},
	{"StudyShow", "GET", "/api/v1/studies/{id}", controllers.StudyShow},
	{"CreateStudy", "POST", "/api/v1/studies", controllers.CreateStudy},
	{"UpdateStudy", "PATCH", "/api/v1/studies/{id}", controllers.UpdateStudy},
	{"UpdateStudy", "PUT", "/api/v1/studies/{id}", controllers.UpdateStudy},
	{"DeleteStudy", "DELETE", "/api/v1/studies/{id}", controllers.DeleteStudy},
	// Journals endpoints
	{"JournalsIndex", "GET", "/api/v1/journals", controllers.JournalsIndex},
	{"JournalShow", "GET", "/api/v1/journals/{id}", controllers.JournalShow},
	{"CreateJournal", "POST", "/api/v1/journals", controllers.CreateJournal},
	{"UpdateJournal", "PATCH", "/api/v1/journals/{id}", controllers.UpdateJournal},
	{"UpdateJournal", "PUT", "/api/v1/journals/{id}", controllers.UpdateJournal},
	{"DeleteJournal", "DELETE", "/api/v1/journals/{id}", controllers.DeleteJournal},
	// Authors endpoints
	{"AuthorsIndex", "GET", "/api/v1/authors", controllers.AuthorsIndex},
	{"AuthorShow", "GET", "/api/v1/authors/{id}", controllers.AuthorShow},
	{"CreateAuthor", "POST", "/api/v1/authors", controllers.CreateAuthor},
	{"UpdateAuthor", "PATCH", "/api/v1/authors/{id}", controllers.UpdateAuthor},
	{"UpdateAuthor", "PUT", "/api/v1/authors/{id}", controllers.UpdateAuthor},
	{"DeleteAuthor", "DELETE", "/api/v1/authors/{id}", controllers.DeleteAuthor},
	// Devices endpoints
	{"DevicesIndex", "GET", "/api/v1/devices", controllers.DevicesIndex},
	{"DeviceShow", "GET", "/api/v1/devices/{id}", controllers.DeviceShow},
	{"CreateDevice", "POST", "/api/v1/devices", controllers.CreateDevice},
	{"UpdateDevice", "PATCH", "/api/v1/devices/{id}", controllers.UpdateDevice},
	{"UpdateDevice", "PUT", "/api/v1/devices/{id}", controllers.UpdateDevice},
	{"DeleteDevice", "DELETE", "/api/v1/devices/{id}", controllers.DeleteDevice},
	// Figures endpoints
	{"FiguresIndex", "GET", "/api/v1/figures", controllers.FiguresIndex},
	{"FigureShow", "GET", "/api/v1/figures/{id}", controllers.FigureShow},
	{"CreateFigure", "POST", "/api/v1/figures", controllers.CreateFigure},
	{"UpdateFigure", "PATCH", "/api/v1/figures/{id}", controllers.UpdateFigure},
	{"UpdateFigure", "PUT", "/api/v1/figures/{id}", controllers.UpdateFigure},
	{"DeleteFigure", "DELETE", "/api/v1/figures/{id}", controllers.DeleteFigure},
	// Data request forms endpoints
	{"DataRequestFormsIndex", "GET", "/api/v1/data_request_forms", controllers.DataRequestFormsIndex},
	{"DataRequestFormShow", "GET", "/api/v1/data_request_forms/{id}", controllers.DataRequestFormShow},
	{"CreateDataRequestForm", "POST", "/api/v1/data_request_forms", controllers.CreateDataRequestForm},
	{"UpdateDataRequestForm", "PATCH", "/api/v1/data_request_forms/{id}", controllers.UpdateDataRequestForm},
	{"UpdateDataRequestForm", "PUT", "/api/v1/data_request_forms/{id}", controllers.UpdateDataRequestForm},
	{"DeleteDataRequestForm", "DELETE", "/api/v1/data_request_forms/{id}", controllers.DeleteDataRequestForm},
	// Figure_papers endpoints
	{"FigurePapersIndex", "GET", "/api/v1/figure_papers", controllers.FigurePapersIndex},
	{"FigurePaperShow", "GET", "/api/v1/figure_papers/{id}", controllers.FigurePaperShow},
	{"CreateFigurePaper", "POST", "/api/v1/figure_papers", controllers.CreateFigurePaper},
	{"UpdateFigurePaper", "PATCH", "/api/v1/figure_papers/{id}", controllers.UpdateFigurePaper},
	{"UpdateFigurePaper", "PUT", "/api/v1/figure_papers/{id}", controllers.UpdateFigurePaper},
	{"DeleteFigurePaper", "DELETE", "/api/v1/figure_papers/{id}", controllers.DeleteFigurePaper},
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
