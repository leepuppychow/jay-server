package models

import (
	"fmt"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
	"github.com/lib/pq"
)

type DataRequestForm struct {
	Id                     int    `json:"id"`
	PaperId                int    `json:"paper_id"`
	Round                  int    `json:"round"`
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
		form_completed           pq.NullTime
		requested_delivery       pq.NullTime
		actual_delivery          pq.NullTime
		data_refinement_complete pq.NullTime
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
			FormCompleted:          NullTimeCheck(form_completed),
			RequestedDelivery:      NullTimeCheck(requested_delivery),
			ActualDelivery:         NullTimeCheck(actual_delivery),
			DataRefinementComplete: NullTimeCheck(data_refinement_complete),
			CreatedAt:              created_at.String(),
			UpdatedAt:              updated_at.String(),
		}
		DRFs = append(DRFs, DRF)
	}

	if err != nil {
		fmt.Println("Error getting paper's DRFs", err)
	}
	kawaiiChan <- DRFs
}
