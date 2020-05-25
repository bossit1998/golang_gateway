package models

type CreateProductKindModel struct {
	Name string `json:"name"`
	Description string `json:"description"`
	OrderNo int64 `json:"order_no,string"`
}

type GetProductKindModel struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Slug string `json:"slug"`
}

type GetAllProductKindsModel struct {
	ProductKinds []GetProductKindModel `json:"product_kinds"`
	Count int64 `json:"count,string"`
}
