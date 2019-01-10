package controllers

import (
	"encoding/json"
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

