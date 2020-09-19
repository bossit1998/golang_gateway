package models

type CalcDeliveryCostRequest struct {
	MinDistance    float64        `json:"min_distance"`
	MinPrice       float64        `json:"min_price"`
	PerKmPrice     float64        `json:"per_km_price"`
	TripsDataModel TripsDataModel `json: "trips_data"`
}

type CalcDeliveryCostResponse struct {
	Distance float64 `json:"distance"`
	Price    float64 `json:"price"`
}
