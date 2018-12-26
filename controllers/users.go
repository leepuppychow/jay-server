package controllers

import (
	"encoding/json"
	"github.com/leepuppychow/jay_medtronic/models"

	"fmt"
	"net/http"
)

func ToJSON(arg interface{}) []byte {
	json, err := json.MarshalIndent(arg, "", "   ")
	if err != nil {
		fmt.Println(err)
	}
	return json
}

func WriteResponse(data interface{}, err error, errorCode int, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(errorCode)
	}
	w.Write(ToJSON(data))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := models.CreateUser(r.Body)
	WriteResponse(data, err, 422, w)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	data, err := models.LoginUser(r.Body)
	WriteResponse(data, err, 401, w)
}

func Test(w http.ResponseWriter, r *http.Request) {
	data, err := models.TestAuth(r.Header.Get("Authorization"))
	WriteResponse(data, err, 401, w)
}
