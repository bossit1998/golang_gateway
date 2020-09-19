package models

type AuthorizationModel struct {
	Token string `header:"Authorization"`
}

type UserInfo struct {
	ID       string `json:"id"`
	UserType string `json:"role"`
	// ShipperID string `json:"shipper_id"`

	UserID     string `json:"user_id"`
	ClientID   string `json:"client_id"`
	ShipperID  string `json:"shipper_id"`
	UserTypeID string `json:"user_type_id"`
}
