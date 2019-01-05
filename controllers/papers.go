package controllers

import (
	"net/http"

	"github.com/leepuppychow/jay_medtronic/models"
)

func PapersIndex(w http.ResponseWriter, r *http.Request) {
	data, err := models.GetAllPapers(r.Header.Get("Authorization"))
	WriteResponse(data, err, 400, w)
}

// func CreatePaper(w http.ResponseWriter, r *http.Request) {
// 	data, err := models.CreatePaper(r.Body)
// 	WriteResponse(data, err, 422, w)
// }
