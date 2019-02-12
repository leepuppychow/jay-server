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

type FigurePaper struct {
	Id        int    `json:"id"`
	PaperId   int    `json:"paper_id"`
	FigureId  int    `json:"figure_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllFigurePapers() ([]FigurePaper, error) {
	var fps []FigurePaper
	var (
		id         int
		paper_id   int
		figure_id  int
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM figure_papers`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&paper_id,
			&figure_id,
			&created_at,
			&updated_at,
		)
		if err != nil {
			log.Println(err)
		}
		fp := FigurePaper{
			Id:        id,
			PaperId:   paper_id,
			FigureId:  figure_id,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		fps = append(fps, fp)
	}
	if err != nil {
		return []FigurePaper{}, err
	}
	log.Println("Successful GET to FigurePapers index")
	return fps, nil
}

func FindFigurePaper(figurePaperId int) (interface{}, error) {
	var (
		id         int
		paper_id   int
		figure_id  int
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM figure_papers WHERE id = $1`
	err := database.DB.QueryRow(query, figurePaperId).Scan(
		&id,
		&paper_id,
		&figure_id,
		&created_at,
		&updated_at,
	)
	if err != nil {
		log.Println(err)
		return h.GeneralResponse{Message: err.Error()}, err
	}
	fp := FigurePaper{
		Id:        id,
		PaperId:   paper_id,
		FigureId:  figure_id,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	log.Println("Successful GET to find FigurePaper:", id)
	return fp, nil
}

func CreateFigurePaper(body io.Reader) (interface{}, error) {
	var fp FigurePaper
	err := json.NewDecoder(body).Decode(&fp)
	if err != nil {
		return h.GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO figure_papers (paper_id, figure_id, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query, fp.PaperId, fp.FigureId)
	if err != nil {
		log.Println(err)
		return h.GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful POST to create FigurePaper")
		return h.GeneralResponse{Message: "FigurePaper created successfully"}, nil
	}
}

func UpdateFigurePaper(figurePaperId int, body io.Reader) (interface{}, error) {
	var fp FigurePaper
	err := json.NewDecoder(body).Decode(&fp)
	query := `
		UPDATE figure_papers
		SET 
			paper_id = $2,
			figure_id = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $1
	`
	_, err = database.DB.Exec(query, figurePaperId, fp.PaperId, fp.FigureId)
	if err != nil {
		log.Println(err)
		return h.GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful PUT/PATCH to update FigurePaper")
		return h.GeneralResponse{Message: "FigurePaper updated successfully"}, nil
	}
}

func DeleteFigurePaper(figurePaperId int) (h.GeneralResponse, error) {
	query := `DELETE FROM figure_papers WHERE id=$1`
	res, err := database.DB.Exec(query, figurePaperId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete FigurePaper with id %d", figurePaperId)
		err = errors.New("Did not find row with specified ID")
		return h.GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return h.GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return h.GeneralResponse{Message: "FigurePaper deleted successfully"}, nil
}
