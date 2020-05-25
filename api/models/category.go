package models

type Category struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	ParentID string `json:"parent_id"`
	Description string `json:"description"`
}

type GetAllCategory struct {
	Categories []Category `json:"categories"`
	Count int64 `json:"count,string"`
} 
