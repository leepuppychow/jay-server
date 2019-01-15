package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func DataRequestFormsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllDataRequestForms(r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func DataRequestFormShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindDataRequestForm(id, r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func CreateDataRequestForm(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateDataRequestForm(r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

func UpdateDataRequestForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateDataRequestForm(id, r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

func DeleteDataRequestForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteDataRequestForm(id, r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}
