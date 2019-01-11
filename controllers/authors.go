package controllers

import (
	"net/http"

	"github.com/leepuppychow/jay_medtronic/models"
)

func AuthorsIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllAuthors(r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateAuthor(r.Body, r.Header.Get("Authorization"))
	WriteResponse(data, err, 422, w)
}
