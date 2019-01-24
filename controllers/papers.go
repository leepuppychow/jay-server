package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func PapersIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllPapers()
	WriteResponse(data, err, 400, w)
}

func PapersShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindPaper(id)
	WriteResponse(data, err, 400, w)
}

func CreatePaper(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreatePaper(r.Body)
	WriteResponse(data, err, 422, w)
}

func SpecialCreatePaper(w http.ResponseWriter, r *http.Request) {
	data, err := models.SpecialCreatePaper(r.Body)
	WriteResponse(data, err, 422, w)
}

func UpdatePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdatePaper(id, r.Body)
	WriteResponse(data, err, 422, w)
}

func DeletePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeletePaper(id)
	WriteResponse(data, err, 400, w)
}
