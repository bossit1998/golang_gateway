package models

//SearchDistributor ...
type SearchDistributor struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//SearchDistributorModel ...
type SearchDistributorModel struct {
	Distributors []SearchDistributor `json:"distributors"`
}

//CreateDistributorRequestModel ...
type CreateDistributorRequestModel struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

//GetDistributorResponseModel ...
type GetDistributorResponseModel struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

//GetDistributorModel ...
type GetDistributorModel struct {
	Distributor GetDistributorResponseModel `json:"distributor"`
}

//GetAllEventsResponseModel ...
type GetAllDistributorsResponseModel struct {
	Count        int                           `json:"count"`
	Distributors []GetDistributorResponseModel `json:"distributors"`
}

//DeleteEventModel ...
type DeleteDistributorModel struct {
	ID string `json:"distributor_id" example:"965b0929-82e1-4a53-ad0b-a16f50c99669"`
}
