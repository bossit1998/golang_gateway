package models

//CreateShipperModel ...
type CreateShipperModel struct {
	Name        string   `json:"name"`
	UserName    string   `json:"username"`
	Password    string   `json:"password"`
	Logo        string   `json:"logo"`
	Description string   `json:"description"`
	Phone       []string `json:"phone"`
	UserRoleID  string   `json:"user_role_id"`
}

//UpdateShipperModel ...
type UpdateShipperModel struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	UserName    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       []string `json:"phone"`
	Logo        string   `json:"logo"`
	Description string   `json:"description"`
}

//GetShipperModel ...
type GetShipperModel struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	UserName    string   `json:"username"`
	Password    string   `json:"password"`
	Phone       []string `json:"phone"`
	Logo        string   `json:"logo"`
	Description string   `json:"description"`
	IsActive    bool     `json:"is_active"`
	CreatedAt   string   `json:"created_at"`
}

//GetAllShippersModel ...
type GetAllShippersModel struct {
	Count    int               `json:"count"`
	Shippers []GetShipperModel `json:"shippers"`
}

//Login Model
type ShipperLogin struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ShipperChangePassword struct {
	Password string `json:"password" binding:"required"`
}
