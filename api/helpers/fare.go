package helpers

import (
	"bitbucket.org/alien_soft/fare_service/api/models"
	"bitbucket.org/alien_soft/fare_service/genproto/fare_service"
)

func FillFares(src []*fare_service.Fare) []models.FareModel {
	var (
		result, fResult []models.FareModel
	)
	for _, e := range src {
		d := models.FareModel{
			ID:           e.GetID(),
			Name:         e.GetName(),
			DeliveryTime: e.GetDeliveryTime(),
			PricePerKm:   e.GetPricePerKm(),
			MinPrice:     e.GetMinPrice(),
			IsActive:     e.GetIsActive(),
			CreatedAt:    e.GetCreatedAt(),
			UpdatedAt:    e.GetUpdatedAt(),
			DeletedAt:    e.GetDeletedAt(),
		}
		result = append(result, d)
	}
	for _, v := range result {
		if v.Err == nil {
			fResult = append(fResult, v)
		}
	}
	return fResult
}
