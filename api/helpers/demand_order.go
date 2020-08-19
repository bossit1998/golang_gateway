package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

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
		"https://services.test.aliftech.uz/api/gate/delever/"+orderID+"/change-status",
		bytes.NewBuffer(values))

	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Access-Token", "lkjISFALKFNQWIOJSALNFLKSMAG;KS;LDD!@3KDKLSAL")
	_, err = client.Do(request)
	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}
}

func SendSms(orderID string, log logger.Logger) {
	client := &http.Client{}
	request, err := http.NewRequest(
		"POST",
		"https://services.test.aliftech.uz/api/gate/delever/"+orderID+"/request-complete",
		nil)
	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}

	request.Header.Add("Authorization", "lkjISFALKFNQWIOJSALNFLKSMAG;KS;LDD!@3KDKLSAL")
	request.Header.Add("Content-Type", "application/json")
	_, err = client.Do(request)
	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}
}

func ConfirmSms(orderID, code string, log logger.Logger) {
	values, err := json.Marshal(map[string]string{
		"otp": code,
	})
	if err != nil {
		log.Error("Error while marshaling", logger.Error(err))
		return
	}

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST",
		"https://services.test.aliftech.uz/api/gate/delever/"+orderID+"/complete",
		bytes.NewBuffer(values))
	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Access-Token", "lkjISFALKFNQWIOJSALNFLKSMAG;KS;LDD!@3KDKLSAL")
	_, err = client.Do(request)
	if err != nil {
		log.Error("Error while sending push", logger.Error(err))
		return
	}
}
