package models

type AuthorizationModel struct {
	Token string `header:"Authorization"`
}

type UserInfo struct {
	ID string `json:"id"`
	Role string `json:"role"`
}
