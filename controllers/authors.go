package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	h "github.com/leepuppychow/jay_medtronic/helpers"
	"github.com/leepuppychow/jay_medtronic/models"
)

func AuthorsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllAuthors()
	h.WriteResponse(data, err, 400, w)
}

func AuthorShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindAuthor(id)
	h.WriteResponse(data, err, 400, w)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateAuthor(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateAuthor(id, r.Body)
	h.WriteResponse(data, err, 422, w)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteAuthor(id)
	h.WriteResponse(data, err, 400, w)
}
