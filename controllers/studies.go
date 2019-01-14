package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func StudiesIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllStudies(r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func StudyShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindStudy(id, r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func CreateStudy(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateStudy(r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

func UpdateStudy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateStudy(id, r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

// func DeleteStudy(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])
// 	data, err := models.DeleteStudy(id, r.Header.Get("Authorization"))
// 	WriteResponse(data, err, 400, w)
// }
