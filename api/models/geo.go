package models

type geozone struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
	Geometry string `json:"geometry"`
}

//GeozoneModel ...
type GeozoneModel struct {
	Geozones []geozone `json:"geozones"`
	Count    uint64    `json:"count"`
}
