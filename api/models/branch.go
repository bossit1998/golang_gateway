package models

//CreateBranchModel ...
type CreateBranchModel struct {
	ShipperID   string   `json:"shipper_id"`
	Name        string   `json:"name"`
	UserName    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       []string `json:"phone"`
	Address     string   `json:"address"`
	Destination string   `json:"destination"`
	Location    Location `json:"location"`
}

//UpdateBranchModel ...
type UpdateBranchModel struct {
	ID          string   `json:"id"`
	ShipperID   string   `json:"shipper_id"`
	Name        string   `json:"name"`
	UserName    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       []string `json:"phone"`
	Address     string   `json:"address"`
	Destination string   `json:"destination"`
	Location    Location `json:"location"`
}

//GetBranchModel ...
type GetBranchModel struct {
	ID          string   `json:"id"`
	ShipperID   string   `json:"shipper_id"`
	Name        string   `json:"name"`
	UserName    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       []string `json:"phone"`
	Address     string   `json:"address"`
	Destination string   `json:"destination"`
	Location    Location `json:"location"`
	IsActive    bool     `json:"is_active"`
	CreatedAt   string   `json:"created_at"`
}

//GetAllBranchesModel ...
type GetAllBranchesModel struct {
	Count    int              `json:"count"`
	Branches []GetBranchModel `json:"branches"`
}

//CheckBranchLoginRequest ...
type CheckBranchLoginRequest struct {
	Phone string `json:"phone"`
}

//CheckBranchLoginResponse ...
type CheckBranchLoginResponse struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmBranchLoginRequest ...
type ConfirmBranchLoginRequest struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmBranchLoginResponse ...
type ConfirmBranchLoginResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}


