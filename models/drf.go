package models

import (
	"fmt"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type DataRequestForm struct {
	Id                     int
	PaperId                int
	Round                  int
	FormCompleted          string `json:"form_completed"`
	RequestedDelivery      string `json:"requested_delivery"`
	ActualDelivery         string `json:"actual_delivery"`
	DataRefinementComplete string `json:"data_refinement_complete"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}

func GetDataRequestFormsForPaper(paperId int, kawaiiChan chan []DataRequestForm) {
	var DRFs []DataRequestForm
	var (
		id                       int
		paper_id                 int
		round                    int
		form_completed           time.Time
		requested_delivery       time.Time
		actual_delivery          time.Time
		data_refinement_complete time.Time
		created_at               time.Time
		updated_at               time.Time
	)
	query := `
		SELECT data_request_forms.* FROM data_request_forms
		INNER JOIN papers ON data_request_forms.paper_id = papers.id
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
			&paper_id,
			&round,
			&form_completed,
			&requested_delivery,
			&actual_delivery,
			&data_refinement_complete,
			&created_at,
			&updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		DRF := DataRequestForm{
			Id:                     id,
			PaperId:                paper_id,
			Round:                  round,
			FormCompleted:          form_completed.String(),
			RequestedDelivery:      requested_delivery.String(),
			ActualDelivery:         actual_delivery.String(),
			DataRefinementComplete: data_refinement_complete.String(),
			CreatedAt:              created_at.String(),
			UpdatedAt:              updated_at.String(),
		}
		DRFs = append(DRFs, DRF)
	}

	if err != nil {
		fmt.Println("Error getting paper's figures", err)
	}
	kawaiiChan <- DRFs
}
