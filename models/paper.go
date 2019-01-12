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
		initial_request_evaluated pq.NullTime
		manuscript_drafted        pq.NullTime
		int_ext_erp               string
		created_at                time.Time
		updated_at                time.Time
		study                     string
		device                    string
	)
	// if !ValidToken(authToken) {
	// 	return papers, errors.New("Unauthorized")
	// }

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
			InitialRequestEvaluated: NullTimeCheck(initial_request_evaluated),
			ManuscriptDrafted:       NullTimeCheck(manuscript_drafted),
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
	fmt.Println("Successful GET to paper index")
	return papers, nil
}

func FindPaper(paperId int, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return papers, errors.New("Unauthorized")
	// }
	var (
		id                        int
		title                     string
		study_id                  int
		device_id                 int
		initial_request_evaluated pq.NullTime
		manuscript_drafted        pq.NullTime
		int_ext_erp               string
		created_at                time.Time
		updated_at                time.Time
		study                     string
		device                    string
	)

	queryString := `
		SELECT papers.*, studies.name AS study, devices.name AS device FROM papers 
		INNER JOIN studies ON papers.study_id = studies.id
		INNER JOIN devices ON papers.device_id = devices.id
		WHERE papers.id=$1;
	`
	err := database.DB.QueryRow(queryString, paperId).Scan(
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
		InitialRequestEvaluated: NullTimeCheck(initial_request_evaluated),
		ManuscriptDrafted:       NullTimeCheck(manuscript_drafted),
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

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: "Error finding paper"}, err
	}
	return paper, nil
}

func CreatePaper(body io.Reader, authToken string) (GeneralResponse, error) {
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
			device_id, 
			initial_request_evaluated,
			manuscript_drafted,
			int_ext_erp,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(queryString,
		p.Title,
		p.Study_Id,
		p.Device_Id,
		p.InitialRequestEvaluated,
		p.ManuscriptDrafted,
		p.IntExtErp)

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: "Unable to create paper"}, err
	} else {
		fmt.Println("Successful POST to create paper")
		return GeneralResponse{Message: "Paper created successfully"}, nil
	}
}

func UpdatePaper(id int, body io.Reader, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }

	var p Paper
	err := json.NewDecoder(body).Decode(&p)

	fmt.Println(p)

	queryString := `
		UPDATE papers
		SET 
			title = $2,
			study_id = $3,
			device_id = $4,
			initial_request_evaluated = $5,
			manuscript_drafted = $6,
			int_ext_erp = $7,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	_, err = database.DB.Exec(queryString,
		id,
		p.Title,
		p.Study_Id,
		p.Device_Id,
		p.InitialRequestEvaluated,
		p.ManuscriptDrafted,
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
		errorMessage := fmt.Sprintf("Error when trying to delete Word with id %d", id)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Paper deleted successfully"}, nil
}
