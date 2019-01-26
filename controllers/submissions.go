package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/leepuppychow/jay_medtronic/helpers"
	"github.com/leepuppychow/jay_medtronic/models"
)

func SubmissionsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllSubmissions()
	h.WriteResponse(data, err, 400, w)
}

func SubmissionShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindSubmission(id)
	h.WriteResponse(data, err, 400, w)
}

func CreateSubmission(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateSubmission(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func UpdateSubmission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateSubmission(id, r.Body)
	h.WriteResponse(data, err, 422, w)
}

func DeleteSubmission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteSubmission(id)
	h.WriteResponse(data, err, 400, w)
}
