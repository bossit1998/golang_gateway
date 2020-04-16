package models

type (
	//FareModel ...
	FareModel struct {
		ID           string `json:"id" example:"965b0929-82e1-4a53-ad0b-a16f50c99573"`
		Name         string `json:"name"`
		DeliveryTime int64  `json:"delivery_time"`
		PricePerKm   int64  `json:"price_per_km"`
		MinPrice     int64  `json:"min_price"`
		IsAcctive    bool   `json:"is_active"`
	}
	//FareGeozoneModel ...
	FareGeozoneModel struct {
		ID   string    `json:"id"`
		Fare FareModel `json: "fare"`
	}
	//CreateFareRequestModel ...
	CreateFareRequestModel struct {
		ID           string `json:"id" example:"965b0929-82e1-4a53-ad0b-a16f50c99573"`
		Name         string `json:"name"`
		DeliveryTime int64  `json:"delivery_time"`
		PricePerKm   int64  `json:"price_per_km"`
		MinPrice     int64  `json:"min_price"`
		IsAcctive    bool   `json:"is_active"`
	}
	//GetFareResponseModel ...
	GetFareResponseModel struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		DeliveryTime int64  `json:"delivery_time"`
		PricePerKm   int64  `json:"price_per_km"`
		MinPrice     int64  `json:"min_price"`
		IsAcctive    bool   `json:"is_active"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		DeletedAt    string `json:"deleted_at"`
	}
	//GetFareResponse ....
	GetFareResponse struct {
		Fare []FareModel `json:"fare"`
	}
	// GetAllFareResponseModel ...
	GetAllFareResponseModel struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		DeliveryTime int64  `json:"delivery_time"`
		PricePerKm   int64  `json:"price_per_km"`
		MinPrice     int64  `json:"min_price"`
		IsAcctive    bool   `json:"is_active"`
	}
	//GetAllFareResponse ...
	GetAllFareResponse struct {
		Count int `json:"count"`
	}
	// DeleteFareModel ...
	DeleteFareModel struct {
		ID string `json:"event_id" example:"965b0929-82e1-4a53-ad0b-a16f50c99669"`
	}
)
