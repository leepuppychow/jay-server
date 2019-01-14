package models

import (
	"fmt"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Study struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllStudies(authToken string) ([]Study, error) {
	// if !ValidToken(authToken) {
	// 	return []Study{}, errors.New("Unauthorized")
	// }
	var studies []Study
	var (
		id         int
		name       string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT studies.* FROM studies;`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&name,
			&created_at,
			&updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		study := Study{
			Id:        id,
			Name:      name,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		studies = append(studies, study)
	}
	if err != nil {
		return []Study{}, err
	}
	fmt.Println("Successful GET to studies index")
	return studies, nil
}
