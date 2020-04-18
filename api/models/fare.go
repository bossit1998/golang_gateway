package models

type (
	//GetFareResponseModel ...
	GetFareResponseModel struct {
		ID           string  `json:"id"`
		Name         string  `json:"name"`
		DeliveryTime uint64  `json:"delivery_time"`
		PricePerKm   float32 `json:"price_per_km"`
		MinPrice     float32 `json:"min_price"`
		MinDistance  float64 `json:"min_distance"`
		CreatedAt    string  `json:"created_at"`
		UpdatedAt    string  `json:"updated_at"`
		DeletedAt    string  `json:"deleted_at"`
	}

	// DeleteFareModel ...
	DeleteFareModel struct {
		ID string `json:"fare_id" example:"965b0929-82e1-4a53-ad0b-a16f50c99669"`
	}

	//GetAllFaresResponseModel ...
	GetAllFaresResponseModel struct {
		Count int                    `json:"count"`
		Fares []GetFareResponseModel `json:"fares"`
	}
)
