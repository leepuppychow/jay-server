package models

import (
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
