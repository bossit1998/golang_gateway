package models

//CreateVendorModel ...
type CreateVendorModel struct {
	CargoOwnerID string `json:"cargo_owner_id"`
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Location     Location `json:"location"`
}

//UpdateVendorModel ...
type UpdateVendorModel struct {
	ID           string `json:"id"`
	CargoOwnerID string `json:"cargo_owner_id"`
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	IsActive     bool   `json:"is_active"`
	Address      string `json:"address"`
	Location     Location `json:"location"`
}

//GetVendorModel ...
type GetVendorModel struct {
	ID           string `json:"id"`
	CargoOwnerID string `json:"cargo_owner_id"`
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	Address      string `json:"address"`
	Location     Location `json:"location"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
}

//GetAllVendorsModel ...
type GetAllVendorsModel struct {
	Count  int              `json:"count"`
	Vendor []GetVendorModel `json:"vendors"`
}

//CheckVendorLoginRequest ...
type CheckVendorLoginRequest struct {
	Phone string `json:"phone"`
}

//CheckVendorLoginResponse ...
type CheckVendorLoginResponse struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmVendorLoginRequest ...
type ConfirmVendorLoginRequest struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmVendorLoginResponse ...
type ConfirmVendorLoginResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}
