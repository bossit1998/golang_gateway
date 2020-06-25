package models

type AuthorizationModel struct {
	Token string `header:"Authorization"`
}

type UserInfo struct {
	ID string `json:"id"`
	UserType string `json:"role"`
	ShipperID string `json:"shipper_id"`
}
