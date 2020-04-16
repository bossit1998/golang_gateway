package models

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Quantity float64 `json:"quantity"`
	Price float64 `json:"price"`
}

type Location struct {
	Long float64 `json:"long"`
	Lat float64`json:"lat"`
} 

type Order struct {
	BranchID string `json:"branch_id"`
	FromLocation Location `json:"from_location"`
	FromAddress string `json:"from_address"`
	ToLocation Location `json:"to_location"`
	ToAddress string `json:"to_address"`
	PhoneNumber string `json:"phone_number"`
	FareID string `json:"fare_id"`
	CoID string `json:"co_id"`
	CreatorTypeID string `json:"creator_type_id"`
	UserID string `json:"user_id"`
	Products Product `json:"products"`
}
