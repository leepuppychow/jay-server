package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/leepuppychow/jay_medtronic/helpers"
	"github.com/leepuppychow/jay_medtronic/models"
)

func JournalsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllJournals()
	h.WriteResponse(data, err, 400, w)
}

func JournalShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindJournal(id)
	h.WriteResponse(data, err, 400, w)
}

func CreateJournal(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateJournal(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func UpdateJournal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateJournal(id, r.Body)
	h.WriteResponse(data, err, 422, w)
}

func DeleteJournal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteJournal(id)
	h.WriteResponse(data, err, 400, w)
}
