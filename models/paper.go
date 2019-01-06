package models

import (
	// "encoding/json"
	// "errors"
	"errors"
	"fmt"
	"time"

	// "io"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Paper struct {
	Id                      int               `json:"id"`
	Title                   string            `json:"title"`
	Study_Id                int               `json:"study_id"`
	Device_Id               int               `json:"device_id"`
	InitialRequestEvaluated string            `json:"initial_request_evaluated"`
	ManuscriptDrafted       string            `json:"manuscript_drafted"`
	IntExtErp               string            `json:"int_ext_erp"`
	CreatedAt               string            `json:"created_at"`
	UpdatedAt               string            `json:"updated_at"`
	Study                   string            `json:"study"`
	Device                  string            `json:"device"`
	Authors                 []Author          `json:"authors"`
	Figures                 []Figure          `json:"figures"`
	DataRequestForms        []DataRequestForm `json:"data_request_forms"`
	Submissions             []Submission      `json:"submissions"`
}

func GetAllPapers(authToken string) ([]Paper, error) {
	var papers []Paper
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
		study                     string
		device                    string
	)
	// Check for valid JWT:
	if !ValidToken(authToken) {
		return papers, errors.New("Unauthorized")
	}

	query := `
		SELECT papers.*, studies.name AS study, devices.name AS device FROM papers 
		INNER JOIN studies ON papers.study_id = studies.id
		INNER JOIN devices ON papers.device_id = devices.id;
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&id,
			&title,
			&study_id,
			&device_id,
			&initial_request_evaluated,
			&manuscript_drafted,
			&int_ext_erp,
			&created_at,
			&updated_at,
			&study,
			&device,
		)
		if err != nil {
			fmt.Println(err)
		}

		authorsChannel := make(chan []Author)
		figuresChannel := make(chan []Figure)
		drfChannel := make(chan []DataRequestForm)
		submissionChannel := make(chan []Submission)
		go GetAuthorsForPaper(id, authorsChannel)
		go GetFiguresForPaper(id, figuresChannel)
		go GetDataRequestFormsForPaper(id, drfChannel)
		go GetSubmissionsForPaper(id, submissionChannel)

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
			Study:                   study,
			Device:                  device,
			Authors:                 <-authorsChannel,
			Figures:                 <-figuresChannel,
			DataRequestForms:        <-drfChannel,
			Submissions:             <-submissionChannel,
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
