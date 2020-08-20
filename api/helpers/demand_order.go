package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"bitbucket.org/alien_soft/api_getaway/config"
	"bitbucket.org/alien_soft/api_getaway/pkg/logger"
)

func SendPush(orderID, statusID string, log logger.Logger) {
	values, err := json.Marshal(map[string]string{
		"status_id": statusID,
	})
	if err != nil {
		log.Error("Error while marshaling", logger.Error(err))
		return
	}

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST",
		config.AliftechURL+orderID+"/change-status",
		bytes.NewBuffer(values))

	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Access-Token", config.Load().AliftechAccessToken)
	_, err = client.Do(request)
	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}
}
