package models

type GetFareModel struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	DeliveryTime uint64  `json:"delivery_time"`
	PricePerKm   float32 `json:"price_per_km"`
	MinPrice     float32 `json:"min_price"`
	MinDistance  float64 `json:"min_distance"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    string  `json:"deleted_at"`
}

type GetAllFaresModel struct {
	Count int            `json:"count"`
	Fares []GetFareModel `json:"fares"`
}

type CreateFareModel struct {
	Name         string  `json:"name"`
	DeliveryTime uint64  `json:"delivery_time"`
	PricePerKm   float32 `json:"price_per_km"`
	MinPrice     float32 `json:"min_price"`
	MinDistance  float64 `json:"min_distance"`
}

type UpdateFareModel struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	DeliveryTime uint64  `json:"delivery_time"`
	PricePerKm   float32 `json:"price_per_km"`
	MinPrice     float32 `json:"min_price"`
	MinDistance  float64 `json:"min_distance"`
}
