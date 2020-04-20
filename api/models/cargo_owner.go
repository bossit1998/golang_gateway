package models

type cargoOwnerModel struct {
	Name string `json:"name" binding:"require"`
	Logo string `json:"logo" binding:"require"`
	Login string `json:"login" binding:"require"`
	Password string `json:"password" binding:"require"`
	PhoneNumber string `json:"phone_number"`
	Description string `json:"description"`
}

type CreateCargoOwner struct {
	cargoOwnerModel
}