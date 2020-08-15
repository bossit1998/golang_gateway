package models

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//LoginResponse
type LoginResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	ClientID     string `json:"client_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserRoleID   string `json:"user_role_id"`
	UserType     string `json:"user_type"`
}

//OTPLoginRequest ...
type OTPLoginRequest struct {
	Phone string `json:"phone" binding:"required"`
}

//OTPLoginResponse ...
type OTPLoginResponse struct {
	Phone string `json:"phone"`
}

//OTPConfirmRequest ...
type OTPConfirmRequest struct {
	Code  string `json:"code" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
