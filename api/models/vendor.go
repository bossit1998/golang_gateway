package models

//CreateVendorModel ...
type CreateVendorModel struct {
	ID           string `json:"id"`
	CargoOwnerID string `json:"cargo_owner_id"`
	Name         string `json:"name"`
	UserName     string `json:"user_name"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
}

//UpdateVendorModel ...
type UpdateVendorModel struct {
	ID           string `json:"id"`
	CargoOwnerID string `json:"cargo_owner_id"`
	Name         string `json:"first_name"`
	UserName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	IsActive     bool   `json:"is_active"`
}

//GetVendorModel ...
type GetVendorModel struct {
	ID           string `json:"id"`
	CargoOwnerID string `json:"cargo_owner_id"`
	Name         string `json:"first_name"`
	UserName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	IsActive     bool   `json:"is_active"`
	CreatedAt    string `json:"created_at"`
}

//GetAllUsersModel ...
type GetAllVendorsModel struct {
	Count  int              `json:"count"`
	Vendor []GetVendorModel `json:"vendors"`
}
