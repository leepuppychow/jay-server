package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/leepuppychow/jay_medtronic/helpers"
	"github.com/leepuppychow/jay_medtronic/models"
)

func DevicePapersIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllDevicePapers()
	h.WriteResponse(data, err, 400, w)
}

func DevicePaperShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindDevicePaper(id)
	h.WriteResponse(data, err, 400, w)
}

func CreateDevicePaper(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateDevicePaper(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func UpdateDevicePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateDevicePaper(id, r.Body)
	h.WriteResponse(data, err, 422, w)
}

func DeleteDevicePaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteDevicePaper(id)
	h.WriteResponse(data, err, 400, w)
}
