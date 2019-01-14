package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Author struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllAuthors(authToken string) ([]Author, error) {
	// if !ValidToken(authToken) {
	// 	return []Author{}, errors.New("Unauthorized")
	// }
	var authors []Author
	var (
		id         int
		first_name string
		last_name  string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT authors.* FROM authors;`
	rows, err := database.DB.Query(query)
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
		return []Author{}, err
	}
	fmt.Println("Successful GET to authors index")
	return authors, nil
}

func FindAuthor(authorId int, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return []Author{}, errors.New("Unauthorized")
	// }
	var (
		id         int
		first_name string
		last_name  string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM authors WHERE authors.id = $1`
	err := database.DB.QueryRow(query, authorId).Scan(
		&id,
		&first_name,
		&last_name,
		&created_at,
		&updated_at,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	author := Author{
		Id:        id,
		FirstName: first_name,
		LastName:  last_name,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	fmt.Println("Successful GET to find author: ", id)
	return author, nil
}

func CreateAuthor(body io.Reader, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return []Author{}, errors.New("Unauthorized")
	// }
	var a Author
	err := json.NewDecoder(body).Decode(&a)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO authors (first_name, last_name, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query, a.FirstName, a.LastName)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful POST to create Author")
		return GeneralResponse{Message: "Author created successfully"}, nil
	}
}

func UpdateAuthor(AuthorId int, body io.Reader, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return []Author{}, errors.New("Unauthorized")
	// }
	var a Author
	err := json.NewDecoder(body).Decode(&a)
	query := `
		UPDATE authors
		SET 
			first_name = $2,
			last_name = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			authors.id = $1
	`
	_, err = database.DB.Exec(query, AuthorId, a.FirstName, a.LastName)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update author")
		return GeneralResponse{Message: "Author updated successfully"}, nil
	}
}

func DeleteAuthor(authorId int, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }
	query := `DELETE FROM authors WHERE id=$1`
	res, err := database.DB.Exec(query, authorId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete author with id %d", authorId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Author deleted successfully"}, nil
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
