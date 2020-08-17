package models

type GetReportModel struct {
	File string `json:"file"`
}

type OperatorReport struct {
	Fullname   string  `json:"fullname"`
	Username   string  `json:"username"`
	Phone      string  `json:"phone"`
	Total      int     `json:"total"`
	AvgPerHour float64 `json:"average_per_hour"`
	Bot        int     `json:"bot"`
	AdminPanel int     `json:"admin_panel"`
	App        int     `json:"app"`
	Website    int     `json:"website"`
}

type OperatorsReport struct {
	Reports []OperatorReport `json:"reports"`
}
