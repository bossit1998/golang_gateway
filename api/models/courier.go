package models

//Courier
type GetCourierModel struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Phone       string `json:"phone"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CreatedAt   string `json:"created_at"`
	IsActive    string `json:"is_active"`
}

type GetAllCouriersModel struct {
	Count    int               `json:"count"`
	Couriers []GetCourierModel `json:"couriers"`
}

type CreateCourierModel struct {
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateCourierModel struct {
	Id        string `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//Courier Details
type CourierDetailsModel struct {
	PassportNumber    string `json:"passport_number"`
	Gender            string `json:"gender"`
	BirthDate         string `json:"birth_date"`
	Address           string `json:"address"`
	Img               string `json:"img"`
	LisenseNumber     string `json:"lisense_number"`
	LisenseGivenDate  string `json:"lisense_given_date"`
	LisenseExpiryDate string `json:"lisense_expiry_date"`
}

//Courier Vehicle
type GetCourierVehicleModel struct {
	ID            string `json:"id"`
	Model         string `json:"model"`
	VehicleNumber string `json:"vehicle_number"`
	CreatedAt     string `json:"created_at"`
}

type GetAllCourierVehiclesModel struct {
	Count           int                      `json:"count"`
	CourierVehicles []GetCourierVehicleModel `json:"courier_vehicles"`
}

type CreateCourierVehicleModel struct {
	Model         string `json:"model"`
	VehicleNumber string `json:"vehicle_number"`
}

type UpdateCourierVehicleModel struct {
	Id            string `json:"id"`
	Model         string `json:"model"`
	VehicleNumber string `json:"vehicle_number"`
}
