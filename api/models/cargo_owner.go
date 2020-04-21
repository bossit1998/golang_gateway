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