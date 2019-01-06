package models

import (
	"github.com/lib/pq"
)

func NullTimeCheck(t pq.NullTime) string {
	if t.Valid {
		return t.Time.String()
	}
	return "NULL"
}
