package utils

import (
	"crews_mock_server/models"
	"encoding/json"
	"net/http"
	"bytes"
	"log"
)

const urlAddress = "http://localhost:3000/api/update_crew"

func SendCoordinates(c *models.Crew) {
	jsonValue, _ := json.Marshal(c)
	_, err := http.Post(urlAddress, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err.Error())
	}
}