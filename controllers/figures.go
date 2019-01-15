package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func FiguresIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllFigures(r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func FigureShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindFigure(id, r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func CreateFigure(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateFigure(r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

func UpdateFigure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateFigure(id, r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

func DeleteFigure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteFigure(id, r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}
