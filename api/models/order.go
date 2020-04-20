package models

type Location struct {
	Long float64 `json:"long" example:"60.123"`
	Lat float64`json:"lat" example:"40.123"`
}

type productModel struct {
	Name string `json:"name" example:"Choyxona Osh"`
	Quantity float64 `json:"quantity" example:"2"`
	Price float64 `json:"price" example:"25000"`
}

type orderModel struct {
	BranchID string `json:"branch_id" example:"a010f178-da52-4373-aacd-e477d871e27a"`
	FromLocation Location `json:"from_location"`
	FromAddress string `json:"from_address" example:"Hamid Olimjon maydoni 10A dom 40-kvartira"`
	ToLocation Location `json:"to_location"`
	ToAddress string `json:"to_address" example:"Hamid Olimjon maydoni 10A dom 40-kvartira"`
	PhoneNumber string `json:"phone_number" example:"998998765432"`
	FareID string `json:"fare_id" example:"a010f178-da52-4373-aacd-e477d871e27a"`
	CoID string `json:"co_id" example:"a010f178-da52-4373-aacd-e477d871e27a"`
	CreatorTypeID string `json:"creator_type_id" example:"a010f178-da52-4373-aacd-e477d871e27a"`
	UserID string `json:"user_id" example:"a010f178-da52-4373-aacd-e477d871e27a"`
	DeliverPrice float64 `json:"deliver_price" example:"10000"`
}

type CreateOrder struct {
	orderModel
	Products []productModel `json:"products"`
}

type getOrderProductModel struct {
	productModel
	ID string `json:"id" `
	TotalAmount float64 `json:"total_amount"`
}

type getOrderModel struct {
	orderModel
	ID string `json:"id" example:"701dc270-0adc-4d00-ae78-4f2f78d794cc"`
	StatusID string `json:"status_id" example:"52f248b4-23a0-4350-80b7-1704eaff6c8c"`
}

type GetOrder struct {
	getOrderModel
	Products []getOrderProductModel `json:"products"`
}

type GetOrders struct {
	Orders []getOrderModel `json:"orders"`
	Count uint64 `json:"count"`
}

type ChangeStatusRequest struct {
	StatusID string `json:"status_id"`
} 