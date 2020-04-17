package models

//SearchCourier ...
type SearchCourier struct {
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//SearchCourierModel ...
type SearchCourierModel struct {
	Courier []SearchCourier `json:"couriers"`
}

//GetCourierResponseModel ...
type GetCourierResponseModel struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

//GetCourierResponseModel ...
type GetCourierDetailsModel struct {
	PassportNumber    string `json:"passport_number"`
	Gender            string `json:"gender"`
	BirthDate         string `json:"birth_date"`
	Address           string `json:"address"`
	Img               string `json:"img"`
	LisenseNumber     string `json:"lisense_number"`
	LisenseGivenDate  string `json:"lisense_given_date"`
	LisenseExpiryDate string `json:"lisense_expiry_date"`
}

//GetCourierModel ...
type GetCourierModel struct {
	Courier GetCourierResponseModel `json:"courier"`
}

//GetAllEventsResponseModel ...
type GetAllCouriersResponseModel struct {
	Count    int                       `json:"count"`
	Couriers []GetCourierResponseModel `json:"couriers"`
}
