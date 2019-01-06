package models

import (
	"fmt"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
	"github.com/lib/pq"
)

type Submission struct {
	Id                     int    `json:"id"`
	PaperId                int    `json:"paper_id"`
	JournalId              int    `json:"journal_id"`
	JournalName            string `json:"journal_name"`
	Attempt                int    `json:"attempt"`
	ManuscriptSubmitted    string `json:"manuscript_submitted"`
	ManuscriptFeedbackDate string `json:"manuscript_feedback_date"`
	Accepted               bool   `json:"accepted"`
	ManuscriptEpub         string `json:"manuscript_epub"`
	ManuscriptPrinted      string `json:"manuscript_printed"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

func GetSubmissionsForPaper(paperId int, kawaiiChan chan []Submission) {
	var submissions []Submission
	var (
		id                       int
		paper_id                 int
		journal_id               int
		journal_name             string
		attempt                  int
		manuscript_submitted     time.Time
		manuscript_feedback_date time.Time
		accepted                 bool
		manuscript_epub          pq.NullTime
		manuscript_printed       pq.NullTime
		created_at               time.Time
		updated_at               time.Time
	)
	query := `
		SELECT submissions.*, journals.name AS journal_name FROM papers
		INNER JOIN submissions ON papers.id = submissions.paper_id
		INNER JOIN journals ON journals.id = submissions.journal_id
		WHERE papers.id = $1;
	`
	rows, err := database.DB.Query(query, paperId)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&paper_id,
			&journal_id,
			&attempt,
			&manuscript_submitted,
			&manuscript_feedback_date,
			&accepted,
			&manuscript_epub,
			&manuscript_printed,
			&created_at,
			&updated_at,
			&journal_name,
		)
		if err != nil {
			fmt.Println(err)
		}
		submission := Submission{
			Id:                     id,
			PaperId:                paper_id,
			JournalId:              journal_id,
			JournalName:            journal_name,
			Attempt:                attempt,
			ManuscriptSubmitted:    manuscript_submitted.String(),
			ManuscriptFeedbackDate: manuscript_feedback_date.String(),
			Accepted:               accepted,
			ManuscriptEpub:         NullTimeCheck(manuscript_epub),
			ManuscriptPrinted:      NullTimeCheck(manuscript_printed),
			CreatedAt:              created_at.String(),
			UpdatedAt:              updated_at.String(),
		}
		submissions = append(submissions, submission)
	}

	if err != nil {
		fmt.Println("Error getting paper's submissions", err)
	}
	kawaiiChan <- submissions
}
