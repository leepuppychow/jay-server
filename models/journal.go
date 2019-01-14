package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Journal struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllJournals(authToken string) ([]Journal, error) {
	// if !ValidToken(authToken) {
	// 	return []Journal{}, errors.New("Unauthorized")
	// }
	var journals []Journal
	var (
		id         int
		name       string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT journals.* FROM journals;`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&name,
			&created_at,
			&updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		journal := Journal{
			Id:        id,
			Name:      name,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		journals = append(journals, journal)
	}
	if err != nil {
		return []Journal{}, err
	}
	fmt.Println("Successful GET to journals index")
	return journals, nil
}

func FindJournal(journalId int, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return []Journal{}, errors.New("Unauthorized")
	// }
	var (
		id         int
		name       string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM journals WHERE journals.id = $1`
	err := database.DB.QueryRow(query, journalId).Scan(
		&id,
		&name,
		&created_at,
		&updated_at,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	journal := Journal{
		Id:        id,
		Name:      name,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	fmt.Println("Successful GET to find journal: ", id)
	return journal, nil
}

func CreateJournal(body io.Reader, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return []Journal{}, errors.New("Unauthorized")
	// }
	var j Journal
	err := json.NewDecoder(body).Decode(&j)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO journals (name, created_at, updated_at)
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query, j.Name)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful POST to create journal")
		return GeneralResponse{Message: "journal created successfully"}, nil
	}
}

func UpdateJournal(journalId int, body io.Reader, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return []Journal{}, errors.New("Unauthorized")
	// }
	var j Journal
	err := json.NewDecoder(body).Decode(&j)
	query := `
		UPDATE journals
		SET 
			name = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			journals.id = $1
	`
	_, err = database.DB.Exec(query, journalId, j.Name)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update journal")
		return GeneralResponse{Message: "Journal updated successfully"}, nil
	}
}

func DeleteJournal(journalId int, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }
	query := `DELETE FROM journals WHERE id=$1`
	res, err := database.DB.Exec(query, journalId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete journal with id %d", journalId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Journal deleted successfully"}, nil
}
