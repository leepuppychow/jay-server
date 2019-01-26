package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/leepuppychow/jay_medtronic/helpers"
	"github.com/leepuppychow/jay_medtronic/models"
)

func FigurePapersIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllFigurePapers()
	h.WriteResponse(data, err, 400, w)
}

func FigurePaperShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindFigurePaper(id)
	h.WriteResponse(data, err, 400, w)
}

func CreateFigurePaper(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateFigurePaper(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func UpdateFigurePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateFigurePaper(id, r.Body)
	h.WriteResponse(data, err, 422, w)
}

func DeleteFigurePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteFigurePaper(id)
	h.WriteResponse(data, err, 400, w)
}
