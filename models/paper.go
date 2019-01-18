package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
	"github.com/lib/pq"
)

type Paper struct {
	Id                      int               `json:"id"`
	Title                   string            `json:"title"`
	Study_Id                int               `json:"study_id"`
	Journal_Id              int               `json:"journal_id"`
	InitialRequestEvaluated string            `json:"initial_request_evaluated"`
	ManuscriptDrafted       string            `json:"manuscript_drafted"`
	ManuscriptSubmitted     string            `json:"manuscript_submitted"`
	ManuscriptAccepted      string            `json:"manuscript_accepted"`
	ManuscriptEpub          string            `json:"manuscript_epub"`
	ManuscriptPrinted       string            `json:"manuscript_printed"`
	SubmissionAttempts      int               `json:"submission_attempts"`
	IntExtErp               string            `json:"int_ext_erp"`
	CreatedAt               string            `json:"created_at"`
	UpdatedAt               string            `json:"updated_at"`
	Study                   string            `json:"study"`
	Journal                 string            `json:"journal"`
	Devices                 []Device          `json:"devices"`
	Authors                 []Author          `json:"authors"`
	Figures                 []Figure          `json:"figures"`
	DataRequestForms        []DataRequestForm `json:"data_request_forms"`
}

type PaperResponse struct {
	PaperId int    `json:"paper_id"`
	Message string `json:"message"`
}

