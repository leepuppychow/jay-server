package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/leepuppychow/jay_medtronic/helpers"
	"github.com/leepuppychow/jay_medtronic/models"
)

func PapersIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllPapers()
	h.WriteResponse(data, err, 400, w)
}

func PapersShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindPaper(id)
	h.WriteResponse(data, err, 400, w)
}

func CreatePaper(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreatePaper(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func SpecialCreatePaper(w http.ResponseWriter, r *http.Request) {
	data, err := models.SpecialCreatePaper(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func UpdatePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdatePaper(id, r.Body)
	h.WriteResponse(data, err, 422, w)
}

func DeletePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeletePaper(id)
	h.WriteResponse(data, err, 400, w)
}
