package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type DevicePaper struct {
	Id        int    `json:"id"`
	PaperId   int    `json:"paper_id"`
	DeviceId  int    `json:"device_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllDevicePapers(authToken string) ([]DevicePaper, error) {
	if !ValidToken(authToken) {
		return []DevicePaper{}, errors.New("Unauthorized")
	}
	var dps []DevicePaper
	var (
		id         int
		paper_id   int
		device_id  int
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM device_papers`
	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&id,
			&paper_id,
			&device_id,
			&created_at,
			&updated_at,
		)
		if err != nil {
			fmt.Println(err)
		}
		dp := DevicePaper{
			Id:        id,
			PaperId:   paper_id,
			DeviceId:  device_id,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		dps = append(dps, dp)
	}
	if err != nil {
		return []DevicePaper{}, err
	}
	fmt.Println("Successful GET to DevicePapers index")
	return dps, nil
}

func FindDevicePaper(devicePaperId int, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []DevicePaper{}, errors.New("Unauthorized")
	}
	var (
		id         int
		paper_id   int
		device_id  int
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM device_papers WHERE id = $1`
	err := database.DB.QueryRow(query, devicePaperId).Scan(
		&id,
		&paper_id,
		&device_id,
		&created_at,
		&updated_at,
	)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	dp := DevicePaper{
		Id:        id,
		PaperId:   paper_id,
		DeviceId:  device_id,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	fmt.Println("Successful GET to find DevicePaper:", id)
	return dp, nil
}

func CreateDevicePaperQuery(dp DevicePaper) (int, error) {
	query := `
		INSERT INTO device_papers (paper_id, device_id, created_at, updated_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id
	`
	lastInsertId := 0
	err := database.DB.QueryRow(query, dp.PaperId, dp.DeviceId).Scan(&lastInsertId)
	if err != nil {
		fmt.Println(err)
	}
	return lastInsertId, err
}

func CreateDevicePaper(body io.Reader, authToken string) (interface{}, error) {
	if !ValidToken(authToken) {
		return []DevicePaper{}, errors.New("Unauthorized")
	}
	var dp DevicePaper
	err := json.NewDecoder(body).Decode(&dp)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	_, err = CreateDevicePaperQuery(dp)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful POST to create DevicePaper")
		return GeneralResponse{Message: "DevicePaper created successfully"}, nil
	}
}

func UpdateDevicePaper(devicePaperId int, body io.Reader, authToken string) (GeneralResponse, error) {
	if !ValidToken(authToken) {
		return []DevicePaper{}, errors.New("Unauthorized")
	}
	var dp DevicePaper
	err := json.NewDecoder(body).Decode(&dp)
	query := `
		UPDATE device_papers
		SET 
			paper_id = $2,
			device_id = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			id = $1
	`
	_, err = database.DB.Exec(query, devicePaperId, dp.PaperId, dp.DeviceId)
	if err != nil {
		fmt.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		fmt.Println("Successful PUT/PATCH to update DevicePaper")
		return GeneralResponse{Message: "DevicePaper updated successfully"}, nil
	}
}

func DeleteDevicePaper(devicePaperId int, authToken string) (GeneralResponse, error) {
	if !ValidToken(authToken) {
		return GeneralResponse{Message: "Unauthorized"}, errors.New("Unauthorized")
	}
	query := `DELETE FROM device_papers WHERE id=$1`
	res, err := database.DB.Exec(query, devicePaperId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete DevicePaper with id %d", devicePaperId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "DevicePaper deleted successfully"}, nil
}
