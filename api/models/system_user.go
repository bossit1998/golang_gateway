package models

//CreateSystemUserModel ...
type CreateSystemUserModel struct {
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Phone      []string `json:"phone"`
	UserRoleID string   `json:"user_role_id"`
	// ShipperID  string   `json:"shipper_id"`
}

//UpdateSystemUserModel ...
type UpdateSystemUserModel struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Phone      []string `json:"phone"`
	UserRoleID string   `json:"user_role_id"`
	// ShipperID  string   `json:"shipper_id"`
}

//GetSystemUserModel ...
type GetSystemUserModel struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Phone     []string `json:"phone"`
	IsActive  bool     `json:"is_active"`
	IsBlocked bool     `json:"is_blocked"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	ShipperID  string `json:"shipper_id"`
	UserRoleID string `json:"user_role_id"`
}

//GetAllSystemUsersModel ...
type GetAllSystemUsersModel struct {
	Count       int                  `json:"count"`
	SystemUsers []GetSystemUserModel `json:"shippers"`
}

//Login Model
type SystemUserLogin struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SystemUserChangePassword struct {
	Password string `json:"password" binding:"required"`
}
