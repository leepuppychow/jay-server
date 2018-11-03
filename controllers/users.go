package users

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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create USER PATH")
	w.Write(ToJSON("create user path"))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login User PATH")
	w.Write(ToJSON("Login User path"))
}