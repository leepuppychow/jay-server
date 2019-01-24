package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Study struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllStudies(authToken string) ([]Study, error) {
	if !ValidToken(authToken) {
		return []Study{}, errors.New("Unauthorized")
	}
	var studies []Study
	var (
		id         int
		name       string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT studies.* FROM studies;`
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
		study := Study{
			Id:        id,
			Name:      name,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		studies = append(studies, study)
	}
	if err != nil {
		return []Study{}, err
	}
	fmt.Println("Successful GET to studies index")
	return studies, nil
}

func FindStudy(studyId int, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []Study{}, errors.New("Unauthorized")
	}
	var (
		id         int
		name       string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM studies WHERE studies.id = $1`
	err := database.DB.QueryRow(query, studyId).Scan(
		&id,
		&name,
		&created_at,
		&updated_at,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	study := Study{
		Id:        id,
		Name:      name,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	fmt.Println("Successful GET to find study: ", id)
	return study, nil
}

func CreateStudy(body io.Reader, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []Study{}, errors.New("Unauthorized")
	}
	var s Study
	err := json.NewDecoder(body).Decode(&s)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO studies (name, created_at, updated_at)
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query, s.Name)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful POST to create study")
		return GeneralResponse{Message: "Study created successfully"}, nil
	}
}

func UpdateStudy(studyId int, body io.Reader, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []Study{}, errors.New("Unauthorized")
	}
	var s Study
	err := json.NewDecoder(body).Decode(&s)
	query := `
		UPDATE studies
		SET 
			name = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			studies.id = $1
	`
	_, err = database.DB.Exec(query, studyId, s.Name)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update study")
		return GeneralResponse{Message: "Study updated successfully"}, nil
	}
}

func DeleteStudy(studyId int, authToken string) (GeneralResponse, error) {
	if !ValidToken(authToken) {
		return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	}
	query := `DELETE FROM studies WHERE id=$1`
	res, err := database.DB.Exec(query, studyId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete study with id %d", studyId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Study deleted successfully"}, nil
}
