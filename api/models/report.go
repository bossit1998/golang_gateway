package models

type GetReportModel struct {
	File string `json:"file"`
}

type OperatorsReportModel struct {
	Fullname   string `json:"fullname"`
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	Total      string `json:"total"`
	AvgPerHour string `json:"average_per_hour"`
	Bot        string `json:"bot"`
	AdminPanel string `json:"admin_panel"`
	App        string `json:"app"`
	Website    string `json:"website"`
}

type GetAllOperatorsReportModel struct {
	Reports []OperatorsReportModel `json:"reports"`
	Count   int                    `json:"count"`
}
