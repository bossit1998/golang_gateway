package models

//GetCourierModel ...
type GetCourierModel struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Phone       string `json:"phone"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CreatedAt   string `json:"created_at"`
	IsActive    string `json:"is_active"`
}

//GetAllCouriersModel ...
type GetAllCouriersModel struct {
	Count    int               `json:"count"`
	Couriers []GetCourierModel `json:"couriers"`
}

//GetAllCourierBranchesModel ...
type GetAllCourierBranchesModel struct {
	BranchIds []string `json:"branch_ids"`
}

//SearchCouriersByModel ...
type SearchCouriersByPhoneModel struct {
	Couriers []GetCourierModel `json:"couriers"`
}

//GetAllCouriersModel ...
type SaveCourierVendorsModel struct {
	CourierId     	  string `json:"courier_id"`
	VendorsIds []string `json:"vendor_ids"`
}

//CreateCourierModel ...
type CreateCourierModel struct {
	Phone     	  string `json:"phone"`
	FirstName 	  string `json:"first_name"`
	LastName  	  string `json:"last_name"`
	DistributorID string `json:"distributor_id"`
}

//CreateBranchCourierModel ...
type CreateBranchCourierModel struct {
	BranchID string `json:"branch_id"`
	CourierID string `json:"courier_id"`
}

//UpdateCourierModel ...
type UpdateCourierModel struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//CourierDetailsModel ...
type CourierDetailsModel struct {
	CourierID         string `json:"courier_id"`
	PassportNumber    string `json:"passport_number"`
	Gender            string `json:"gender"`
	BirthDate         string `json:"birth_date"`
	Address           string `json:"address"`
	Img               string `json:"img"`
	LisenseNumber     string `json:"lisense_number"`
	LisenseGivenDate  string `json:"lisense_given_date"`
	LisenseExpiryDate string `json:"lisense_expiry_date"`
}

//GetCourierVehicleModel ...
type GetCourierVehicleModel struct {
	ID            string `json:"id"`
	Model         string `json:"model"`
	VehicleNumber string `json:"vehicle_number"`
	CreatedAt     string `json:"created_at"`
}

//GetAllCourierVehiclesModel ...
type GetAllCourierVehiclesModel struct {
	Count           int                      `json:"count"`
	CourierVehicles []GetCourierVehicleModel `json:"courier_vehicles"`
}

//CreateCourierVehicleModel ...
type CreateCourierVehicleModel struct {
	Model         string `json:"model"`
	VehicleNumber string `json:"vehicle_number"`
}

//UpdateCourierVehicleModel ...
type UpdateCourierVehicleModel struct {
	ID            string `json:"id"`
	Model         string `json:"model"`
	VehicleNumber string `json:"vehicle_number"`
}

//CheckLoginResponse ...
type CheckLoginResponse struct {
	Code string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmLoginRequest ...
type ConfirmLoginRequest struct {
	Code string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmLoginResponse ...
type ConfirmLoginResponse struct {
	ID string `json:"id"`
	AccessToken string `json:"access_token"`
}