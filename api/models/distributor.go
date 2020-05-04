package models

//Distributor
type GetDistributorModel struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	CreatedAt   string `json:"created_at"`
	IsActive    string `json:"is_active"`
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

//Park
type GetParkModel struct {
	ID            string        `json:"id"`
	DistributorId string        `json:"distributor_id"`
	Name          string        `json:"name"`
	Location      LocationModel `json:"location"`
	Address       string        `json:"address"`
	CreatedAt     string        `json:"created_at"`
	IsActive      string        `json:"is_active"`
}

type CreateParkModel struct {
	Phone    string        `json:"phone"`
	Name     string        `json:"name"`
	Location LocationModel `json:"location"`
	Address  string        `json:"address"`
}

type LocationModel struct {
	Long float64 `json:"long"`
	Lat  float64 `json:"lat"`
}

type GetAllParksModel struct {
	Count int            `json:"count"`
	Parks []GetParkModel `json:"parks"`
}
