package models

type Location struct {
	Long float64 `json:"long" example:"60.123"`
	Lat  float64 `json:"lat" example:"40.123"`
}

type productModel struct {
	Name     string  `json:"name" example:"Choyxona Osh"`
	Quantity float64 `json:"quantity" example:"2"`
	Price    float64 `json:"price" example:"25000"`
}

type orderModel struct {
	ToLocation          Location `json:"to_location"`
	ToAddress           string   `json:"to_address" example:"Hamid Olimjon maydoni 10A dom 40-kvartira"`
	CustomerName        string   `json:"customer_name" example:"Oybek"`
	CustomerPhoneNumber string   `json:"customer_phone_number" example:"998998765432"`
	FareID              string   `json:"fare_id" example:"a010f178-da52-4373-aacd-e477d871e27a"`
	CoDeliveryPrice     float64  `json:"co_delivery_price" example:"10000"`
	Description         string   `json:"description"`
	ExternalOrderID 	int64	 `json:"external_order_id"`
}

type updateProduct struct {
	productModel
	ID string `json:"id"`
}

type updateStep struct {
	step
	ID string `json:"id"`
}

type updateStepModel struct {
	updateStep
	Products []updateProduct `json:"products"`
}

type updateOrder struct {
	ID string `json:"id"`
	orderModel
}

type UpdateOrder struct {
	updateOrder
	steps []updateStepModel `json:"steps"`
}

type step struct {
	BranchName         string   `json:"branch_name"`
	Location           Location `json:"location"`
	Address            string   `json:"address"`
	DestinationAddress string   `json:"destination_address"`
	PhoneNumber        string   `json:"phone_number"`
	Description        string   `json:"description"`
}

type stepModel struct {
	step
	Products []productModel `json:"products"`
}

type CreateOrder struct {
	orderModel
	Steps []stepModel `json:"steps"`
}

type getOrderProductModel struct {
	productModel
	ID          string  `json:"id" `
	TotalAmount float64 `json:"total_amount"`
}

type getOrderModel struct {
	orderModel
	ID       string `json:"id" example:"701dc270-0adc-4d00-ae78-4f2f78d794cc"`
	StatusID string `json:"status_id" example:"52f248b4-23a0-4350-80b7-1704eaff6c8c"`
}

type GetOrder struct {
	getOrderModel
	Products []getOrderProductModel `json:"products"`
}

type GetOrders struct {
	Orders []getOrderModel `json:"orders"`
	Count  uint64          `json:"count"`
}

type ChangeStatusRequest struct {
	StatusID string `json:"status_id"`
}

type Status struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetStatuses struct {
	Statuses []Status `json:"statuses"`
}

type AddCourierRequest struct {
	CourierID string `json:"courier_id"`
}
