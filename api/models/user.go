package models

//CreateUserModel ...
type CreateUserModel struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//UpdateUserModel ...
type UpdateUserModel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//GetUserModel ...
type GetUserModel struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
}

//GetAllUsersModel ...
type GetAllUsersModel struct {
	Count int            `json:"count"`
	Users []GetUserModel `json:"users"`
}

//CheckUserLoginRequest ...
type CheckUserLoginRequest struct {
	Phone string `json:"phone"`
}

//CheckUserLoginResponse ...
type CheckUserLoginResponse struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmUserLoginRequest ...
type ConfirmUserLoginRequest struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmUserLoginResponse ...
type ConfirmUserLoginResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}

//SearchByPhoneResponse ...
type SearchByPhoneResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
