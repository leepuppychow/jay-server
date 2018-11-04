package users

import (
	"encoding/json"
	"jay_medtronic/models"

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
	enableCORS(&w)
	if err != nil {
		w.WriteHeader(errorCode)
	}
	w.Write(ToJSON(data))
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := user.Create(r.Body)
	WriteResponse(data, err, 422, w)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	data, err := user.Login(r.Body)
	WriteResponse(data, err, 401, w)
}
