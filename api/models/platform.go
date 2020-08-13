package models

//GetPlatformModel ...
type GetPlatformModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//GetAllPlatformsModel ...
type GetAllPlatformsModel struct {
	Count     uint64             `json:"count"`
	Platforms []GetPlatformModel `json:"platforms"`
}
