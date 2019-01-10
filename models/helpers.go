package models

import (
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
