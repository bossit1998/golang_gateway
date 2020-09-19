package models

type CreateCategoryModel struct {
	Name        string `json:"name"`
	ParentID    string `json:"parent_id"`
	Description string `json:"description"`
	OrderNo     int64  `json:"order_no"`
	Image       string `json:"image"`
}

type GetCategoryModel struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	Slug            string             `json:"slug"`
	ParentID        string             `json:"parent_id"`
	Description     string             `json:"description"`
	ChildCategories []GetCategoryModel `json:"child_categories"`
	Image           string             `json:"image"`
	// OrderNo         uint64             `json:"order_no"`
}

type GetAllCategoriesModel struct {
	Categories []GetCategoryModel `json:"categories"`
	Count      int64              `json:"count,string"`
}
