package models

type CreateMeasureModel struct {
	Name string `json:"name"`
	ShortName string `json:"short_name"`
	ProductKindID string `json:"product_kind_id"`
	OrderNo int64 `json:"order_no,string"`
	Description string `json:"description"`
}

type GetMeasureModel struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ShortName string `json:"short_name"`
	Slug string `json:"slug"`
	ProductKindID string `json:"product_kind_id"`
	Description string `json:"description"`
}

type GetAllMeasuresModel struct {
	Measures []GetMeasureModel `json:"measures"`
	Count int64 `json:"count,string"`
} 
