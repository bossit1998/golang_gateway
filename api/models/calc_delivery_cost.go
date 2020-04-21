package models

type CalcDeliveryCostRequest struct {
	MinDistance float64 `json:"min_distance"`
	MinPrice float64 `json:"min_price"`
	PerKmPrice float64 `json:"per_km_price"`
	FromLocation Location `json:"from_location"`
	ToLocation Location `json:"to_location"`
}

type CalcDeliveryCostResponse struct {
	Distance float64 `json:"distance"`
	Price float64 `json:"price"`
}