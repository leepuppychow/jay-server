package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func JournalsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllJournals(r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func JournalShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindJournal(id, r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func CreateJournal(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateJournal(r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

func UpdateJournal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateJournal(id, r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}

func DeleteJournal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteJournal(id, r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}