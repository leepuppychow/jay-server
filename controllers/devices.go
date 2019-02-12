package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/leepuppychow/jay-server/helpers"
	"github.com/leepuppychow/jay-server/models"
)

func DevicesIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllDevices()
	h.WriteResponse(data, err, 400, w)
}

func DeviceShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindDevice(id)
	h.WriteResponse(data, err, 400, w)
}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateDevice(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateDevice(id, r.Body)
	h.WriteResponse(data, err, 422, w)
}

func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteDevice(id)
	h.WriteResponse(data, err, 400, w)
}
