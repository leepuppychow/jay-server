package models

import (
	"fmt"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Author struct {
	Id        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAuthorsForPaper(paperId int, kawaiiChan chan []Author) {
	var authors []Author
	var (
		id         int
		first_name string
		last_name  string
		created_at time.Time
		updated_at time.Time
	)
	query := `
		SELECT authors.* FROM authors 
		INNER JOIN author_papers ON authors.id = author_papers.author_id
		INNER JOIN papers ON papers.id = author_papers.paper_id
		WHERE papers.id = $1
	`
	rows, err := database.DB.Query(query, paperId)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&first_name,
			&last_name,
			&created_at,
			&updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		author := Author{
			Id:        id,
			FirstName: first_name,
			LastName:  last_name,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		authors = append(authors, author)
	}

	if err != nil {
		fmt.Println("Error getting paper's authors", err)
	}
	kawaiiChan <- authors
}
