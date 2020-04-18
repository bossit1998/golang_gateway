package models

//Distributor
type GetDistributorModel struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type GetAllDistributorsModel struct {
	Count        int                   `json:"count"`
	Distributors []GetDistributorModel `json:"distributors"`
}

type CreateDistributorModel struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

type UpdateDistributorModel struct {
	ID    string `json:"id"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}
