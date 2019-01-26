package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/lib/pq"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Submission struct {
	Id                  int    `json:"id"`
	PaperId             int    `json:"paper_id"`
	JournalId           int    `json:"journal_id"`
	Journal             string `json:"journal"`
	Attempt             int    `json:"attempt"`
	ManuscriptSubmitted string `json:"manuscript_submitted"`
	ManuscriptRejected  string `json:"manuscript_rejected"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

func GetAllSubmissions() ([]Submission, error) {
	var submissions []Submission
	var (
		id                   int
		paper_id             int
		journal_id           int
		attempt              int
		manuscript_submitted pq.NullTime
		manuscript_rejected  pq.NullTime
		created_at           time.Time
		updated_at           time.Time
		journal              string
	)
	query := `
		SELECT submissions.*, journals.name AS journal FROM submissions
		INNER JOIN journals ON submissions.journal_id = journals.id
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&paper_id,
			&journal_id,
			&attempt,
			&manuscript_submitted,
			&manuscript_rejected,
			&created_at,
			&updated_at,
			&journal,
		)
		if err != nil {
			log.Println(err)
		}
		submission := Submission{
			Id:                  id,
			PaperId:             paper_id,
			JournalId:           journal_id,
			Attempt:             attempt,
			ManuscriptSubmitted: NullTimeCheck(manuscript_submitted),
			ManuscriptRejected:  NullTimeCheck(manuscript_rejected),
			CreatedAt:           created_at.String(),
			UpdatedAt:           updated_at.String(),
			Journal:             journal,
		}
		submissions = append(submissions, submission)
	}
	if err != nil {
		return []Submission{}, err
	}
	log.Println("Successful GET to Submissions index")
	return submissions, nil
}

func FindSubmission(submissionId int) (interface{}, error) {
	var (
		id                   int
		paper_id             int
		journal_id           int
		attempt              int
		manuscript_submitted pq.NullTime
		manuscript_rejected  pq.NullTime
		created_at           time.Time
		updated_at           time.Time
		journal              string
	)
	query := `
		SELECT submissions.*, journals.name AS journal FROM submissions
		INNER JOIN journals ON submissions.journal_id = journals.id
		WHERE submissions.id = $1
	`
	err := database.DB.QueryRow(query, submissionId).Scan(
		&id,
		&paper_id,
		&journal_id,
		&attempt,
		&manuscript_submitted,
		&manuscript_rejected,
		&created_at,
		&updated_at,
		&journal,
	)
	if err != nil {
		log.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	s := Submission{
		Id:                  id,
		PaperId:             paper_id,
		JournalId:           journal_id,
		Attempt:             attempt,
		ManuscriptSubmitted: NullTimeCheck(manuscript_submitted),
		ManuscriptRejected:  NullTimeCheck(manuscript_rejected),
		CreatedAt:           created_at.String(),
		UpdatedAt:           updated_at.String(),
		Journal:             journal,
	}
	log.Println("Successful GET to find Submission: ", id)
	return s, nil
}

func CreateSubmissionQuery(s Submission) (int, error) {
	query := `
		INSERT INTO submissions (
			paper_id,
			journal_id,
			attempt,
			manuscript_submitted,
			manuscript_rejected,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`
	lastInsertId := 0
	err := database.DB.QueryRow(query,
		s.PaperId,
		s.JournalId,
		s.Attempt,
		InvalidTimeWillBeNull(s.ManuscriptSubmitted),
		InvalidTimeWillBeNull(s.ManuscriptRejected),
	).Scan(&lastInsertId)

	return lastInsertId, err
}

func CreateSubmission(body io.Reader) (interface{}, error) {
	var s Submission
	err := json.NewDecoder(body).Decode(&s)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	lastInsertId, err := CreateSubmissionQuery(s)
	if err != nil {
		log.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful POST to create Submission", lastInsertId)
		return GeneralResponse{Message: "Submission created successfully"}, nil
	}
}

func UpdateSubmission(submissionId int, body io.Reader) (interface{}, error) {
	var s Submission
	err := json.NewDecoder(body).Decode(&s)
	query := `
		UPDATE submissions
		SET 
			paper_id = $2,
			journal_id = $3,
			attempt = $4,
			manuscript_submitted = $5,
			manuscript_rejected = $6,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			Submissions.id = $1
	`
	_, err = database.DB.Exec(query,
		submissionId,
		s.PaperId,
		s.JournalId,
		s.Attempt,
		s.ManuscriptSubmitted,
		s.ManuscriptRejected,
	)
	if err != nil {
		log.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful PUT/PATCH to update Submission")
		return GeneralResponse{Message: "Submission updated successfully"}, nil
	}
}

func DeleteSubmission(submissionId int) (GeneralResponse, error) {
	query := `DELETE FROM submissions WHERE id=$1`
	res, err := database.DB.Exec(query, submissionId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete Submission with id %d", submissionId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Submission deleted successfully"}, nil
}

func GetSubmissionsForPaper(paperId int) <-chan []Submission {
	ch := make(chan []Submission)
	go func() {
		var submissions []Submission
		var (
			id                   int
			paper_id             int
			journal_id           int
			attempt              int
			manuscript_submitted pq.NullTime
			manuscript_rejected  pq.NullTime
			created_at           time.Time
			updated_at           time.Time
			journal              string
		)
		query := `
			SELECT submissions.*, journals.name AS journal FROM submissions 
			INNER JOIN papers ON papers.id = submissions.paper_id
			INNER JOIN journals ON submissions.journal_id = journals.id
			WHERE papers.id = $1
		`
		rows, err := database.DB.Query(query, paperId)
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(
				&id,
				&paper_id,
				&journal_id,
				&attempt,
				&manuscript_submitted,
				&manuscript_rejected,
				&created_at,
				&updated_at,
				&journal,
			)
			if err != nil {
				log.Println(err)
			}
			s := Submission{
				Id:                  id,
				PaperId:             paper_id,
				JournalId:           journal_id,
				Attempt:             attempt,
				ManuscriptSubmitted: NullTimeCheck(manuscript_submitted),
				ManuscriptRejected:  NullTimeCheck(manuscript_rejected),
				CreatedAt:           created_at.String(),
				UpdatedAt:           updated_at.String(),
				Journal:             journal,
			}
			submissions = append(submissions, s)
		}
		if err != nil {
			log.Println("Error getting paper's Submissions", err)
		}
		ch <- submissions
	}()
	return ch
}
