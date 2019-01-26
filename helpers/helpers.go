package h

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/lib/pq"
)

type GeneralResponse struct {
	Message string `json:"message"`
}

func NullTimeCheck(t pq.NullTime) string {
	if t.Valid {
		return t.Time.String()
	}
	return "NULL"
}

func InvalidTimeWillBeNull(t string) interface{} {
	_, err := time.Parse("2006-01-02", t)
	if err != nil {
		return nil
	}
	return t
}

func ToJSON(arg interface{}) []byte {
	json, err := json.MarshalIndent(arg, "", "   ")
	if err != nil {
		log.Println(err)
	}
	return json
}

func WriteResponse(data interface{}, err error, errorCode int, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(errorCode)
	}
	w.Write(ToJSON(data))
}
