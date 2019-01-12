package models

import (
	"fmt"
	"time"

	"github.com/leepuppychow/jay_medtronic/database"
)

type Device struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetDevicesForPaper(paperId int, kawaiiChan chan []Device) {
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
	}
	if err != nil {
		fmt.Println(err)
	}
	device := Device{
		Id:        id,
		Name:      name,
		CreatedAt: created_at.String(),
		UpdatedAt: updated_at.String(),
	}
	devices = append(devices, device)
	if err != nil {
		fmt.Println("Error getting paper's devices", err)
	}
	kawaiiChan <- devices
}
