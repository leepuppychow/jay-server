package models

import (
	// "encoding/json"
	// "errors"
	"errors"
	"fmt"
	"time"

	// "io"

	"github.com/leepuppychow/jay_medtronic/database"
	_ "github.com/leepuppychow/jay_medtronic/env"
)

type Paper struct {
	Id                      int
	Title                   string `json:"title"`
	Study_Id                int    `json:"study_id"`
	Device_Id               int    `json:"device_id"`
	InitialRequestEvaluated string `json:"initial_request_evaluated"`
	ManuscriptDrafted       string `json:"manuscript_drafted"`
	IntExtErp               string `json:"int_ext_erp"`
	CreatedAt               string `json:"created_at"`
	UpdatedAt               string `json:"updated_at"`
}

var (
	id                        int
	title                     string
	study_id                  int
	device_id                 int
	initial_request_evaluated time.Time
	manuscript_drafted        time.Time
	int_ext_erp               string
	created_at                time.Time
	updated_at                time.Time
)

func GetAllPapers(authToken string) ([]Paper, error) {
	var papers []Paper
	// Check for valid JWT:
	if !ValidToken(authToken) {
		return papers, errors.New("Unauthorized")
	}

	query := "SELECT * FROM papers"
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id,
			&title,
			&study_id,
			&device_id,
			&initial_request_evaluated,
			&manuscript_drafted,
			&int_ext_erp,
			&created_at,
			&updated_at)
		if err != nil {
			fmt.Println(err)
		}
		paper := Paper{
			Id:                      id,
			Title:                   title,
			Study_Id:                study_id,
			Device_Id:               device_id,
			InitialRequestEvaluated: initial_request_evaluated.String(),
			ManuscriptDrafted:       manuscript_drafted.String(),
			IntExtErp:               int_ext_erp,
			CreatedAt:               created_at.String(),
			UpdatedAt:               updated_at.String(),
		}
		papers = append(papers, paper)
	}
	if err != nil {
		return papers, err
	}
	return papers, nil
}

// func CreatePaper(body io.Reader) {

// }
