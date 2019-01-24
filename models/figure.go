package models

import (
"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Figure struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	FigureType string `json:"figure_type"`
	ImageFile  string `json:"image_file"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func GetAllFigures() ([]Figure, error) {
	var figures []Figure
	var (
		id          int
		name        string
		figure_type string
		image_file  string
		created_at  time.Time
		updated_at  time.Time
	)
	query := `SELECT * FROM figures;`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&name,
			&figure_type,
			&image_file,
			&created_at,
			&updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		figure := Figure{
			Id:         id,
			Name:       name,
			FigureType: figure_type,
			ImageFile:  image_file,
			CreatedAt:  created_at.String(),
			UpdatedAt:  updated_at.String(),
		}
		figures = append(figures, figure)
	}
	if err != nil {
		return []Figure{}, err
	}
	fmt.Println("Successful GET to figures index")
	return figures, nil
}

func FindFigure(figureId int) (interface{}, error) {
	var (
		id          int
		name        string
		figure_type string
		image_file  string
		created_at  time.Time
		updated_at  time.Time
	)
	query := `SELECT * FROM figures WHERE figures.id = $1`
	err := database.DB.QueryRow(query, figureId).Scan(
		&id,
		&name,
		&figure_type,
		&image_file,
		&created_at,
		&updated_at,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	figure := Figure{
		Id:         id,
		Name:       name,
		FigureType: figure_type,
		ImageFile:  image_file,
		CreatedAt:  created_at.String(),
		UpdatedAt:  updated_at.String(),
	}
	fmt.Println("Successful GET to find figure:", id)
	return figure, nil
}

func CreateFigure(body io.Reader) (interface{}, error) {
	var f Figure
	err := json.NewDecoder(body).Decode(&f)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO figures (name, figure_type, image_file, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query, f.Name, f.FigureType, f.ImageFile)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful POST to create figure")
		return GeneralResponse{Message: "Figure created successfully"}, nil
	}
}

func UpdateFigure(figureId int, body io.Reader) (interface{}, error) {
	var f Figure
	err := json.NewDecoder(body).Decode(&f)
	query := `
		UPDATE figures
		SET 
			name = $2,
			figure_type = $3,
			image_file = $4,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			Figures.id = $1
	`
	_, err = database.DB.Exec(query, figureId, f.Name, f.FigureType, f.ImageFile)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update figure")
		return GeneralResponse{Message: "Figure updated successfully"}, nil
	}
}

func DeleteFigure(figureId int) (GeneralResponse, error) {
	query := `DELETE FROM figures WHERE id=$1`
	res, err := database.DB.Exec(query, figureId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete figure with id %d", figureId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Figure deleted successfully"}, nil
}

func GetFiguresForPaper(paperId int) <-chan []Figure {
	ch := make(chan []Figure)
	go func() {
		var figures []Figure
		var (
			id          int
			name        string
			figure_type string
			image_file  string
			created_at  time.Time
			updated_at  time.Time
		)
		query := `
			SELECT figures.* FROM figures 
			INNER JOIN figure_papers ON figures.id = figure_papers.figure_id
			INNER JOIN papers ON papers.id = figure_papers.paper_id
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
				&name,
				&figure_type,
				&image_file,
				&created_at,
				&updated_at,
			)
			if err != nil {
				fmt.Println(err)
			}
			figure := Figure{
				Id:         id,
				Name:       name,
				FigureType: figure_type,
				ImageFile:  image_file,
				CreatedAt:  created_at.String(),
				UpdatedAt:  updated_at.String(),
			}
			figures = append(figures, figure)
		}
		if err != nil {
			fmt.Println("Error getting paper's figures", err)
		}
		ch <- figures
	}()
	return ch
}
