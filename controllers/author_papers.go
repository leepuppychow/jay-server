package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/leepuppychow/jay_medtronic/models"
)

func AuthorPapersIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllAuthorPapers()
	WriteResponse(data, err, 400, w)
}

func AuthorPaperShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.FindAuthorPaper(id)
	WriteResponse(data, err, 400, w)
}

func CreateAuthorPaper(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateAuthorPaper(r.Body)
	WriteResponse(data, err, 422, w)
}

func UpdateAuthorPaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.UpdateAuthorPaper(id, r.Body)
	WriteResponse(data, err, 422, w)
}

func DeleteAuthorPaper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	data, err := models.DeleteAuthorPaper(id)
	WriteResponse(data, err, 400, w)
}
