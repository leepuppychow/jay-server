package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/leepuppychow/jay-server/database"
	h "github.com/leepuppychow/jay-server/helpers"
)

type Journal struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllJournals() ([]Journal, error) {
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
		log.Println(err)
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
			log.Println(err)
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
	log.Println("Successful GET to journals index")
	return journals, nil
}

func FindJournal(journalId int) (interface{}, error) {
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
		log.Println(err)
		return h.GeneralResponse{Message: err.Error()}, err
	}
	journal := Journal{
		Id:        id,
		Name:      name,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	log.Println("Successful GET to find journal: ", id)
	return journal, nil
}

func CreateJournal(body io.Reader) (interface{}, error) {
	var j Journal
	err := json.NewDecoder(body).Decode(&j)
	if err != nil {
		return h.GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO journals (name, created_at, updated_at)
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query, j.Name)
	if err != nil {
		log.Println(err)
		return h.GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful POST to create journal")
		return h.GeneralResponse{Message: "journal created successfully"}, nil
	}
}

func UpdateJournal(journalId int, body io.Reader) (interface{}, error) {
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
		log.Println(err)
		return h.GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful PUT/PATCH to update journal")
		return h.GeneralResponse{Message: "Journal updated successfully"}, nil
	}
}

func DeleteJournal(journalId int) (h.GeneralResponse, error) {
	query := `DELETE FROM journals WHERE id=$1`
	res, err := database.DB.Exec(query, journalId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete journal with id %d", journalId)
		err = errors.New("Did not find row with specified ID")
		return h.GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return h.GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return h.GeneralResponse{Message: "Journal deleted successfully"}, nil
}
