package models

import (
	"fmt"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Figure struct {
	Id         int
	Name       string `json:"name"`
	FigureType string `json:"figure_type"`
	ImageFile  string `json:"image_file"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func GetFiguresForPaper(paperId int, channel chan []Figure) {
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
	channel <- figures
}
