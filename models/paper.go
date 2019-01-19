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
	Id                      int          `json:"id"`
	Title                   string       `json:"title"`
	Study_Id                int          `json:"study_id"`
	InitialRequestEvaluated string       `json:"initial_request_evaluated"`
	DrfRounds               int          `json:"drf_rounds"`
	DrfCompleted            string       `json:"drf_completed"`
	DrfRequestedDelivery    string       `json:"drf_requested_delivery"`
	DrfActualDelivery       string       `json:"drf_actual_delivery"`
	DataRefinementComplete  string       `json:"data_refinement_complete"`
	ManuscriptDrafted       string       `json:"manuscript_drafted"`
	ManuscriptAccepted      string       `json:"manuscript_accepted"`
	ManuscriptEpub          string       `json:"manuscript_epub"`
	ManuscriptPrinted       string       `json:"manuscript_printed"`
	IntExtErp               string       `json:"int_ext_erp"`
	CreatedAt               string       `json:"created_at"`
	UpdatedAt               string       `json:"updated_at"`
	Study                   string       `json:"study"`
	Journal                 string       `json:"journal"`
	Submissions             []Submission `json:"submissions"`
	Submission_Ids          []int        `json:"submission_ids"`
	Devices                 []Device     `json:"devices"`
	Device_Ids              []int        `json:"device_ids"`
	Authors                 []Author     `json:"authors"`
	Author_Ids              []int        `json:"author_ids"`
	Figures                 []Figure     `json:"figures"`
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
		initial_request_evaluated pq.NullTime
		drf_rounds                int
		drf_completed             pq.NullTime
		drf_requested_delivery    pq.NullTime
		drf_actual_delivery       pq.NullTime
		data_refinement_complete  pq.NullTime
		manuscript_drafted        pq.NullTime
		manuscript_accepted       pq.NullTime
		manuscript_epub           pq.NullTime
		manuscript_printed        pq.NullTime
		int_ext_erp               string
		created_at                time.Time
		updated_at                time.Time
		study                     string
		journal                   string
	)
	query := `
		SELECT papers.*, studies.name AS study, journals.name AS journal FROM papers 
		INNER JOIN studies ON papers.study_id = studies.id
		INNER JOIN submissions ON papers.id = submissions.paper_id
		INNER JOIN journals ON submissions.journal_id = journals.id
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&id,
			&study_id,
			&title,
			&int_ext_erp,
			&drf_rounds,
			&initial_request_evaluated,
			&drf_completed,
			&drf_requested_delivery,
			&drf_actual_delivery,
			&data_refinement_complete,
			&manuscript_drafted,
			&manuscript_accepted,
			&manuscript_epub,
			&manuscript_printed,
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
		devicesChannel := make(chan []Device)
		submissionsChannel := make(chan []Submission)
		go GetAuthorsForPaper(id, authorsChannel)
		go GetFiguresForPaper(id, figuresChannel)
		go GetDevicesForPaper(id, devicesChannel)
		go GetSubmissionsForPaper(id, submissionsChannel)

		paper := Paper{
			Id:                      id,
			Title:                   title,
			Study_Id:                study_id,
			InitialRequestEvaluated: NullTimeCheck(initial_request_evaluated),
			DrfRounds:               drf_rounds,
			DrfCompleted:            NullTimeCheck(drf_completed),
			DrfRequestedDelivery:    NullTimeCheck(drf_requested_delivery),
			DrfActualDelivery:       NullTimeCheck(drf_actual_delivery),
			DataRefinementComplete:  NullTimeCheck(data_refinement_complete),
			ManuscriptDrafted:       NullTimeCheck(manuscript_drafted),
			ManuscriptAccepted:      NullTimeCheck(manuscript_accepted),
			ManuscriptEpub:          NullTimeCheck(manuscript_epub),
			ManuscriptPrinted:       NullTimeCheck(manuscript_printed),
			IntExtErp:               int_ext_erp,
			CreatedAt:               created_at.String(),
			UpdatedAt:               updated_at.String(),
			Study:                   study,
			Journal:                 journal,
			Authors:                 <-authorsChannel,
			Figures:                 <-figuresChannel,
			Devices:                 <-devicesChannel,
			Submissions:             <-submissionsChannel,
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
		study_id                  int
		title                     string
		initial_request_evaluated pq.NullTime
		drf_rounds                int
		drf_completed             pq.NullTime
		drf_requested_delivery    pq.NullTime
		drf_actual_delivery       pq.NullTime
		data_refinement_complete  pq.NullTime
		manuscript_drafted        pq.NullTime
		manuscript_accepted       pq.NullTime
		manuscript_epub           pq.NullTime
		manuscript_printed        pq.NullTime
		int_ext_erp               string
		created_at                time.Time
		updated_at                time.Time
		study                     string
		journal                   string
	)

	queryString := `
		SELECT papers.*, studies.name AS study, journals.name AS journal FROM papers 
		INNER JOIN studies ON papers.study_id = studies.id
		INNER JOIN submissions ON papers.id = submissions.paper_id
		INNER JOIN journals ON submissions.journal_id = journals.id
		WHERE papers.id=$1
	`
	err := database.DB.QueryRow(queryString, paperId).Scan(
		&id,
		&study_id,
		&title,
		&int_ext_erp,
		&drf_rounds,
		&initial_request_evaluated,
		&drf_completed,
		&drf_requested_delivery,
		&drf_actual_delivery,
		&data_refinement_complete,
		&manuscript_drafted,
		&manuscript_accepted,
		&manuscript_epub,
		&manuscript_printed,
		&created_at,
		&updated_at,
		&study,
		&journal,
	)

	authorsChannel := make(chan []Author)
	figuresChannel := make(chan []Figure)
	devicesChannel := make(chan []Device)
	submissionsChannel := make(chan []Submission)
	go GetAuthorsForPaper(id, authorsChannel)
	go GetFiguresForPaper(id, figuresChannel)
	go GetDevicesForPaper(id, devicesChannel)
	go GetSubmissionsForPaper(id, submissionsChannel)

	paper := Paper{
		Id:                      id,
		Title:                   title,
		Study_Id:                study_id,
		InitialRequestEvaluated: NullTimeCheck(initial_request_evaluated),
		DrfRounds:               drf_rounds,
		DrfCompleted:            NullTimeCheck(drf_completed),
		DrfRequestedDelivery:    NullTimeCheck(drf_requested_delivery),
		DrfActualDelivery:       NullTimeCheck(drf_actual_delivery),
		DataRefinementComplete:  NullTimeCheck(data_refinement_complete),
		ManuscriptDrafted:       NullTimeCheck(manuscript_drafted),
		ManuscriptAccepted:      NullTimeCheck(manuscript_accepted),
		ManuscriptEpub:          NullTimeCheck(manuscript_epub),
		ManuscriptPrinted:       NullTimeCheck(manuscript_printed),
		IntExtErp:               int_ext_erp,
		CreatedAt:               created_at.String(),
		UpdatedAt:               updated_at.String(),
		Study:                   study,
		Journal:                 journal,
		Authors:                 <-authorsChannel,
		Figures:                 <-figuresChannel,
		Devices:                 <-devicesChannel,
		Submissions:             <-submissionsChannel,
	}

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: "Error finding paper"}, err
	}
	return paper, nil
}

func CreatePaperQuery(p Paper) (int, error) {
	queryString := `
		INSERT INTO papers (
			title,
			study_id,
			initial_request_evaluated,
			drf_rounds,
			drf_completed,
			drf_requested_delivery,
			drf_actual_delivery,
			data_refinement_complete,
			manuscript_drafted,
			manuscript_accepted,
			manuscript_epub,
			manuscript_printed,
			int_ext_erp,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`
	lastInsertId := 0
	err := database.DB.QueryRow(queryString,
		p.Title,
		p.Study_Id,
		p.InitialRequestEvaluated,
		p.DrfRounds,
		p.DrfCompleted,
		p.DrfRequestedDelivery,
		p.DrfActualDelivery,
		p.DataRefinementComplete,
		p.ManuscriptDrafted,
		p.ManuscriptAccepted,
		p.ManuscriptEpub,
		p.ManuscriptPrinted,
		p.IntExtErp,
	).Scan(&lastInsertId)

	return lastInsertId, err
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
	lastInsertId, err := CreatePaperQuery(p)

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: "Unable to create paper"}, err
	} else {
		fmt.Println("Successful POST to create paper")
		return PaperResponse{PaperId: lastInsertId, Message: "Paper created successfully"}, nil
	}
}

func SpecialCreatePaper(body io.Reader, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }
	var p Paper
	err := json.NewDecoder(body).Decode(&p)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	paperId, err := CreatePaperQuery(p)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}

	// Create the author_papers and device_papers entries now
	for _, authorId := range p.Author_Ids {
		ap := AuthorPaper{
			PaperId:  paperId,
			AuthorId: authorId,
		}
		_, err = CreateAuthorPaperQuery(ap)
		if err != nil {
			fmt.Println(err)
			return GeneralResponse{Message: err.Error()}, err
		}
	}
	for _, deviceId := range p.Device_Ids {
		dp := DevicePaper{
			PaperId:  paperId,
			DeviceId: deviceId,
		}
		_, err = CreateDevicePaperQuery(dp)
		if err != nil {
			fmt.Println(err)
			return GeneralResponse{Message: err.Error()}, err
		}
	}
	return GeneralResponse{Message: "Paper, author_papers, and device_papers created successfully"}, nil
}

func UpdatePaper(paperId int, body io.Reader, authToken string) (GeneralResponse, error) {
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
			initial_request_evaluated = $4,
			drf_rounds = $5,
			drf_completed = $6,
			drf_requested_delivery = $7,
			drf_actual_delivery = $8,
			data_refinement_complete = $9,
			manuscript_drafted = $10,
			manuscript_accepted = $11,
			manuscript_epub = $12,
			manuscript_printed = $13,
			int_ext_erp = $14,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`
	_, err = database.DB.Exec(queryString,
		paperId,
		p.Title,
		p.Study_Id,
		p.InitialRequestEvaluated,
		p.DrfRounds,
		p.DrfCompleted,
		p.DrfRequestedDelivery,
		p.DrfActualDelivery,
		p.DataRefinementComplete,
		p.ManuscriptDrafted,
		p.ManuscriptAccepted,
		p.ManuscriptEpub,
		p.ManuscriptPrinted,
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
