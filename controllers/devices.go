package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func DevicesIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllDevices()
	WriteResponse(data, err, 400, w)
}

func DeviceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindDevice(id)
	WriteResponse(data, err, 400, w)
}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateDevice(r.Body)
	WriteResponse(data, err, 422, w)
}

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateDevice(id, r.Body)
	WriteResponse(data, err, 422, w)
}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteDevice(id)
	WriteResponse(data, err, 400, w)
}
