package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func AuthorsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllAuthors()
	WriteResponse(data, err, 400, w)
}

func AuthorShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindAuthor(id)
	WriteResponse(data, err, 400, w)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateAuthor(r.Body)
	WriteResponse(data, err, 422, w)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateAuthor(id, r.Body)
	WriteResponse(data, err, 422, w)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteAuthor(id)
	WriteResponse(data, err, 400, w)
}
