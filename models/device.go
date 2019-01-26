package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Device struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetAllDevices() ([]Device, error) {
	var devices []Device
	var (
		id         int
		name       string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM devices;`
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
		device := Device{
			Id:        id,
			Name:      name,
			CreatedAt: created_at.String(),
			UpdatedAt: updated_at.String(),
		}
		devices = append(devices, device)
	}
	if err != nil {
		return []Device{}, err
	}
	log.Println("Successful GET to devices index")
	return devices, nil
}

func FindDevice(deviceId int) (interface{}, error) {
	var (
		id         int
		name       string
		created_at time.Time
		updated_at time.Time
	)
	query := `SELECT * FROM devices WHERE devices.id = $1`
	err := database.DB.QueryRow(query, deviceId).Scan(
		&id,
		&name,
		&created_at,
		&updated_at,
	)
	if err != nil {
		log.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	}
	device := Device{
		Id:        id,
		Name:      name,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	log.Println("Successful GET to find device: ", id)
	return device, nil
}

func CreateDevice(body io.Reader) (interface{}, error) {
	var d Device
	err := json.NewDecoder(body).Decode(&d)
	if err != nil {
		return GeneralResponse{Message: err.Error()}, err
	}
	query := `
		INSERT INTO devices (name, created_at, updated_at)
		VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`
	_, err = database.DB.Exec(query, d.Name)
	if err != nil {
		log.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful POST to create device")
		return GeneralResponse{Message: "Device created successfully"}, nil
	}
}

func UpdateDevice(deviceId int, body io.Reader) (interface{}, error) {
	var d Device
	err := json.NewDecoder(body).Decode(&d)
	query := `
		UPDATE devices
		SET 
			name = $2,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			devices.id = $1
	`
	_, err = database.DB.Exec(query, deviceId, d.Name)
	if err != nil {
		log.Println(err)
		return GeneralResponse{Message: err.Error()}, err
	} else {
		log.Println("Successful PUT/PATCH to update device")
		return GeneralResponse{Message: "Device updated successfully"}, nil
	}
}

func DeleteDevice(deviceId int) (GeneralResponse, error) {
	query := `DELETE FROM devices WHERE id=$1`
	res, err := database.DB.Exec(query, deviceId)
	rowCount, err := res.RowsAffected()
	if rowCount == 0 {
		errorMessage := fmt.Sprintf("Error when trying to delete device with id %d", deviceId)
		err = errors.New("Did not find row with specified ID")
		return GeneralResponse{Message: errorMessage}, err
	} else if err != nil {
		return GeneralResponse{Message: "Error with DELETE request"}, err
	}
	return GeneralResponse{Message: "Device deleted successfully"}, nil
}

func GetDevicesForPaper(paperId int) <-chan []Device {
	ch := make(chan []Device)
	go func() {
		var devices []Device
		var (
			id         int
			name       string
			created_at time.Time
			updated_at time.Time
		)
		query := `
			SELECT devices.* FROM devices 
			INNER JOIN device_papers ON devices.id = device_papers.device_id
			INNER JOIN papers ON device_papers.paper_id = papers.id
			WHERE papers.id = $1
		`
		rows, err := database.DB.Query(query, paperId)
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
			device := Device{
				Id:        id,
				Name:      name,
				CreatedAt: created_at.String(),
				UpdatedAt: updated_at.String(),
			}
			devices = append(devices, device)
		}
		if err != nil {
			log.Println("Error getting paper's devices", err)
		}
		ch <- devices
	}()
	return ch
}
