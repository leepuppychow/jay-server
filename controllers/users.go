package controllers

import (
	"github.com/leepuppychow/jay_medtronic/models"

	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateUser(r.Body)
	WriteResponse(data, err, 422, w)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	data, err := models.LoginUser(r.Body)
	WriteResponse(data, err, 401, w)
}
