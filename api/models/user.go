package models

//CreateUserModel ...
type CreateUserModel struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
}

//UpdateUserModel ...
type UpdateUserModel struct {
	ID    string `json:"id"`
	Name  string `json:"first_name"`
	Phone string `json:"phone"`
}

//GetUserModel ...
type GetUserModel struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
	Name        string `json:"first_name"`
	Phone       string `json:"phone"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
}

//GetAllUsersModel ...
type GetAllUsersModel struct {
	Count int            `json:"count"`
	Users []GetUserModel `json:"users"`
}
