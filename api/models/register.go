package models

// RegisterModel ...
type RegisterModel struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

//RegisterConfirmModel ...
type RegisterConfirmModel struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
