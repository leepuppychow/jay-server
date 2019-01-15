package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func GetAllDataRequestForms(authToken string) ([]DataRequestForm, error) {
	// if !ValidToken(authToken) {
	// 	return []DataRequestForm{}, errors.New("Unauthorized")
	// }
	var drfs []DataRequestForm
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
	query := `SELECT * FROM data_request_forms;`
	rows, err := database.DB.Query(query)
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
		drf := DataRequestForm{
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
		drfs = append(drfs, drf)
	}
	if err != nil {
		return []DataRequestForm{}, err
	}
	fmt.Println("Successful GET to Data Request Forms index")
	return drfs, nil
}

func FindDataRequestForm(dataRequestFormId int, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return []DataRequestForm{}, errors.New("Unauthorized")
	// }
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
	query := `SELECT * FROM data_request_forms WHERE id = $1`
	err := database.DB.QueryRow(query, dataRequestFormId).Scan(
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
		return GeneralResponse{Message: err.Error()}, err
	}
	drf := DataRequestForm{
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
	fmt.Println("Successful GET to find Data Request Form:", id)
	return drf, nil
}

func CreateDataRequestForm(body io.Reader, authToken string) (interface{}, error) {
	// if !ValidToken(authToken) {
	// 	return []DataRequestForm{}, errors.New("Unauthorized")
	// }
	var drf DataRequestForm
	err := json.NewDecoder(body).Decode(&drf)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO data_request_forms (
			paper_id,
			round,
			form_completed,
			requested_delivery,
			actual_delivery,
			data_refinement_complete,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query,
		drf.PaperId,
		drf.Round,
		drf.FormCompleted,
		drf.RequestedDelivery,
		drf.ActualDelivery,
		drf.DataRefinementComplete,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful POST to create DataRequestForm")
		return GeneralResponse{Message: "DataRequestForm created successfully"}, nil
	}
}

func UpdateDataRequestForm(dataRequestFormId int, body io.Reader, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return []DataRequestForm{}, errors.New("Unauthorized")
	// }
	var drf DataRequestForm
	err := json.NewDecoder(body).Decode(&drf)
	query := `
		UPDATE data_request_forms
		SET 
			paper_id=$2,
			round=$3,
			form_completed=$4,
			requested_delivery=$5,
			actual_delivery=$6,
			data_refinement_complete=$7,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			data_request_forms.id = $1
	`
	_, err = database.DB.Exec(query,
		dataRequestFormId,
		drf.PaperId,
		drf.Round,
		drf.FormCompleted,
		drf.RequestedDelivery,
		drf.ActualDelivery,
		drf.DataRefinementComplete,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update DataRequestForm")
		return GeneralResponse{Message: "DataRequestForm updated successfully"}, nil
	}
}

func DeleteDataRequestForm(dataRequestFormId int, authToken string) (GeneralResponse, error) {
	// if !ValidToken(authToken) {
	// 	return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	// }
	query := `DELETE FROM data_request_forms WHERE id=$1`
	res, err := database.DB.Exec(query, dataRequestFormId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete DataRequestForm with id %d", dataRequestFormId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "DataRequestForm deleted successfully"}, nil
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
