package models

//CreateCustomerModel ...
type CreateCustomerModel struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//UpdateCustomerModel ...
type UpdateCustomerModel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//GetCustomerModel ...
type GetCustomerModel struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
}

//GetAllCustomersModel ...
type GetAllCustomersModel struct {
	Count     int                `json:"count"`
	Customers []GetCustomerModel `json:"customers"`
}

//CheckCustomerLoginRequest ...
type CustomerLoginRequest struct {
	Phone string `json:"phone"`
}

//CheckCustomerLoginResponse ...
type CustomerLoginResponse struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmCustomerLoginRequest ...
type ConfirmCustomerLoginRequest struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmCustomerLoginResponse ...
type ConfirmCustomerLoginResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}

//SearchByPhoneResponse ...
type SearchByPhoneResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//CustomerExists
type CustomerExists struct {
	Exists bool `json:"exists"`
}
