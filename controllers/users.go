package controllers

import (
	"errors"
	"net/http"

	"github.com/leepuppychow/jay_medtronic/auth"
	h "github.com/leepuppychow/jay_medtronic/helpers"
	"github.com/leepuppychow/jay_medtronic/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateUser(r.Body)
	h.WriteResponse(data, err, 422, w)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	data, err := models.LoginUser(r.Body)
	h.WriteResponse(data, err, 401, w)
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	valid := auth.ValidToken(r.Header.Get("Authorization"))
	if valid {
		h.WriteResponse("User Authenticated", nil, 200, w)
	} else {
		h.WriteResponse("", errors.New("Unauthorized"), 401, w)
	}
}
