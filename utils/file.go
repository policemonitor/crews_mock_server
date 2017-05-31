package utils

import (
	"crews_mock_server/models"
	"io/ioutil"
	"encoding/json"
)

func ReadCrewsConfigs() (crews []*models.Crew) {
	data, err := ioutil.ReadFile("dump")
	if err != nil {
		return
	}

	json.Unmarshal(data, &crews)
	return
}
