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

//GetCourierModel ...
type GetCourierModel struct {
	Courier GetCourierResponseModel `json:"courier"`
}

//GetAllEventsResponseModel ...
type GetAllCouriersResponseModel struct {
	Count    int                       `json:"count"`
	Couriers []GetCourierResponseModel `json:"couriers"`
}
