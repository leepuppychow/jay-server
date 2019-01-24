package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type AuthorPaper struct {
	Id        int    `json:"id"`
	PaperId   int    `json:"paper_id"`
	AuthorId  int    `json:"author_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllAuthorPapers(authToken string) ([]AuthorPaper, error) {
	if !ValidToken(authToken) {
		return []AuthorPaper{}, errors.New("Unauthorized")
	}
	var aps []AuthorPaper
	var (
		id         int
		paper_id   int
		author_id  int
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM author_papers`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&paper_id,
			&author_id,
			&created_at,
			&updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		ap := AuthorPaper{
			Id:        id,
			PaperId:   paper_id,
			AuthorId:  author_id,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		aps = append(aps, ap)
	}
	if err != nil {
		return []AuthorPaper{}, err
	}
	fmt.Println("Successful GET to AuthorPapers index")
	return aps, nil
}

func FindAuthorPaper(authorPaperId int, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []AuthorPaper{}, errors.New("Unauthorized")
	}
	var (
		id         int
		paper_id   int
		author_id  int
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM author_papers WHERE id = $1`
	err := database.DB.QueryRow(query, authorPaperId).Scan(
		&id,
		&paper_id,
		&author_id,
		&created_at,
		&updated_at,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	fp := AuthorPaper{
		Id:        id,
		PaperId:   paper_id,
		AuthorId:  author_id,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	fmt.Println("Successful GET to find AuthorPaper:", id)
	return fp, nil
}

func CreateAuthorPaperQuery(ap AuthorPaper) (int, error) {
	query := `
		INSERT INTO author_papers (paper_id, author_id, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`
	lastInsertId := 0
	err := database.DB.QueryRow(query, ap.PaperId, ap.AuthorId).Scan(&lastInsertId)
	if err != nil {
		fmt.Println(err)
	}
	return lastInsertId, err
}

func CreateAuthorPaper(body io.Reader, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []AuthorPaper{}, errors.New("Unauthorized")
	}
	var ap AuthorPaper
	err := json.NewDecoder(body).Decode(&ap)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	_, err = CreateAuthorPaperQuery(ap)

	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful POST to create AuthorPaper")
		return GeneralResponse{Message: "AuthorPaper created successfully"}, nil
	}
}

func UpdateAuthorPaper(authorPaperId int, body io.Reader, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []AuthorPaper{}, errors.New("Unauthorized")
	}
	var fp AuthorPaper
	err := json.NewDecoder(body).Decode(&fp)
	query := `
		UPDATE author_papers
		SET 
			paper_id = $2,
			author_id = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $1
	`
	_, err = database.DB.Exec(query, authorPaperId, fp.PaperId, fp.AuthorId)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update AuthorPaper")
		return GeneralResponse{Message: "AuthorPaper updated successfully"}, nil
	}
}

func DeleteAuthorPaper(authorPaperId int, authToken string) (GeneralResponse, error) {
	if !ValidToken(authToken) {
		return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	}
	query := `DELETE FROM author_papers WHERE id=$1`
	res, err := database.DB.Exec(query, authorPaperId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete AuthorPaper with id %d", authorPaperId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "AuthorPaper deleted successfully"}, nil
}
