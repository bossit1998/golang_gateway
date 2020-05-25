package models

type CreateProductModel struct {
	Name string `json:"name"`
	Code string `json:"code"`
	ShortName string `json:"short_name"`
	PreviewText string `json:"preview_text"`
	MeasureID string `json:"measure_id"`
	ProductKindID string `json:"product_kind_id"`
	CategoryID string `json:"category_id"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
	OrderNo int64 `json:"order_no,string"`
	Price int64 `json:"price,string"`
}

type GetProductModel struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Slug string `json:"slug"`
	ShortName string `json:"short_name"`
	PreviewText string `json:"preview_text"`
	MeasureID string `json:"measure_id"`
	ProductKindID string `json:"product_kind_id"`
	CategoryID string `json:"category_id"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
	Price int64 `json:"price,string"`
}

type GetAllProductsModel struct {
	Products []GetProductModel `json:"products"`
	Count int64 `json:"count,string"`
} 