func GetAllPapers(authToken string) ([]Paper, error) {
	// if !ValidToken(authToken) {
	// 	return []Paper{}, errors.New("Unauthorized")
	// }
	var papers []Paper
	var (
		id                        int
		title                     string
		study_id                  int
		journal_id                int
		initial_request_evaluated pq.NullTime
		manuscript_drafted        pq.NullTime
		manuscript_submitted      pq.NullTime
		manuscript_accepted       pq.NullTime
		manuscript_epub           pq.NullTime
		manuscript_printed        pq.NullTime
		submission_attempts       int
		int_ext_erp               string
		created_at                time.Time
		updated_at                time.Time
		study                     string
		journal                   string
	)
	query := `
		SELECT papers.*, studies.name AS study, journals.name AS journal FROM papers 
		INNER JOIN studies ON papers.study_id = studies.id
		INNER JOIN journals ON papers.journal_id = journals.id;
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
			&journal_id,
			&initial_request_evaluated,
			&manuscript_drafted,
			&manuscript_submitted,
			&manuscript_accepted,
			&manuscript_epub,
			&manuscript_printed,
			&submission_attempts,
			&int_ext_erp,
			&created_at,
			&updated_at,
			&study,
			&journal,
		)
		if err != nil {
			fmt.Println(err)
		}

		authorsChannel := make(chan []Author)
		figuresChannel := make(chan []Figure)
		drfChannel := make(chan []DataRequestForm)
		devicesChannel := make(chan []Device)
		go GetAuthorsForPaper(id, authorsChannel)
		go GetFiguresForPaper(id, figuresChannel)
		go GetDataRequestFormsForPaper(id, drfChannel)
		go GetDevicesForPaper(id, devicesChannel)

		paper := Paper{
			Id:                      id,
			Title:                   title,
			Study_Id:                study_id,
			Journal_Id:              journal_id,
			InitialRequestEvaluated: NullTimeCheck(initial_request_evaluated),
			ManuscriptDrafted:       NullTimeCheck(manuscript_drafted),
			ManuscriptSubmitted:     NullTimeCheck(manuscript_submitted),
			ManuscriptAccepted:      NullTimeCheck(manuscript_accepted),
			ManuscriptEpub:          NullTimeCheck(manuscript_epub),
			ManuscriptPrinted:       NullTimeCheck(manuscript_printed),
			SubmissionAttempts:      submission_attempts,
			IntExtErp:               int_ext_erp,
			CreatedAt:               created_at.String(),
			UpdatedAt:               updated_at.String(),
			Study:                   study,
			Journal:                 journal,
			Authors:                 <-authorsChannel,
			Figures:                 <-figuresChannel,
			DataRequestForms:        <-drfChannel,
			Devices:                 <-devicesChannel,
		}
		papers = append(papers, paper)
	}
	if err != nil {
		return papers, err
	}
	fmt.Println("Successful GET to paper index")
	return papers, nil
}

func FindPaper(paperId int, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return Paper{}, errors.New("Unauthorized")
	// }
	var (
		id                        int
		title                     string
		study_id                  int
		journal_id                int
		initial_request_evaluated pq.NullTime
		manuscript_drafted        pq.NullTime
		manuscript_submitted      pq.NullTime
		manuscript_accepted       pq.NullTime
		manuscript_epub           pq.NullTime
		manuscript_printed        pq.NullTime
		submission_attempts       int
		int_ext_erp               string
		created_at                time.Time
		updated_at                time.Time
		study                     string
		journal                   string
	)

	queryString := `
		SELECT papers.*, studies.name AS study, journals.name AS journal FROM papers 
		INNER JOIN studies ON papers.study_id = studies.id
		INNER JOIN journals ON papers.journal_id = journals.id
		WHERE papers.id=$1;
	`
	err := database.DB.QueryRow(queryString, paperId).Scan(
		&id,
		&title,
		&study_id,
		&journal_id,
		&initial_request_evaluated,
		&manuscript_drafted,
		&manuscript_submitted,
		&manuscript_accepted,
		&manuscript_epub,
		&manuscript_printed,
		&submission_attempts,
		&int_ext_erp,
		&created_at,
		&updated_at,
		&study,
		&journal,
	)

	authorsChannel := make(chan []Author)
	figuresChannel := make(chan []Figure)
	drfChannel := make(chan []DataRequestForm)
	devicesChannel := make(chan []Device)
	go GetAuthorsForPaper(id, authorsChannel)
	go GetFiguresForPaper(id, figuresChannel)
	go GetDataRequestFormsForPaper(id, drfChannel)
	go GetDevicesForPaper(id, devicesChannel)

	paper := Paper{
		Id:                      id,
		Title:                   title,
		Study_Id:                study_id,
		Journal_Id:              journal_id,
		InitialRequestEvaluated: NullTimeCheck(initial_request_evaluated),
		ManuscriptDrafted:       NullTimeCheck(manuscript_drafted),
		ManuscriptSubmitted:     NullTimeCheck(manuscript_submitted),
		ManuscriptAccepted:      NullTimeCheck(manuscript_accepted),
		ManuscriptEpub:          NullTimeCheck(manuscript_epub),
		ManuscriptPrinted:       NullTimeCheck(manuscript_printed),
		SubmissionAttempts:      submission_attempts,
		IntExtErp:               int_ext_erp,
		CreatedAt:               created_at.String(),
		UpdatedAt:               updated_at.String(),
		Study:                   study,
		Journal:                 journal,
		Authors:                 <-authorsChannel,
		Figures:                 <-figuresChannel,
		DataRequestForms:        <-drfChannel,
		Devices:                 <-devicesChannel,
	}

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: "Error finding paper"}, err
	}
	return paper, nil
}

func CreatePaper(body io.Reader, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }
	var p Paper
	err := json.NewDecoder(body).Decode(&p)

	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	queryString := `
		INSERT INTO papers (
			title,
			study_id,
			journal_id, 
			initial_request_evaluated,
			manuscript_drafted,
			manuscript_submitted,
			manuscript_accepted,
			manuscript_epub,
			manuscript_printed,
			submission_attempts,
			int_ext_erp,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`

	lastInsertId := 0
	err = database.DB.QueryRow(queryString,
		p.Title,
		p.Study_Id,
		p.Journal_Id,
		p.InitialRequestEvaluated,
		p.ManuscriptDrafted,
		p.ManuscriptSubmitted,
		p.ManuscriptAccepted,
		p.ManuscriptEpub,
		p.ManuscriptPrinted,
		p.SubmissionAttempts,
		p.IntExtErp,
	).Scan(&lastInsertId)

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: "Unable to create paper"}, err
	} else {
		fmt.Println("Successful POST to create paper")
		return PaperResponse{PaperId: lastInsertId, Message: "Paper created successfully"}, nil
	}
}

func UpdatePaper(id int, body io.Reader, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }
	var p Paper
	err := json.NewDecoder(body).Decode(&p)
	queryString := `
		UPDATE papers
		SET
			title = $2,
			study_id = $3,
			journal_id = $4,
			initial_request_evaluated = $5,
			manuscript_drafted = $6,
			manuscript_submitted = $7,
			manuscript_accepted = $8,
			manuscript_epub = $9,
			manuscript_printed = $10,
			submission_attempts = $11,
			int_ext_erp = $12,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	_, err = database.DB.Exec(queryString,
		id,
		p.Title,
		p.Study_Id,
		p.Journal_Id,
		p.InitialRequestEvaluated,
		p.ManuscriptDrafted,
		p.ManuscriptSubmitted,
		p.ManuscriptAccepted,
		p.ManuscriptEpub,
		p.ManuscriptPrinted,
		p.SubmissionAttempts,
		p.IntExtErp,
	)

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: "Unable to update paper"}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update paper")
		return GeneralResponse{Message: "Paper updated successfully"}, nil
	}
}

func DeletePaper(id int, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }
	queryString := `DELETE FROM papers WHERE id=$1`
	res, err := database.DB.Exec(queryString, id)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete study with id %d", id)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Paper deleted successfully"}, nil
}
