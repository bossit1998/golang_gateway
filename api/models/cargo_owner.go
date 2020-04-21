package models

import pbco "genproto/co_service"

type cargoOwnerModel struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo" binding:"required"`
	Login string `json:"login" binding:"required"`
	Password string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
}

type CreateCargoOwner struct {
	cargoOwnerModel
}

type CheckNameRequest struct {
	Name string `json:"name" binding:"required"`
}

type CheckExistsResponse struct {
	pbco.CheckExistsResponse
}

type CheckLoginRequest struct {
	Login string `json:"login" binding:"required"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type ChangeLoginPasswordRequest struct {
	Login string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetCO struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
	Login string `json:"login"`
	PhoneNumber string `json:"phone_number"`
	Token string `json:"token"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
}