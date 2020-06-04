package models

//CreateShipperModel ...
type CreateShipperModel struct {
	Name        string   `json:"name"`
	UserName    string   `json:"username"`
	Password    string   `json:"password"`
	Logo        string   `json:"logo"`
	Description string   `json:"description"`
	Phone       []string `json:"phone"`
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
	Count     int               `json:"count"`
	Shippers []GetShipperModel `json:"shippers"`
}

//CheckShipperLoginRequest ...
type CheckShipperLoginRequest struct {
	Phone string `json:"phone"`
}

//CheckShipperLoginResponse ...
type CheckShipperLoginResponse struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmShipperLoginRequest ...
type ConfirmShipperLoginRequest struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

//ConfirmShipperLoginResponse ...
type ConfirmShipperLoginResponse struct {
	ID          string `json:"id"`
	AccessToken string `json:"access_token"`
}
